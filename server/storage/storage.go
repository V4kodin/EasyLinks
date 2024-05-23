package storage

import (
	"EasyLinks/server/pkg/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type URLWorker interface {
	AddValue(key, value, userID string) error
	Ping() (bool, error)
}

func (coll *Coll) InsertOne(ShortURL *models.ShortURL) *models.ShortURL {
	result, err := coll.c.InsertOne(context.TODO(), ShortURL)
	ShortURL.ID = result.InsertedID.(string)
	if err == nil {
		ShortURL.Error = 0
		return ShortURL
	} else if mongo.IsDuplicateKeyError(err) {
		ShortURL.Error = 6
		return ShortURL
	} else if mongo.IsTimeout(err) {
		ShortURL.Error = 14
		return ShortURL
	} else {
		ShortURL.Error = 2
		log.Fatal(err)
		return ShortURL
	}
}

func (coll *Coll) FindOne(ShortURL *models.ShortURL) *models.ShortURL {

	filter := bson.D{{"_id", ShortURL.ID}}
	err := coll.c.FindOne(context.TODO(), filter).Decode(&ShortURL)

	switch {
	case err == nil:
		ShortURL.Error = 0
		return ShortURL
	case errors.Is(err, mongo.ErrNilValue):
		ShortURL.Error = 5
		return ShortURL
	case errors.Is(err, mongo.MarshalError{Err: err}):
		ShortURL.Error = 3
		return ShortURL
	case errors.Is(err, mongo.ErrNoDocuments):
		ShortURL.Error = 5
		return ShortURL
	default:
		ShortURL.Error = 2
		log.Fatal(err)
		return ShortURL
	}
}
