package app

import (
	"github.com/jansvabik/go-nodejs-performance/go/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// App is a structure of app data
type App struct {
	MongoClient *mongo.Client
	Cfg         config.Config
}

// State contains the app data
var State App

// Create creates new app data structure
func Create() {
	State = App{}
}
