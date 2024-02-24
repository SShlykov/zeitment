package chapter

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
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
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (ch *service) CreateChapter(ctx context.Context, chapter *models.Chapter) (*models.Chapter, error) {
	id, err := ch.repo.Create(ctx, chapter)
	if err != nil {
		return nil, err
	}
	chapter.ID = id

	return chapter, err
}

func (ch *service) GetChapterByID(ctx context.Context, id string) (*models.Chapter, error) {
	return ch.repo.FindByID(ctx, id)
}

func (ch *service) UpdateChapter(ctx context.Context, id string, chapter *models.Chapter) (*models.Chapter, error) {
	return ch.repo.Update(ctx, id, chapter)
}

func (ch *service) DeleteChapter(ctx context.Context, id string) (*models.Chapter, error) {
	return ch.repo.Delete(ctx, id)
}

func (ch *service) ListChapters(ctx context.Context) ([]models.Chapter, error) {
	return ch.repo.List(ctx)
}

func (ch *service) GetChapterByBookID(ctx context.Context, bookID string) ([]models.Chapter, error) {
	return ch.repo.GetChapterByBookID(ctx, bookID)
}
