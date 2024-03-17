package adapters

import (
	"github.com/SShlykov/zeitment/bookback/internal/domain/entity"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

func TocSectionsEntityToModel(entity []*entity.Section) []*models.Section {
	var sections []*models.Section
	for _, s := range entity {
		sections = append(sections, TocSectionEntityToModel(s))
	}
	return sections
}

func TocSectionEntityToModel(entity *entity.Section) *models.Section {
	return &models.Section{
		ID:       entity.ID,
		Title:    entity.Title,
		Order:    entity.Order,
		Level:    entity.Level,
		IsPublic: entity.IsPublic,
	}
}

func TocSectionsModelToEntity(model []*models.Section) []*entity.Section {
	var sections []*entity.Section
	for _, s := range model {
		sections = append(sections, TocSectionModelToEntity(s))
	}
	return sections
}

func TocSectionModelToEntity(model *models.Section) *entity.Section {
	return &entity.Section{
		ID:       model.ID,
		Title:    model.Title,
		Order:    model.Order,
		Level:    model.Level,
		IsPublic: model.IsPublic,
	}
}
