package app

import (
	"context"
	"github.com/SShlykov/zeitment/auth/pkg/config"
	"github.com/SShlykov/zeitment/bookback/pkg/postgres"
)

type App struct {
	configPath string
	logger     loggerPkg.Logger
	config     *config.Config
	db         postgres.Client

	ctx      context.Context
	closeCtx func()
}
