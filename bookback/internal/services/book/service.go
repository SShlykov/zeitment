package book

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

// Service описывает сервис для работы с книгами.
type Service interface {
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	UpdateBook(ctx context.Context, id string, book *models.Book) (*models.Book, error)
	DeleteBook(ctx context.Context, id string) (*models.Book, error)
	ListBooks(ctx context.Context) ([]models.Book, error)
}

type service struct {
	repo Repository
}

// NewService создает новый экземпляр Service.
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	id, err := s.repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	book.ID = id

	return book, err
}

func (s *service) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateBook(ctx context.Context, id string, book *models.Book) (*models.Book, error) {
	return s.repo.Update(ctx, id, book)
}

func (s *service) DeleteBook(ctx context.Context, id string) (*models.Book, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) ListBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.List(ctx)
}
