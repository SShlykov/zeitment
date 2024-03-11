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

func TestBookEventController_GetBookEventByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().GetBookEventByID(gomock.Any(), fixture.ID).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fixture.IDPath, nil)
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_CreateBookEvent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().CreateBookEvent(gomock.Any(), gomock.Any()).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.CreateBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_UpdateBookEvent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().UpdateBookEvent(gomock.Any(), fixture.ID, gomock.Any()).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, fixture.IDPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.UpdateBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_DeleteBookEvent(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().DeleteBookEvent(gomock.Any(), fixture.ID).Return(bookEvent, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, fixture.IDPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + v1.IDVar)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.DeleteBookEvent(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_GetBookEventsByBookID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().GetBookEventsByBookID(gomock.Any(), fixture.ID, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		v1.BookEventsPath+v1.BookSubPath+"/"+fixture.ID,
		strings.NewReader(fixture.RequestPageOptions),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/book/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventsByBookID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_GetBookEventsByChapterID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().GetBookEventsByChapterID(gomock.Any(), fixture.ID, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		v1.BookEventsPath+v1.ChapterSubPath+"/"+fixture.ID,
		strings.NewReader(fixture.RequestPageOptions),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/chapter/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventsByChapterID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_GetBookEventsByPageID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().GetBookEventsByPageID(gomock.Any(), fixture.ID, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/page/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/page/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventsByPageID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_GetBookEventsByParagraphID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)
	bookEvent := &models.BookEvent{ID: fixture.ID}
	service.EXPECT().GetBookEventsByParagraphID(gomock.Any(), fixture.ID, gomock.Any()).Return([]*models.BookEvent{bookEvent}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/paragraph/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/paragraph/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventsByParagraphID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestBookEventController_GetBookEventsByParagraphID_BindError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.BookEventsPath)

	service := mocks.NewMockBookEventsService(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BookEventsPath+"/paragraph/"+fixture.ID, strings.NewReader("invalid json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(v1.BookEventsPath + "/paragraph/:id")
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookEventController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetBookEventsByParagraphID(c)

	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
}
