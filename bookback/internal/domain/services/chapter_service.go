package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/converter"
)

type ChapterService interface {
	CreateChapter(ctx context.Context, request models.CreateChapterRequest) (*models.Chapter, error)
	GetChapterByID(ctx context.Context, id string) (*models.Chapter, error)
	UpdateChapter(ctx context.Context, id string, request models.UpdateChapterRequest) (*models.Chapter, error)
	DeleteChapter(ctx context.Context, id string) (*models.Chapter, error)
	ListChapters(ctx context.Context, limit uint64, offset uint64) ([]*models.Chapter, error)

	GetChapterByBookID(ctx context.Context, bookID string) ([]*models.Chapter, error)
}

type chapterService struct {
	chapterRepo SimpleRepo[*entity.Chapter]
}

func NewChapterService(chapterRepo SimpleRepo[*entity.Chapter]) ChapterService {
	return &chapterService{chapterRepo: chapterRepo}
}

func (ch *chapterService) CreateChapter(ctx context.Context, request models.CreateChapterRequest) (*models.Chapter, error) {
	chapter := converter.ChapterModelToEntity(request.Chapter)

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

	return converter.ChapterEntityToModel(chapter), nil
}

func (ch *chapterService) UpdateChapter(ctx context.Context, id string, request models.UpdateChapterRequest) (*models.Chapter, error) {
	chapter := converter.ChapterModelToEntity(request.Chapter)

	updatedChapter, err := ch.chapterRepo.Update(ctx, id, chapter)
	if err != nil {
		return nil, err
	}

	return converter.ChapterEntityToModel(updatedChapter), err
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

func (ch *chapterService) ListChapters(ctx context.Context, limit uint64, offset uint64) ([]*models.Chapter, error) {
	chapters, err := ch.chapterRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return converter.ChaptersEntityToModel(chapters), nil
}

func (ch *chapterService) GetChapterByBookID(ctx context.Context, bookID string) ([]*models.Chapter, error) {
	chapters, err := ch.chapterRepo.FindByKV(ctx, "book_id", bookID)
	if err != nil {
		return nil, err
	}
	return converter.ChaptersEntityToModel(chapters), nil
}
