package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jansvabik/go-nodejs-performance/go/app"
	"github.com/jansvabik/go-nodejs-performance/go/config"
	"github.com/jansvabik/go-nodejs-performance/go/routers"
)

func main() {
	// create an app and store its configuration
	app.Create()
	err := config.Load(&app.State.Cfg)
	if err != nil {
		log.Fatalln("Failed to load the configuration")
		os.Exit(1)
	}

	// try to connect to the database
	client, err := app.DatabaseConnect()
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
		os.Exit(1)
	}
	app.State.MongoClient = client

	// init router
	r := routers.InitRoutes()

	// create server
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
