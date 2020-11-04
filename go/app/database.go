package app

import (
	"context"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseConnect creates the database connection
func DatabaseConnect() (*mongo.Client, error) {
	// try to connect to the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(
		"mongodb://" +
			State.Cfg.Database.User +
			":" +
			url.QueryEscape(State.Cfg.Database.Password) +
			"@" +
			State.Cfg.Database.Host +
			":" +
			State.Cfg.Database.Port +
			"/?authSource=" +
			State.Cfg.Database.Name +
			"&connect=direct")

	return mongo.Connect(ctx, options)
}
