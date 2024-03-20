package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

func ChaptersEntityToModel(chapters []*entity.Chapter) []*models.Chapter {
	var result []*models.Chapter
	for _, chapter := range chapters {
		result = append(result, ChapterEntityToModel(chapter))
	}
	return result
}

func ChapterEntityToModel(chapter *entity.Chapter) *models.Chapter {
	return &models.Chapter{
		ID:          chapter.ID,
		CreatedAt:   chapter.CreatedAt,
		UpdatedAt:   chapter.UpdatedAt,
		DeletedAt:   types.Null[time.Time]{Valid: chapter.DeletedAt.Valid, Value: chapter.DeletedAt.Time},
		Title:       chapter.Title,
		Number:      chapter.Number,
		Text:        chapter.Text,
		BookID:      chapter.BookID,
		IsPublic:    chapter.IsPublic,
		MapLink:     types.Null[string]{Valid: chapter.MapLink.Valid, Value: chapter.MapLink.String},
		MapParamsID: types.Null[string]{Valid: chapter.MapLink.Valid, Value: chapter.MapLink.String},
	}
}

func ChapterModelToEntity(chapter *models.Chapter) *entity.Chapter {
	return &entity.Chapter{
		ID:          chapter.ID,
		CreatedAt:   chapter.CreatedAt,
		UpdatedAt:   chapter.UpdatedAt,
		DeletedAt:   sql.NullTime{Valid: chapter.DeletedAt.Valid, Time: chapter.DeletedAt.Value},
		Title:       chapter.Title,
		Number:      chapter.Number,
		Text:        chapter.Text,
		BookID:      chapter.BookID,
		IsPublic:    chapter.IsPublic,
		MapLink:     sql.NullString{Valid: chapter.MapLink.Valid, String: chapter.MapLink.Value},
		MapParamsID: sql.NullString{Valid: chapter.MapLink.Valid, String: chapter.MapLink.Value},
	}
}
