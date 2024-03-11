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

func init() {
	logger = loggerPkg.SetupLogger("debug")
	metrics = localmetrics.NewLocalMetrics(logger)
	context = contextPkg.Background()
	requestPageOptions = `{"options": {"page": 1, "page_size": 10}}`
	return
}

func TestBookController_ListBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockBookService(ctrl)
	listBooks := make([]*models.Book, 0)
	service.EXPECT().ListBooks(gomock.Any(), gomock.Any()).Return(listBooks, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath+"/list", strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, metrics, logger, context)
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

	bc := NewBookController(service, metrics, logger, context)
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
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, metrics, logger, context)
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
	req := httptest.NewRequest(http.MethodPut, idPath, strings.NewReader(requestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	bc := NewBookController(service, metrics, logger, context)
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

	bc := NewBookController(service, metrics, logger, context)
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
