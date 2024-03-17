package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/adapters"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/postgres/dbutils"
)

//go:generate mockgen -destination=../../../tests/mocks/domain/services/chapter_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services ChapterService
type ChapterService interface {
	CreateChapter(ctx context.Context, request models.CreateChapterRequest) (*models.Chapter, error)
	GetChapterByID(ctx context.Context, id string) (*models.Chapter, error)
	UpdateChapter(ctx context.Context, id string, request models.UpdateChapterRequest) (*models.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (*models.Chapter, error)
	ListChapters(ctx context.Context, request models.RequestChapter) ([]*models.Chapter, error)

	GetChapterByBookID(ctx context.Context, bookID string, request models.RequestChapter) ([]*models.Chapter, error)

	TogglePublic(ctx context.Context, request models.ToggleChapterRequest) (*models.Chapter, error)
}

type chapterService struct {
	chapterRepo SimpleRepo[*entity.Chapter]
}

func NewChapterService(chapterRepo SimpleRepo[*entity.Chapter]) ChapterService {
	return &chapterService{chapterRepo: chapterRepo}
}

func (ch *chapterService) TogglePublic(ctx context.Context, request models.ToggleChapterRequest) (*models.Chapter, error) {
	chapter, err := ch.chapterRepo.FindByID(ctx, request.ChapterID)
	if err != nil {
		return nil, err
	}
	chapter.IsPublic = !chapter.IsPublic

	var updated *entity.Chapter
	updated, err = ch.chapterRepo.Update(ctx, request.ChapterID, chapter)
	if err != nil {
		return nil, err
	}

	return adapters.ChapterEntityToModel(updated), nil
}

func (ch *chapterService) CreateChapter(ctx context.Context, request models.CreateChapterRequest) (*models.Chapter, error) {
	chapter := adapters.ChapterModelToEntity(request.Chapter)

	id, err := ch.chapterRepo.Create(ctx, chapter)
	if err != nil {
		return nil, err
	}

	var newChapter *models.Chapter
	newChapter, err = ch.GetChapterByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return newChapter, nil
}

func (ch *chapterService) GetChapterByID(ctx context.Context, id string) (*models.Chapter, error) {
	chapter, err := ch.chapterRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return adapters.ChapterEntityToModel(chapter), nil
}

func (ch *chapterService) UpdateChapter(ctx context.Context, id string, request models.UpdateChapterRequest) (*models.Chapter, error) {
	chapter := adapters.ChapterModelToEntity(request.Chapter)

	updatedChapter, err := ch.chapterRepo.Update(ctx, id, chapter)
	if err != nil {
		return nil, err
	}

	return adapters.ChapterEntityToModel(updatedChapter), err
}

func (ch *chapterService) DeleteChapter(ctx context.Context, id string) (*models.Chapter, error) {
	chapter, err := ch.GetChapterByID(ctx, id)
	if err != nil {
		return nil, errors.Join(errors.New("chapter not found"), err)
	}
	err = ch.chapterRepo.HardDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return chapter, err
}

func (ch *chapterService) ListChapters(ctx context.Context, request models.RequestChapter) ([]*models.Chapter, error) {
	options := dbutils.NewPagination(&request.Options)

	chapters, err := ch.chapterRepo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return adapters.ChaptersEntityToModel(chapters), nil
}

func (ch *chapterService) GetChapterByBookID(ctx context.Context, bookID string, request models.RequestChapter) ([]*models.Chapter, error) {
	options := dbutils.NewQueryOptions(
		dbutils.NewFilter("book_id", bookID),
		dbutils.NewPagination(&request.Options),
	)

	chapters, err := ch.chapterRepo.FindByKV(ctx, options)
	if err != nil {
		return nil, err
	}
	return adapters.ChaptersEntityToModel(chapters), nil
}
