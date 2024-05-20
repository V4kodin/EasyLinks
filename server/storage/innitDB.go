package storage

import (
	"EasyLinks/server"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Coll struct {
	c *mongo.Collection
}

func InitDB() (*Coll, error) {

	clientOptions := options.Client().ApplyURI(server.BdAddress)

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
