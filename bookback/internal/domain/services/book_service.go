package services

import (
	"context"
	adapters2 "github.com/SShlykov/zeitment/bookback/internal/domain/adapters"
	entity2 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

// BookService описывает сервис для работы с книгами.
//
//go:generate mockgen -destination=../../../tests/mocks/domain/services/book_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services BookService
type BookService interface {
	CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, id string) (*models.Book, error)
	ListBooks(ctx context.Context, request models.RequestBook) ([]*models.Book, error)

	GetTableOfContentsByBookID(ctx context.Context, request models.RequestTOC) (*models.TableOfContents, error)
	TogglePublic(ctx context.Context, request models.ToggleBookRequest) (*models.Book, error)
}

type BookRepo interface {
	SimpleRepo[*entity2.Book]
	GetTOCSectionsFromChapters(ctx context.Context, bookID string) ([]*entity2.Section, error)
	GetTOCSectionsFromPages(ctx context.Context, bookID string) ([]*entity2.Section, error)
}

type bookService struct {
	repo BookRepo
}

// NewBookService создает новый экземпляр Service.
func NewBookService(repo BookRepo) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(ctx context.Context, request models.CreateBookRequest) (*models.Book, error) {
	book := adapters2.BookModelToEntity(request.Book)

	if book.Variables == nil {
		book.Variables = []string{}
	}
	id, err := s.repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return s.GetBookByID(ctx, id)
}

func (s *bookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return adapters2.BookEntityToModel(book), nil
}

func (s *bookService) UpdateBook(ctx context.Context, id string, request models.UpdateBookRequest) (*models.Book, error) {
	book, err := s.repo.Update(ctx, id, adapters2.BookModelToEntity(request.Book))
	if err != nil {
		return nil, err
	}

	return adapters2.BookEntityToModel(book), nil
}

func (s *bookService) DeleteBook(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, err
}

func (s *bookService) ListBooks(ctx context.Context, request models.RequestBook) ([]*models.Book, error) {
	options := dbutils.NewPagination(&request.Options)

	books, err := s.repo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters2.BooksEntityToModel(books), nil
}

func (s *bookService) GetTableOfContentsByBookID(ctx context.Context, request models.RequestTOC) (*models.TableOfContents, error) {
	book, err := s.repo.FindByID(ctx, request.BookID)
	if err != nil {
		return nil, err
	}

	toc := &models.TableOfContents{
		BookID:    book.ID,
		BookTitle: book.Title,
		Author:    book.Owner,        // TODO: change to author name when Auth service will be implemented
		Tags:      make([]string, 0), // TODO: implement tags
	}

	var chapters []*entity2.Section
	chapters, err = s.repo.GetTOCSectionsFromChapters(ctx, request.BookID)
	if err != nil {
		return nil, err
	}

	var pages []*entity2.Section
	pages, err = s.repo.GetTOCSectionsFromPages(ctx, request.BookID)
	if err != nil {
		return nil, err
	}

	toc.Sections = joinSections(chapters, pages)

	return toc, nil
}

func (s *bookService) TogglePublic(ctx context.Context, request models.ToggleBookRequest) (*models.Book, error) {
	book, err := s.repo.FindByID(ctx, request.BookID)
	if err != nil {
		return nil, err
	}

	book.IsPublic = !book.IsPublic
	book, err = s.repo.Update(ctx, request.BookID, book)
	if err != nil {
		return nil, err
	}

	return adapters2.BookEntityToModel(book), nil
}

func joinSections(chapters, pages []*entity2.Section) []*models.Section {
	sections := make([]*models.Section, 0)
	pageSectionSet := make(map[string][]*models.Section)

	for _, page := range pages { // тут страницы отсортированы верно
		pageSection := adapters2.TocSectionEntityToModel(page)
		pageSectionSet[page.ParentID] = append(pageSectionSet[page.ParentID], pageSection)
	}

	for _, chapter := range chapters {
		var chapID = chapter.ID
		var chapIsPublic = chapter.IsPublic
		chapOrder := chapter.Order * 1_000 // предполагаем, что у нас не будет больше 1000 страниц в главе
		chapterSection := adapters2.TocSectionEntityToModel(chapter)
		chapterSection.Order = chapOrder
		sections = append(sections, chapterSection)

		for _, pageSection := range pageSectionSet[chapID] {
			pageSection.Order = chapOrder + pageSection.Order
			pageSection.IsPublic = chapIsPublic && pageSection.IsPublic
			sections = append(sections, pageSection)
		}
	}

	return sections
}
