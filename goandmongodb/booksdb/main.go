package main

import (
	"GoWebDevCourse/goandmongodb/booksdb/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Book struct {
	Isbn   string  `json:"isbn" bson:"isbn"`
	Title  string  `json:"title" bson:"title"`
	Author string  `json:"author" bson:"author"`
	Price  float64 `json:"price" bson:"price"`
}

func main() {
	AllBooks2()
	fmt.Println("-----")
	OneBook("978-1503379640")
}

func AllBooks2() {
	var books []Book
	filter := bson.D{}
	//filter := bson.E{"isbn", bson.E{"$ne", "null"}}

	cur, err := db.BooksCollection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("error executing find({}) request " + err.Error())
		return
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), &books); err != nil {
		fmt.Println("error retrieving books " + err.Error())
		//return
	}
	fmt.Println(books)
}

func AllBooks() {
	var books []Book
	filter := bson.D{}
	//filter := bson.E{"isbn", bson.E{"$ne", "null"}}

	cur, err := db.BooksCollection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("error executing find({}) request " + err.Error())
		return
	}

	var book *Book
	for {
		if !cur.Next(context.Background()) {
			break
		}
		if err := cur.Decode(&book); err != nil {
			fmt.Println("error unmarshalling to book " + err.Error())
		}
		fmt.Println("book = ")
		fmt.Println(book)
		books = append(books, *book)
	}

	fmt.Println(books)
}

func OneBook(isbn string) {
	var book *Book

	find := bson.D{bson.E{"isbn", isbn}}

	//find := bson.E{"isbn", isbn}
	res := db.BooksCollection.FindOne(context.Background(), find)

	if err := res.Decode(&book); err != nil {
		fmt.Println("Error returning book " + err.Error())
	}

	fmt.Println(book)
}
