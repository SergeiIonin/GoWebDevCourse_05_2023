package main

import (
	"GoWebDevCourse/goandmongodb/02_app/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	client := getSession()

	/*collection := client.Database("go_webdev").Collection("users")

	res, err := collection.InsertOne(context.Background(), bson.M{"Name": "James Bond", "Gender": "male", "Age": 32, "Id": "777"})

	id := res.InsertedID

	if err != nil {
		fmt.Println("Error saving user")
		return
	}

	fmt.Println("inserted id = ", id)*/

	var u *models.User

	filter := bson.D{bson.E{Key: "Id", Value: "777"}}
	//filter := bson.E{Key: "Id", Value: fmt.Sprintf(`"%s"`, "777")}
	//filter := bson.D{}

	cur, err := client.Database("go_webdev").Collection("users").Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	/*if err = cur.All(context.Background(), &u); err != nil { // NB!
		fmt.Println("error during decoding: ", err.Error())
	}*/

	for cur.Next(context.Background()) {
		err = cur.Decode(&u)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("user is ", u)

}

func getSession() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic("can't setup db connection " + err.Error())
	}

	return client
}
