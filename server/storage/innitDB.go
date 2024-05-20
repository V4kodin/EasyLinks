package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Coll struct {
	c *mongo.Collection
}

func InitDB() (*Coll, error) {

	bdHost := os.Getenv("MONGO_ADDRES")
	bdPort := os.Getenv("MONGO_PORT")
	bdUser := os.Getenv("MONGO_USERNAME")
	bdPass := os.Getenv("MONGO_PASSWORD")
	conn := "mongodb://" + bdUser + ":" + bdPass + "@" + bdHost + ":" + bdPort
	clientOptions := options.Client().ApplyURI(conn)
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
