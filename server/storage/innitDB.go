package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Coll struct {
	c *mongo.Collection
}

func InitDB() (*Coll, error) {
	port := "8081"
	clientOptions := options.Client().ApplyURI("mongodb://localhost:" + port)

	connection, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	coll := &Coll{
		c: connection.Database("core").Collection("shortUrls"),
	}

	log.Println("Connected to MongoDB!")

	return coll, err
}
