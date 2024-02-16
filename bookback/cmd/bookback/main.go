package main

import (
	"flag"
	"fmt"
	appPkg "github.com/SShlykov/zeitment/bookback/internal/pkg/app"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./config/default.yml", "path to the configuration file")
}

func main() {
	app, err := appPkg.NewApp(configPath)
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
