package storage

import (
	"context"
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

func (coll *Coll) InsertOne(shortURL *ShortURL) {
	result, err := coll.c.InsertOne(context.TODO(), shortURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted document with _id:", result.InsertedID)

}
