package adapters

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
)

func MapVariablesEntityToModel(variables []*entity.MapVariable) []*models.MapVariable {
	var result []*models.MapVariable
	for _, variable := range variables {
		result = append(result, MapVariableEntityToModel(variable))
	}
	return result
}

func MapVariableModelToEntity(variable *models.MapVariable) *entity.MapVariable {
	return &entity.MapVariable{
		ID:          variable.ID,
		CreatedAt:   variable.CreatedAt,
		UpdatedAt:   variable.UpdatedAt,
		BookID:      variable.BookID,
		ChapterID:   sql.NullString{Valid: variable.ChapterID.Valid, String: variable.ChapterID.Value},
		PageID:      sql.NullString{Valid: variable.PageID.Valid, String: variable.PageID.Value},
		ParagraphID: sql.NullString{Valid: variable.ParagraphID.Valid, String: variable.ParagraphID.Value},
		MapLink:     variable.MapLink,
		Lat:         variable.Lat,
		Lng:         variable.Lng,
		Zoom:        sql.NullInt64{Valid: variable.Zoom.Valid, Int64: variable.Zoom.Value},
		Date:        sql.NullString{Valid: variable.Date.Valid, String: variable.Date.Value},
		Description: sql.NullString{Valid: variable.Description.Valid, String: variable.Description.Value},
		Link:        sql.NullString{Valid: variable.Link.Valid, String: variable.Link.Value},
		LinkText:    sql.NullString{Valid: variable.LinkText.Valid, String: variable.LinkText.Value},
		LinkType:    sql.NullString{Valid: variable.LinkType.Valid, String: variable.LinkType.Value},
		LinkImage:   sql.NullString{Valid: variable.LinkImage.Valid, String: variable.LinkImage.Value},
		Image:       sql.NullString{Valid: variable.Image.Valid, String: variable.Image.Value},
	}
}

func MapVariableEntityToModel(variable *entity.MapVariable) *models.MapVariable {
	return &models.MapVariable{
		ID:          variable.ID,
		CreatedAt:   variable.CreatedAt,
		UpdatedAt:   variable.UpdatedAt,
		BookID:      variable.BookID,
		ChapterID:   types.Null[string]{Valid: variable.ChapterID.Valid, Value: variable.ChapterID.String},
		PageID:      types.Null[string]{Valid: variable.PageID.Valid, Value: variable.PageID.String},
		ParagraphID: types.Null[string]{Valid: variable.ParagraphID.Valid, Value: variable.ParagraphID.String},
		MapLink:     variable.MapLink,
		Lat:         variable.Lat,
		Lng:         variable.Lng,
		Zoom:        types.Null[int64]{Valid: variable.Zoom.Valid, Value: variable.Zoom.Int64},
		Date:        types.Null[string]{Valid: variable.Date.Valid, Value: variable.Date.String},
		Description: types.Null[string]{Valid: variable.Description.Valid, Value: variable.Description.String},
		Link:        types.Null[string]{Valid: variable.Link.Valid, Value: variable.Link.String},
		LinkText:    types.Null[string]{Valid: variable.LinkText.Valid, Value: variable.LinkText.String},
		LinkType:    types.Null[string]{Valid: variable.LinkType.Valid, Value: variable.LinkType.String},
		LinkImage:   types.Null[string]{Valid: variable.LinkImage.Valid, Value: variable.LinkImage.String},
		Image:       types.Null[string]{Valid: variable.Image.Valid, Value: variable.Image.String},
	}
}
