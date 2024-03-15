package main

import (
	"flag"
	"fmt"
	pkg "github.com/SShlykov/zeitment/auth/internal/bootstrap/app"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config", "path to the configuration files")

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
