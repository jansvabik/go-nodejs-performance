package data

import (
	"context"
	"time"

	"github.com/jansvabik/go-nodejs-performance/go/random"

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
	// create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try to find the data
	URLs := []URL{}
	cursor, err := collection().Find(ctx, bson.M{}, options.Find().SetProjection(bson.M{
		"_id":     0,
		"url":     1,
		"target":  1,
		"used":    1,
		"lastUse": 1,
	}))
	defer cursor.Close(ctx)
	if err != nil {
		return URLs, err
	}

	// create an array of URLs
	for cursor.Next(ctx) {
		var url URL
		cursor.Decode(&url)
		URLs = append(URLs, url)
	}

	return URLs, err
}

// GetByURL returns data about one specific URL
func GetByURL(URLID string) (*URL, error) {
	// create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try to get data from database
	var result URL
	err := collection().FindOne(ctx, bson.M{
		"url": URLID,
	}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create creates new document of URL
func Create(target string) (*URL, error) {
	t := time.Now()
	doc := URL{
		URL:      random.String(6),
		Target:   target,
		Used:     0,
		LastUse:  nil,
		Created:  &t,
		Modified: &t,
		Password: random.String(32),
	}

	// create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try to store the document in db
	_, err := collection().InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

// Update updates existing document
func Update() {

}

// Delete deletes specified document permanently
func Delete(urlID string) error {
	// create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// try to remove the document from database
	_, err := collection().DeleteOne(ctx, bson.M{"url": urlID})
	if err != nil {
		return err
	}

	return nil
}

// collection returns the set connection
func collection() *mongo.Collection {
	return app.State.MongoClient.Database("gonodejsperf").Collection("url")
}
