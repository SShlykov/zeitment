package book

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/config"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

// Service описывает сервис для работы с книгами.
type Service interface {
	CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	GetBookByID(ctx context.Context, id string) (*entity.Book, error)
	UpdateBook(ctx context.Context, id string, book *entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) (*entity.Book, error)
	ListBooks(ctx context.Context) ([]entity.Book, error)
}

type service struct {
	repo Repo
}

// NewService создает новый экземпляр Service.
func NewService(repo Repo) Service {
	return &service{repo}
}

func (s *service) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	var newBook *entity.Book
	id, err := s.repo.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	newBook, err = s.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return newBook, err
}

func (s *service) GetBookByID(ctx context.Context, id string) (*entity.Book, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateBook(ctx context.Context, id string, book *entity.Book) (*entity.Book, error) {
	if !s.isBookExisted(ctx, id) {
		return nil, config.ErrorNotFound
	}
	return s.repo.Update(ctx, id, book)
}

func (s *service) DeleteBook(ctx context.Context, id string) (*entity.Book, error) {
	if !s.isBookExisted(ctx, id) {
		return nil, config.ErrorNotFound
	}
	return s.repo.Delete(ctx, id)
}

func (s *service) ListBooks(ctx context.Context) ([]entity.Book, error) {
	return s.repo.List(ctx)
}

func (s *service) isBookExisted(ctx context.Context, id string) bool {
	_, err := s.GetBookByID(ctx, id)
	return err == nil
}
