package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

func PagesEntityToModel(pages []*entity.Page) []*models.Page {
	var result []*models.Page
	for _, page := range pages {
		result = append(result, PageEntityToModel(page))
	}
	return result
}

func PageModelToEntity(m *models.Page) *entity.Page {
	return &entity.Page{
		ID:          m.ID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   sql.NullTime{Valid: m.DeletedAt.Valid, Time: m.DeletedAt.Value},
		Title:       m.Title,
		Text:        m.Text,
		Number:      m.Number,
		ChapterID:   m.ChapterID,
		IsPublic:    m.IsPublic,
		MapParamsID: sql.NullString{Valid: m.MapParamsID.Valid, String: m.MapParamsID.Value},
	}
}

func PageEntityToModel(e *entity.Page) *models.Page {
	return &models.Page{
		ID:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   types.Null[time.Time]{Valid: e.DeletedAt.Valid, Value: e.DeletedAt.Time},
		Title:       e.Title,
		Text:        e.Text,
		Number:      e.Number,
		ChapterID:   e.ChapterID,
		IsPublic:    e.IsPublic,
		MapParamsID: types.Null[string]{Valid: e.MapParamsID.Valid, Value: e.MapParamsID.String},
	}
}
