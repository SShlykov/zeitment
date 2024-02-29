package chapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/services/book"
)

type Service interface {
	CreateChapter(ctx context.Context, chapter *models.Chapter) (*models.Chapter, error)
	GetChapterByID(ctx context.Context, id string) (*models.Chapter, error)
	UpdateChapter(ctx context.Context, id string, chapter *models.Chapter) (*models.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (*models.Chapter, error)
	ListChapters(ctx context.Context) ([]models.Chapter, error)

	GetChapterByBookID(ctx context.Context, bookID string) ([]models.Chapter, error)
}

type service struct {
	chapterRepo Repository
	bookRepo    book.Repository
}

func NewService(chapterRepo Repository, bookRepo book.Repository) Service {
	return &service{chapterRepo: chapterRepo, bookRepo: bookRepo}
}

func (ch *service) CreateChapter(ctx context.Context, chapter *models.Chapter) (*models.Chapter, error) {
	if !ch.isBookExisted(ctx, chapter.BookID) {
		return nil, errors.New("book not found")
	}

	id, err := ch.chapterRepo.Create(ctx, chapter)
	if err != nil {
		return nil, err
	}

	chapter, err = ch.GetChapterByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return chapter, nil
}

func (ch *service) GetChapterByID(ctx context.Context, id string) (*models.Chapter, error) {
	return ch.chapterRepo.FindByID(ctx, id)
}

func (ch *service) UpdateChapter(ctx context.Context, id string, chapter *models.Chapter) (*models.Chapter, error) {
	if !ch.isBookExisted(ctx, chapter.BookID) {
		return nil, errors.New("book not found")
	}
	return ch.chapterRepo.Update(ctx, id, chapter)
}

func (ch *service) DeleteChapter(ctx context.Context, id string) (*models.Chapter, error) {
	return ch.chapterRepo.Delete(ctx, id)
}

func (ch *service) ListChapters(ctx context.Context) ([]models.Chapter, error) {
	return ch.chapterRepo.List(ctx)
}

func (ch *service) GetChapterByBookID(ctx context.Context, bookID string) ([]models.Chapter, error) {
	if !ch.isBookExisted(ctx, bookID) {
		return nil, errors.New("book not found")
	}
	return ch.chapterRepo.GetChapterByBookID(ctx, bookID)
}

func (ch *service) isBookExisted(ctx context.Context, id string) bool {
	depBook, err := ch.bookRepo.FindByID(ctx, id)
	if err != nil || depBook == nil {
		return false
	}
	return true
}
