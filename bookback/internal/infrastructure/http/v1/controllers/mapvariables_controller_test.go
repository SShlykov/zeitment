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

func TestMapVariablesController_CreateMapVariable(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().CreateMapVariable(gomock.Any(), gomock.Any()).Return(mapVariable, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.MapVariablesPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.CreateMapVariable(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"created"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_DeleteMapVariable(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().DeleteMapVariable(gomock.Any(), gomock.Any()).Return(mapVariable, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, v1.MapVariablesPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.DeleteMapVariable(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"deleted"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_GetMapVariablesByPageID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().GetMapVariablesByPageID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.MapVariable{mapVariable}, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.MapVariablesPath+"/page/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetMapVariablesByPageID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_GetMapVariablesByChapterID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().GetMapVariablesByChapterID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.MapVariable{mapVariable}, nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		v1.MapVariablesPath+v1.ChapterSubPath+"/"+fixture.ID,
		strings.NewReader(fixture.RequestPageOptions),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetMapVariablesByChapterID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_GetMapVariablesByBookID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().GetMapVariablesByBookID(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.MapVariable{mapVariable}, nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		v1.MapVariablesPath+v1.BookSubPath+"/"+fixture.ID,
		strings.NewReader(fixture.RequestPageOptions),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetMapVariablesByBookID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_UpdateMapVariable(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().UpdateMapVariable(gomock.Any(), gomock.Any(), gomock.Any()).Return(mapVariable, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, v1.MapVariablesPath+"/"+fixture.ID, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.UpdateMapVariable(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"updated"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}

func TestMapVariablesController_GetMapVariableByID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MapVariablesPath)

	service := mocks.NewMockMapVariablesService(ctrl)
	mapVariable := &models.MapVariable{ID: fixture.ID}
	service.EXPECT().GetMapVariableByID(gomock.Any(), gomock.Any()).Return(mapVariable, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, v1.MapVariablesPath+"/"+fixture.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fixture.ID)

	cc := NewMapVariablesController(service, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetMapVariableByID(c)
	if err != nil {
		return
	}

	assert.Empty(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, strings.Contains(rec.Body.String(), `"status":"ok"`))
	assert.True(t, strings.Contains(rec.Body.String(), `"data":`))
	assert.True(t, strings.Contains(rec.Body.String(), fixture.ID))
}
