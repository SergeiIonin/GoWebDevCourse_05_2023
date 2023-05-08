package main

import (
	"GoWebDevCourse/goandmongodb/booksdb/db"
	"GoWebDevCourse/goandmongodb/webapp/config"
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
	AllBooks()
	fmt.Println("-----")
	OneBook("978-1503379640")
	fmt.Println("-----")
	bookWithOldPrice := Book{
		Isbn:   "121343",
		Title:  "Doctor Zhivago",
		Author: "B.Pasternak",
		Price:  122.38,
	}
	UpdateBook(bookWithOldPrice, 110.11)
	fmt.Println("-----")
	OneBook("121343")
	fmt.Println("-----")
	DeleteBook("978-1505255607")
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
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), &books); err != nil {
		fmt.Println("error retrieving books " + err.Error())
		//return
	}
	fmt.Println(books)
}

/*
	func AllBooks() {
		var books []Book
		filter := bson.D{}

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
*/
func OneBook(isbn string) {
	var book *Book

	find := bson.D{bson.E{"isbn", isbn}}
	//find := bson.E{"isbn", isbn}

	//find := bson.E{"isbn", isbn}
	res := db.BooksCollection.FindOne(context.Background(), find)

	if err := res.Decode(&book); err != nil {
		fmt.Println("Error returning book " + err.Error())
	}

	fmt.Println(book)
}

func UpdateBook(bookIn Book, priceIn float64) {
	book := Book{}

	book.Isbn = bookIn.Isbn
	book.Title = bookIn.Title
	book.Author = bookIn.Author
	book.Price = priceIn

	fmt.Println("[UPDATE] received isbn is " + book.Isbn)

	// update
	filter := bson.D{bson.E{"isbn", book.Isbn}}

	update := bson.D{{"$set",
		bson.D{
			bson.E{"isbn", book.Isbn},
			bson.E{"title", book.Title},
			bson.E{"author", book.Author},
			bson.E{"price", priceIn},
		}},
	}

	_, err := config.BooksCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		fmt.Println("Error updating the doc " + err.Error())
		return
	}

	fmt.Println(book)
}

func InsertBook(bk Book) {
	doc := bson.D{
		bson.E{"isbn", bk.Isbn},
		bson.E{"title", bk.Title},
		bson.E{"author", bk.Author},
		bson.E{"price", bk.Price},
	}

	if _, err := config.BooksCollection.InsertOne(context.Background(), doc); err != nil {
		fmt.Println("Error inserting doc into mongo")
	}
}

func DeleteBook(isbn string) {
	doc := bson.D{bson.E{"isbn", isbn}}
	if _, err := config.BooksCollection.DeleteOne(context.Background(), doc); err != nil {
		fmt.Println("Error deleting doc from mongo")
	}
}
