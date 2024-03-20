package main

import (
	"context"
	"fmt"
	userService "github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
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
	proxyAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	HTTPProxy(proxyAddr)
}

func HTTPProxy(proxyAddr string) {
	grpcGwMux := runtime.NewServeMux()

	//----------------------------------------------------------------
	// настройка подключений со стороны gRPC
	//----------------------------------------------------------------

	//grpc.WithPerRPCCredentials(&reqData{}),
	grpcUserConn, err := grpc.Dial(authServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Filed to connect to User service", err)
	}
	defer grpcUserConn.Close()

	err = userService.RegisterUserServiceHandler(context.Background(), grpcGwMux, grpcUserConn)
	if err != nil {
		log.Fatalln("Filed to start HTTP server", err)
	}

	//----------------------------------------------------------------
	//	Настройка маршрутов со стороны REST
	//----------------------------------------------------------------
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", grpcGwMux)
	mux.HandleFunc("/", helloWorld)

	fmt.Println("starting HTTP server at " + proxyAddr)
	log.Fatal(http.ListenAndServe(proxyAddr, mux))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, world!"))
}
