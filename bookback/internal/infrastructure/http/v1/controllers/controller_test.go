package controllers

import (
	ctx "context"
	v1 "github.com/SShlykov/zeitment/bookback/internal/infrastructure/http/v1"
	localMetrics "github.com/SShlykov/zeitment/bookback/internal/infrastructure/metrics"
	"log/slog"
)

var (
	logger             *slog.Logger
	metrics            localMetrics.Metrics
	context            ctx.Context
	requestPageOptions string
	id                 = "12b9b045-0845-462c-b372-0fca3180a6af"
	idPath             = v1.BooksPath + "/id"
)
