package data

import (
	"context"
	"time"

	"github.com/jansvabik/go-nodejs-performance/go/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// URL is the model of URL in database
type URL struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	URL      string              `json:"url,omitempty" bson:"url"`
	Target   string              `json:"target,omitempty" bson:"target"`
	Used     int32               `json:"used" bson:"used"`
	LastUse  *time.Time          `json:"lastUse" bson:"lastUse,omitempty"`
	Created  *time.Time          `json:"created,omitempty" bson:"created"`
	Modified *time.Time          `json:"modified,omitempty" bson:"modified,omitempty"`
	Password string              `json:"password,omitempty" bson:"password"`
}

// GetList gets all URLs from database and returns it
func GetList() ([]URL, error) {
	URLs := []URL{}

	// create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try to find the data
	p := bson.M{
		"_id":     0,
		"url":     1,
		"target":  1,
		"used":    1,
		"lastUse": 1,
	}
	cursor, err := collection().Find(ctx, bson.M{}, options.Find().SetProjection(p))
	defer cursor.Close(ctx)
	if err != nil {
		return URLs, err
	}

	// create and array of URLs
	for cursor.Next(ctx) {
		var url URL
		cursor.Decode(&url)
		URLs = append(URLs, url)
	}

	return URLs, err
}

// collection returns the set connection
func collection() *mongo.Collection {
	return app.State.MongoClient.Database("gonodejsperf").Collection("url")
}
