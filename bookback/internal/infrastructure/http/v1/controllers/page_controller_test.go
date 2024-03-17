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

func TestPageController_TogglePublic(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	page := &models.Page{ID: fixture.ID}
	service.EXPECT().TogglePublic(gomock.Any(), gomock.Any()).Return(page, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.PagesPath+v1.ToggleSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.TogglePublic(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestPageController_ListPages(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	service.EXPECT().ListPages(gomock.Any(), gomock.Any()).Return([]*models.Page{}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.PagesPath+v1.ListSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.ListPages(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":[]`))
}

func TestPageController_CreatePage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	page := &models.Page{ID: fixture.ID}
	service.EXPECT().CreatePage(gomock.Any(), gomock.Any()).Return(page, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.PagesPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.CreatePage(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestPageController_GetPageByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	page := &models.Page{ID: fixture.ID}
	service.EXPECT().GetPageByID(gomock.Any(), gomock.Any()).Return(page, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, v1.PagesPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetPageByID(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestPageController_UpdatePage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	page := &models.Page{ID: fixture.ID}
	service.EXPECT().UpdatePage(gomock.Any(), gomock.Any(), gomock.Any()).Return(page, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, v1.PagesPath+"/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.UpdatePage(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestPageController_DeletePage(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	page := &models.Page{ID: fixture.ID}
	service.EXPECT().DeletePage(gomock.Any(), gomock.Any()).Return(page, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, v1.PagesPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.DeletePage(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestPageController_GetPagesByChapterID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.PagesPath)

	service := mocks.NewMockPageService(ctrl)
	service.EXPECT().GetPagesByChapterID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Page{}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.PagesPath+v1.ChapterSubPath+"/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewPageController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetPagesByChapterID(c)

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":[]`))
}
