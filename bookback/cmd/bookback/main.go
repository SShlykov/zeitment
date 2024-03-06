package main

import (
	"flag"
	"fmt"
	pkg "github.com/SShlykov/zeitment/bookback/internal/bootstrap/app"
	"os"
)

var configPath string

// @title Book API
// @version 0.1
// @description Это API для работы с книгами
// @host localhost:7077
// @BasePath /api/v1
// @schemes http
// @produces application/json
// @consumes application/json
func main() {
	flag.StringVar(&configPath, "config", "./config/default.yml", "path to the configuration file")
	app, err := pkg.NewApp(configPath)
	if err != nil {
		fmt.Printf("failed to create app: %+v\n", err)
		os.Exit(2)
	}

	err = app.Run()
	if err != nil {
		fmt.Printf("failed to run app: %+v\n", err)
		os.Exit(2)
	}
}
