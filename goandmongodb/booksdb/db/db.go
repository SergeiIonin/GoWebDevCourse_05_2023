package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Db *mongo.Database
var BooksCollection *mongo.Collection

func init() {
	client := getMongoClient()

	err := client.Ping(context.Background(), nil)

	if err != nil {
		panic("can't get response from db " + err.Error())
	}

	Db = client.Database("bookstore")
	BooksCollection = Db.Collection("bookstore")
}

func getMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/bookstore"))

	if err != nil {
		panic("can't setup db connection " + err.Error())
	}

	return client
}
