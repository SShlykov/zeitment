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

func TestChapterController_ListChapters(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	listChapters := make([]*models.Chapter, 0)
	service.EXPECT().ListChapters(gomock.Any(), gomock.Any()).Return(listChapters, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath+v1.ListSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.ListChapters(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
}

func TestChapterController_GetChapterByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: fixture.ID}
	service.EXPECT().GetChapterByID(gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fixture.IDPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetChapterByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestChapterController_CreateChapter(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: fixture.ID}
	service.EXPECT().CreateChapter(gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.CreateChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestChapterController_UpdateChapter(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: fixture.ID}
	service.EXPECT().UpdateChapter(gomock.Any(), gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, fixture.IDPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.UpdateChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestChapterController_DeleteChapter(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: fixture.ID}
	service.EXPECT().DeleteChapter(gomock.Any(), fixture.ID).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, fixture.IDPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.DeleteChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestChapterController_GetChaptersByBookID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.ChaptersPath)

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: fixture.ID}
	service.EXPECT().GetChapterByBookID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{chapter}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath+v1.BookSubPath+"/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + "/book/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewChapterController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetChapterByBookID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}
