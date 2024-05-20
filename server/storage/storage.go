package storage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type ShortURL struct {
	ID       string     `bson:"_id"`
	URL      string     `bson:"url"`
	ExpireAt *time.Time `bson:"expireAt,omitempty"`
}

type URLWorker interface {
	AddValue(key, value, userID string) error
	Ping() (bool, error)
}

func (coll *Coll) InsertOne(shortURL *ShortURL) (string, int) {
	result, err := coll.c.InsertOne(context.TODO(), shortURL)

	if err == nil {
		return result.InsertedID.(string), 0
	} else if mongo.IsDuplicateKeyError(err) {
		return "", 6
	} else if mongo.IsTimeout(err) {
		return "", 14
	} else {
		log.Fatal(err)
	}
	return "", 2
}

func (coll *Coll) FindOne(id string) (*ShortURL, int) {
	var shortURL ShortURL
	filter := bson.D{{"_id", id}}
	err := coll.c.FindOne(context.TODO(), filter).Decode(&shortURL)

	switch {
	case err == nil:
		return &shortURL, 0
	case errors.Is(err, mongo.ErrNilValue):
		return nil, 5
	case errors.Is(err, mongo.MarshalError{Err: err}):
		return nil, 3
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, 5
	default:
		log.Fatal(err)
		return nil, 2
	}
	return &shortURL, 2

}
