package chapter

import (
	"context"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
)

type Service interface {
	CreateChapter(ctx context.Context, chapter *entity.Chapter) (*entity.Chapter, error)
	GetChapterByID(ctx context.Context, id string) (*entity.Chapter, error)
	UpdateChapter(ctx context.Context, id string, chapter *entity.Chapter) (*entity.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (*entity.Chapter, error)
	ListChapters(ctx context.Context) ([]entity.Chapter, error)

	GetChapterByBookID(ctx context.Context, bookID string) ([]entity.Chapter, error)
}

type service struct {
	chapterRepo Repository
}

func NewService(chapterRepo Repository) Service {
	return &service{chapterRepo: chapterRepo}
}

func (ch *service) CreateChapter(ctx context.Context, chapter *entity.Chapter) (*entity.Chapter, error) {
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

func (ch *service) GetChapterByID(ctx context.Context, id string) (*entity.Chapter, error) {
	return ch.chapterRepo.FindByID(ctx, id)
}

func (ch *service) UpdateChapter(ctx context.Context, id string, chapter *entity.Chapter) (*entity.Chapter, error) {
	return ch.chapterRepo.Update(ctx, id, chapter)
}

func (ch *service) DeleteChapter(ctx context.Context, id string) (*entity.Chapter, error) {
	return ch.chapterRepo.Delete(ctx, id)
}

func (ch *service) ListChapters(ctx context.Context) ([]entity.Chapter, error) {
	return ch.chapterRepo.List(ctx)
}

func (ch *service) GetChapterByBookID(ctx context.Context, bookID string) ([]entity.Chapter, error) {
	return ch.chapterRepo.GetChapterByBookID(ctx, bookID)
}
