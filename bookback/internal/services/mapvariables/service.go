package mapvariables

import (
	"context"
	"github.com/SShlykov/zeitment/bookback/internal/models"
)

type Service interface {
	CreateMapVariable(ctx context.Context, variable *models.MapVariable) (*models.MapVariable, error)
	GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error)
	UpdateMapVariable(ctx context.Context, id string, variable *models.MapVariable) (*models.MapVariable, error)
	DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error)
	GetMapVariablesByBookID(ctx context.Context, mapID string) ([]models.MapVariable, error)
	GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]models.MapVariable, error)
	GetMapVariablesByPageID(ctx context.Context, pageID string) ([]models.MapVariable, error)
	GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]models.MapVariable, error)

	GetMapVariablesByMapLinkAndBookID(ctx context.Context, mapLink, bookID string) ([]models.MapVariable, error)
	GetMapVariablesByMapLinkAndChapterID(ctx context.Context, mapLink, chapterID string) ([]models.MapVariable, error)
	GetMapVariablesByMapLinkAndPageID(ctx context.Context, mapLink, pageID string) ([]models.MapVariable, error)
	GetMapVariablesByMapLinkAndParagraphID(ctx context.Context, mapLink, paragraphID string) ([]models.MapVariable, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateMapVariable(ctx context.Context, variable *models.MapVariable) (*models.MapVariable, error) {
	id, err := s.repo.Create(ctx, variable)
	if err != nil {
		return nil, err
	}
	variable.ID = id

	return variable, err
}

func (s *service) GetMapVariableByID(ctx context.Context, id string) (*models.MapVariable, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateMapVariable(ctx context.Context, id string, variable *models.MapVariable) (*models.MapVariable, error) {
	return s.repo.Update(ctx, id, variable)
}

func (s *service) DeleteMapVariable(ctx context.Context, id string) (*models.MapVariable, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetMapVariablesByBookID(ctx context.Context, mapID string) ([]models.MapVariable, error) {
	return s.repo.GetByBookID(ctx, mapID)
}

func (s *service) GetMapVariablesByChapterID(ctx context.Context, chapterID string) ([]models.MapVariable, error) {
	return s.repo.GetByChapterID(ctx, chapterID)
}

func (s *service) GetMapVariablesByPageID(ctx context.Context, pageID string) ([]models.MapVariable, error) {
	return s.repo.GetByPageID(ctx, pageID)
}

func (s *service) GetMapVariablesByParagraphID(ctx context.Context, paragraphID string) ([]models.MapVariable, error) {
	return s.repo.GetByParagraphID(ctx, paragraphID)
}

func (s *service) GetMapVariablesByMapLinkAndBookID(ctx context.Context, mapLink, bookID string) ([]models.MapVariable, error) {
	return s.repo.GetByMapLinkAndBookID(ctx, mapLink, bookID)
}

func (s *service) GetMapVariablesByMapLinkAndChapterID(ctx context.Context, mapLink, chapterID string) ([]models.MapVariable, error) {
	return s.repo.GetByMapLinkAndChapterID(ctx, mapLink, chapterID)
}

func (s *service) GetMapVariablesByMapLinkAndPageID(ctx context.Context, mapLink, pageID string) ([]models.MapVariable, error) {
	return s.repo.GetByMapLinkAndPageID(ctx, mapLink, pageID)
}

func (s *service) GetMapVariablesByMapLinkAndParagraphID(ctx context.Context, mapLink, paragraphID string) ([]models.MapVariable, error) {
	return s.repo.GetByMapLinkAndParagraphID(ctx, mapLink, paragraphID)
}
