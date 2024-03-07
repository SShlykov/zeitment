package converter

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
	"time"
)

func ParagraphsEntityToModel(paragraphs []*entity.Paragraph) []*models.Paragraph {
	var result []*models.Paragraph
	for _, paragraph := range paragraphs {
		result = append(result, ParagraphEntityToModel(paragraph))
	}
	return result
}

func ParagraphEntityToModel(paragraph *entity.Paragraph) *models.Paragraph {
	return &models.Paragraph{
		ID:        paragraph.ID,
		CreatedAt: paragraph.CreatedAt,
		UpdatedAt: paragraph.UpdatedAt,
		DeletedAt: types.Null[time.Time]{Valid: paragraph.DeletedAt.Valid, Value: paragraph.DeletedAt.Time},
		Title:     paragraph.Title,
		Text:      paragraph.Text,
		Type:      paragraph.Type,
		IsPublic:  paragraph.IsPublic,
		PageID:    paragraph.PageID,
	}
}

func ParagraphModelToEntity(paragraph *models.Paragraph) *entity.Paragraph {
	return &entity.Paragraph{
		ID:        paragraph.ID,
		CreatedAt: paragraph.CreatedAt,
		UpdatedAt: paragraph.UpdatedAt,
		DeletedAt: sql.NullTime{Valid: paragraph.DeletedAt.Valid, Time: paragraph.DeletedAt.Value},
		Title:     paragraph.Title,
		Text:      paragraph.Text,
		Type:      paragraph.Type,
		IsPublic:  paragraph.IsPublic,
		PageID:    paragraph.PageID,
	}
}
