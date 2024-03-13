package app

import (
	"context"
	"github.com/SShlykov/zeitment/auth/pkg/config"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/postgres"
)

type App struct {
	configPath string
	logger     loggerPkg.Logger
	config     *config.Config
	db         postgres.Client

	ctx      context.Context
	closeCtx func()
}
