package storage

import (
	"context"
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
	_, err := coll.c.InsertOne(context.TODO(), shortURL)
	if err == nil {
		return ("Inserted document with _id:" + shortURL.ID), 0
	} else if mongo.IsDuplicateKeyError(err) {
		return "Alread in db ", 6
	} else if mongo.IsTimeout(err) {
		return "Timout errr ", 14
	} else {
		log.Fatal(err)
	}
	return "", 2
}

func (coll *Coll) FindOne(id string) *ShortURL {
	var shortURL ShortURL
	err := coll.c.FindOne(context.TODO(), id).Decode(&shortURL)
	if err != nil {
		log.Fatal(err)
	}
	return &shortURL

}
