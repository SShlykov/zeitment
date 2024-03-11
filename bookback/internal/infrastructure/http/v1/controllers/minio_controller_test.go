package controllers

import (
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	"github.com/SShlykov/zeitment/bookback/internal/models"
	"github.com/SShlykov/zeitment/bookback/tests/mocks/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMinioController_GetMinioObject(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	fixture := NewTestFixture(v1.MinioPath)
	bucket, object := "test", "xxt.jpg"

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, v1.MinioPath+"/get-object/:bucket/:object", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("bucket", "object")
	c.SetParamValues(bucket, object)

	minioUseCase := mock_usecase.NewMockMinioUseCase(ctrl)
	resp := &models.MinioResp{ContentType: "image/jpeg", Content: []byte("test content"), Name: object}
	minioUseCase.EXPECT().GetMinioObject(gomock.Any(), gomock.Any()).Return(resp, nil)

	cc := NewMinioController(minioUseCase, fixture.Metrics, fixture.Logger, fixture.Context)
	err := cc.GetMinioObject(c)

	assert.Empty(t, err)
}
