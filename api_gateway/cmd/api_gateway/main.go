package main

import (
	"context"
	"fmt"
	userService "github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	loggerPkg "github.com/SShlykov/zeitment/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var (
	authServiceURL = fmt.Sprintf("%s:%s", os.Getenv("AUTH_HOST"), os.Getenv("AUTH_PORT"))
)

func main() {
	logger := loggerPkg.SetupLogger("info")
	proxyAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	HTTPProxy(proxyAddr, logger)
}

func HTTPProxy(proxyAddr string, logger loggerPkg.Logger) {
	grpcGwMux := runtime.NewServeMux()

	//----------------------------------------------------------------
	// настройка подключений со стороны gRPC
	//----------------------------------------------------------------

	//grpc.WithPerRPCCredentials(&reqData{}),
	grpcUserConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure())
	if err != nil {
		logger.Error("Filed to connect to User service", loggerPkg.Err(err))
	}
	defer grpcUserConn.Close()

	err = userService.RegisterUserServiceHandler(context.Background(), grpcGwMux, grpcUserConn)
	if err != nil {
		logger.Error("Filed to register User service", loggerPkg.Err(err))
	}

	//----------------------------------------------------------------
	//	Настройка маршрутов со стороны REST
	//----------------------------------------------------------------
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", grpcGwMux)

	logger.Info("starting HTTP server at " + proxyAddr)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}
