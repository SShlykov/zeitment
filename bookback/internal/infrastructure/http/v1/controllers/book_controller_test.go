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

func TestBookController_GetTableOfContentsByBookID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	toc := &models.TableOfContents{BookID: fixture.ID}
	service.EXPECT().GetTableOfContentsByBookID(gomock.Any(), gomock.Any()).Return(toc, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath+v1.ToggleSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := bc.GetTableOfContentsByBookID(c)
	if err != nil {
		return
	}
	body := rec.Body.String()

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(body, `"status":"ok"`))
	assert.True(t, strings.Contains(body, `"data":`))
	assert.True(t, strings.Contains(body, `"book_id":"12b9b045-0845-462c-b372-0fca3180a6af"`))
}

func TestBookController_ListBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	listBooks := make([]*models.Book, 0)
	service.EXPECT().ListBooks(gomock.Any(), gomock.Any()).Return(listBooks, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath+v1.ListSubPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
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

	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: fixture.ID}
	service.EXPECT().GetBookByID(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fixture.IDPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
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
	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: "12b9b045-0845-462c-b372-0fca3180a6af"}
	service.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.BooksPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
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
	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: fixture.ID}
	service.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, fixture.IDPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
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
	fixture := NewTestFixture(v1.BooksPath)

	service := mocks.NewMockBookService(ctrl)
	book := &models.Book{ID: fixture.ID}
	service.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(book, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, fixture.IDPath, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	bc := NewBookController(service, fixture.Metrics, fixture.Logger, fixture.Context)
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
