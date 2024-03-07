package converter

import (
	"database/sql"
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/internal/models/types"
)

func BookEventsEntityToModel(bookevents []*entity.BookEvent) []*models.BookEvent {
	bookEventModels := make([]*models.BookEvent, 0)
	for _, bookevent := range bookevents {
		bookEventModel := BookEventEntityToModel(bookevent)
		bookEventModels = append(bookEventModels, bookEventModel)
	}
	return bookEventModels
}

func BookEventEntityToModel(bookevent *entity.BookEvent) *models.BookEvent {
	return &models.BookEvent{
		ID:          bookevent.ID,
		CreatedAt:   bookevent.CreatedAt,
		UpdatedAt:   bookevent.UpdatedAt,
		BookID:      bookevent.BookID,
		ChapterID:   types.Null[string]{Valid: bookevent.ChapterID.Valid, Value: bookevent.ChapterID.String},
		PageID:      types.Null[string]{Valid: bookevent.PageID.Valid, Value: bookevent.PageID.String},
		ParagraphID: types.Null[string]{Valid: bookevent.ParagraphID.Valid, Value: bookevent.ParagraphID.String},
		EventType:   types.Null[string]{Valid: bookevent.EventType.Valid, Value: bookevent.EventType.String},
		IsPublic:    bookevent.IsPublic,
		Key:         bookevent.Key,
		Value:       bookevent.Value,
		Link:        types.Null[string]{Valid: bookevent.Link.Valid, Value: bookevent.Link.String},
		LinkText:    types.Null[string]{Valid: bookevent.LinkText.Valid, Value: bookevent.LinkText.String},
		LinkType:    types.Null[string]{Valid: bookevent.LinkType.Valid, Value: bookevent.LinkType.String},
		LinkImage:   types.Null[string]{Valid: bookevent.LinkImage.Valid, Value: bookevent.LinkImage.String},
		Description: types.Null[string]{Valid: bookevent.Description.Valid, Value: bookevent.Description.String},
	}
}

func BookEventModelToEntity(bookeventModel *models.BookEvent) *entity.BookEvent {
	return &entity.BookEvent{
		ID:          bookeventModel.ID,
		CreatedAt:   bookeventModel.CreatedAt,
		UpdatedAt:   bookeventModel.UpdatedAt,
		BookID:      bookeventModel.BookID,
		ChapterID:   sql.NullString{Valid: bookeventModel.ChapterID.Valid, String: bookeventModel.ChapterID.Value},
		PageID:      sql.NullString{Valid: bookeventModel.PageID.Valid, String: bookeventModel.PageID.Value},
		ParagraphID: sql.NullString{Valid: bookeventModel.ParagraphID.Valid, String: bookeventModel.ParagraphID.Value},
		EventType:   sql.NullString{Valid: bookeventModel.EventType.Valid, String: bookeventModel.EventType.Value},
		IsPublic:    bookeventModel.IsPublic,
		Key:         bookeventModel.Key,
		Value:       bookeventModel.Value,
		Link:        sql.NullString{Valid: bookeventModel.Link.Valid, String: bookeventModel.Link.Value},
		LinkText:    sql.NullString{Valid: bookeventModel.LinkText.Valid, String: bookeventModel.LinkText.Value},
		LinkType:    sql.NullString{Valid: bookeventModel.LinkType.Valid, String: bookeventModel.LinkType.Value},
		LinkImage:   sql.NullString{Valid: bookeventModel.LinkImage.Valid, String: bookeventModel.LinkImage.Value},
		Description: sql.NullString{Valid: bookeventModel.Description.Valid, String: bookeventModel.Description.Value},
	}
}
