package endpoint

import (
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"net/http"
)

func RunServer(handler *Handler, logger loggerPkg.Logger) error {
	httpServer := &http.Server{
		ReadHeaderTimeout: handler.Timeout,
		ReadTimeout:       handler.Timeout,
		WriteTimeout:      handler.Timeout,
		IdleTimeout:       handler.IddleTimeout,
		Addr:              handler.Address,
		Handler:           handler.e,
	}

	logger.Info("HTTP server started")
	return httpServer.ListenAndServe()
}
