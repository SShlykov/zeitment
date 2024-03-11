package controllers

import (
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthController_GetHealthCheck(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.HealthPath)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, v1.HealthPath, strings.NewReader(fixture.RequestPageOptions))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	hc := NewHealthController(fixture.Metrics, fixture.Logger, fixture.Context)
	err := hc.GetHealthCheck(c)

	assert.Empty(t, err)
	assert.True(t, strings.Contains(rec.Body.String(), `"healthy"`))
}
