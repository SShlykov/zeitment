package controllers

import (
	"context"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics/localmetrics"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/pkg/logger"
	mocks "github.com/SShlykov/zeitment/bookback/tests/mocks/domain/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	lggr        *slog.Logger
	mtrcs       metrics.Metrics
	ctx         context.Context
	RequestBook string
	id          = "12b9b045-0845-462c-b372-0fca3180a6af"
	idPath      = v1.BooksPath + "/id"
)

func init() {
	lggr = logger.SetupLogger("debug")
	mtrcs = localmetrics.NewLocalMetrics(lggr)
	ctx = context.Background()
	RequestBook = `{"options": {"page": 1, "page_size": 10}}`
	return
}

func TestBookController_ListBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	listBooks := make([]*models.Book, 0)
	service.EXPECT().ListBooks(gomock.Any(), gomock.Any()).Return(listBooks, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath+"/list", strings.NewReader(RequestBook))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, mtrcs, lggr, ctx)
	err := bc.ListBooks(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":[]`))
}

func TestBookController_GetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: id}
	service.EXPECT().GetBookByID(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, idPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookController(service, mtrcs, lggr, ctx)
	err := bc.GetBookByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), `"id":"12b9b045-0845-462c-b372-0fca3180a6af"`))
}

func TestBookController_CreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: "12b9b045-0845-462c-b372-0fca3180a6af"}
	service.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath, strings.NewReader(RequestBook))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, mtrcs, lggr, ctx)
	err := bc.CreateBook(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), `"id":"12b9b045-0845-462c-b372-0fca3180a6af"`))
}

func TestBookController_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: id}
	service.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, idPath, strings.NewReader(RequestBook))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookController(service, mtrcs, lggr, ctx)
	err := bc.UpdateBook(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), `"id":"12b9b045-0845-462c-b372-0fca3180a6af"`))
}

func TestBookController_DeleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: id}
	service.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, idPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookController(service, mtrcs, lggr, ctx)
	err := bc.DeleteBook(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), `"id":"12b9b045-0845-462c-b372-0fca3180a6af"`))
}
