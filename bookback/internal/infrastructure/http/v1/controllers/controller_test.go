package controllers

import (
	ctx "context"
	localMetrics "github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics/localmetrics"
	loggerPkg "github.com/SShlykov/zeitment/logger"
)

type TestFixture struct {
	Logger             loggerPkg.Logger
	Metrics            localMetrics.Metrics
	Context            ctx.Context
	RequestPageOptions string
	ID                 string
	IDPath             string
}

func NewTestFixture(basePath string) *TestFixture {
	fixture := &TestFixture{}
	fixture.Logger = loggerPkg.SetupLogger("debug")
	fixture.Metrics = localmetrics.NewLocalMetrics(fixture.Logger)
	fixture.Context = ctx.Background()
	fixture.RequestPageOptions = `{"options": {"page": 1, "page_size": 10}}`
	fixture.ID = "12b9b045-0845-462c-b372-0fca3180a6af"
	fixture.IDPath = basePath + "/id"

	return fixture
}
