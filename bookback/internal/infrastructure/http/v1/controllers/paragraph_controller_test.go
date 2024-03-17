package controllers

import (
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	mocks "github.com/SShlykov/zeitment/bookback/tests/mocks/domain/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParagraphController_TogglePublic(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraph := &models.Paragraph{ID: fixture.ID}
	service.EXPECT().TogglePublic(gomock.Any(), gomock.Any()).Return(paragraph, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ParagraphsPath+v1.ToggleSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.TogglePublic(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestParagraphController_CreateParagraph(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraph := &models.Paragraph{ID: fixture.ID}
	service.EXPECT().CreateParagraph(gomock.Any(), gomock.Any()).Return(paragraph, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ParagraphsPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.CreateParagraph(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestParagraphController_DeleteParagraph(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraph := &models.Paragraph{ID: fixture.ID}
	service.EXPECT().DeleteParagraph(gomock.Any(), gomock.Any()).Return(paragraph, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, v1.ParagraphsPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.DeleteParagraph(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestParagraphController_GetParagraphByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraph := &models.Paragraph{ID: fixture.ID}
	service.EXPECT().GetParagraphByID(gomock.Any(), gomock.Any()).Return(paragraph, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, v1.ParagraphsPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetParagraphByID(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestParagraphController_GetParagraphsByPageID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraphs := []*models.Paragraph{{ID: fixture.ID}}
	service.EXPECT().GetParagraphsByPageID(gomock.Any(), gomock.Any(), gomock.Any()).Return(paragraphs, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, v1.ParagraphsPath+v1.ChapterSubPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetParagraphsByPageID(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestParagraphController_UpdateParagraph(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ParagraphsPath)

	service := mocks.NewMockParagraphService(ctrl)
	paragraph := &models.Paragraph{ID: fixture.ID}
	service.EXPECT().UpdateParagraph(gomock.Any(), gomock.Any(), gomock.Any()).Return(paragraph, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, v1.ParagraphsPath+"/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewParagraphController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.UpdateParagraph(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}
