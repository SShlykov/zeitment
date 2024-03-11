package controllers

import (
	contextPkg "context"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics/localmetrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	loggerPkg "github.com/SShlykov/zeitment/bookback/pkg/logger"
	mocks "github.com/SShlykov/zeitment/bookback/tests/mocks/domain/services"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	beTestId     = "12b9b045-0845-462c-b372-0fca3180a6af"
	beTestIdPath = v1.BookEventsPath + "/id"
)

func init() {
	logger = loggerPkg.SetupLogger("debug")
	metrics = localmetrics.NewLocalMetrics(logger)
	context = contextPkg.Background()
	requestPageOptions = `{"options": {"page": 1, "page_size": 10}}`
	return
}

func TestBookEventController_GetBookEventByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().GetBookEventByID(gomock.Any(), beTestId).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, beTestIdPath, nil)
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(beTestId)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_CreateBookEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().CreateBookEvent(gomock.Any(), gomock.Any()).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.CreateBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_UpdateBookEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().UpdateBookEvent(gomock.Any(), beTestId, gomock.Any()).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, beTestIdPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(beTestId)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.UpdateBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_DeleteBookEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().DeleteBookEvent(gomock.Any(), beTestId).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, beTestIdPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(beTestId)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.DeleteBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_GetBookEventsByBookID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().GetBookEventsByBookID(gomock.Any(), beTestId, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/book/"+id, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/book/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventsByBookID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_GetBookEventsByChapterID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().GetBookEventsByChapterID(gomock.Any(), beTestId, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/chapter/"+id, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/chapter/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventsByChapterID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_GetBookEventsByPageID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().GetBookEventsByPageID(gomock.Any(), beTestId, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/page/"+id, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/page/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventsByPageID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_GetBookEventsByParagraphID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: beTestId}
	service.EXPECT().GetBookEventsByParagraphID(gomock.Any(), beTestId, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/paragraph/"+id, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/paragraph/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventsByParagraphID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), beTestId))
}

func TestBookEventController_GetBookEventsByParagraphID_BindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookEventsService(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/paragraph/"+id, strings.NewReader("invalid json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/paragraph/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookEventController(service, metrics, logger, context)
	err := bc.GetBookEventsByParagraphID(c)

	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
}
