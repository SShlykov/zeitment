package controllers

import (
	contextPkg "context"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics/localmetrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	mocks "github.com/SShlykov/zeitment/bookback/internal/tests/mocks/domain/services"
	loggerPkg "github.com/SShlykov/zeitment/bookback/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	cTestId     = "12b9b045-0845-462c-b372-0fca3180a6af"
	cTestIdPath = v1.BookEventsPath + "/id"
)

func init() {
	logger = loggerPkg.SetupLogger("debug")
	metrics = localmetrics.NewLocalMetrics(logger)
	context = contextPkg.Background()
	requestPageOptions = `{"options": {"page": 1, "page_size": 10}}`
	return
}

func TestChapterController_ListChapters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	listChapters := make([]*models.Chapter, 0)
	service.EXPECT().ListChapters(gomock.Any(), gomock.Any()).Return(listChapters, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath+"/list", strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.ListChapters(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
}

func TestChapterController_GetChapterByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: cTestId}
	service.EXPECT().GetChapterByID(gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, cTestIdPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(cTestId)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.GetChapterByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), cTestId))
}

func TestChapterController_CreateChapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: cTestId}
	service.EXPECT().CreateChapter(gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.CreateChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), cTestId))
}

func TestChapterController_UpdateChapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: cTestId}
	service.EXPECT().UpdateChapter(gomock.Any(), gomock.Any(), gomock.Any()).Return(chapter, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, cTestIdPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(cTestId)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.UpdateChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), cTestId))
}

func TestChapterController_DeleteChapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	service.EXPECT().DeleteChapter(gomock.Any(), gomock.Any()).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, cTestIdPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(cTestId)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.DeleteChapter(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Empty(t, rec.Body.String())
}

func TestChapterController_GetChaptersByBookID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockChapterService(ctrl)
	chapter := &models.Chapter{ID: cTestId}
	service.EXPECT().GetChapterByBookID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{chapter}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.ChaptersPath+"/book/"+cTestId, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.ChaptersPath + "/book/:id")
	c.SetParamNames("id")
	c.SetParamValues(cTestId)

	cc := NewChapterController(service, metrics, logger, context)
	err := cc.GetChapterByBookID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), cTestId))
}
