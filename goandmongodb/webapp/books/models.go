package books

import (
	"GoWebDevCourse/goandmongodb/webapp/config"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

type Book struct {
	Isbn   string  `json:"isbn" bson:"isbn"`
	Title  string  `json:"title" bson:"title"`
	Author string  `json:"author" bson:"author"`
	Price  float64 `json:"price" bson:"price"`
}

func AllBooks() ([]Book, error) {
	var books []Book
	filter := bson.D{}
	cur, err := config.BooksCollection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("error executing find({}) request " + err.Error())
		return nil, err
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), &books); err != nil {
		fmt.Println("error retrieving books " + err.Error())
		return nil, err
	}
	fmt.Println("Books: ")
	fmt.Println(books)
	return books, nil
}

func OneBook(r *http.Request) (Book, error) {
	var book Book
	isbn := r.FormValue("isbn")

	if isbn == "" {
		return book, errors.New("400. Isbn is not provided.")
	}

	find := bson.D{bson.E{"isbn", isbn}}
	res := config.BooksCollection.FindOne(context.Background(), find)

	if err := res.Decode(&book); err != nil {
		return book, errors.New("500. Error returning book")
	}

	return book, nil
}

func PutBook(r *http.Request) (Book, error) {
	book := Book{}

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		return book, errors.New("400. Form is invalid")
	}

	// convert form values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return book, errors.New("400. Price value is invalid")
	}
	book.Price = float64(f64)

	// update

	put := bson.M{"isbn": book.Isbn, "title": book.Title, "author": book.Author, "price": book.Price}

	_, err = config.BooksCollection.InsertOne(context.Background(), put)

	if err != nil {
		return book, errors.New("500. Error inserting the doc")
	}

	return book, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	book := Book{}

	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// validate form values
	if book.Isbn == "" || book.Title == "" || book.Author == "" || price == "" {
		return book, errors.New("400. Form is invalid")
	}

	// convert form values
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return book, errors.New("400. Price value is invalid")
	}
	book.Price = float64(f64)

	fmt.Println("[UPDATE] received isbn is " + book.Isbn)

	filter := bson.D{bson.E{"isbn", book.Isbn}}

	update := bson.D{{"$set",
		bson.D{
			bson.E{"isbn", book.Isbn},
			bson.E{"title", book.Title},
			bson.E{"author", book.Author},
			bson.E{"price", book.Price},
		}},
	}

	_, err = config.BooksCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return book, errors.New("500. Error updating the doc")
	}

	return book, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")

	if isbn == "" {
		return errors.New("400. isbn is invalid")
	}

	filter := bson.D{bson.E{"isbn", isbn}}
	_, err := config.BooksCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.New("500. Error deleting doc")
	}
	return nil
}
