package controllers

import (
	"GoWebDevCourse/goandmongodb/02_app/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// added session to our userController
type UserController struct {
	client *mongo.Client
}

// added session to our userController
func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

// return 404 if user in nf
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var u *models.User

	/*filter := bson.D{bson.E{Key: "id", Value: fmt.Sprintf(`"%s"`, p.ByName("id"))}}

	res := uc.client.Database("go_webdev").Collection("users").FindOne(context.Background(), filter)

	err := res.Decode(u)

	if err != nil {
		fmt.Println(err.Error())
	}*/

	filter := bson.D{bson.E{Key: "Id", Value: p.ByName("id")}}

	cur, err := uc.client.Database("go_webdev").Collection("users").Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	for cur.Next(context.Background()) {
		err = cur.Decode(&u)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	uj, err := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error decoding user", http.StatusInternalServerError)
	}

	doc := bson.D{bson.E{Key: "Name", Value: u.Name}, bson.E{Key: "Gender", Value: u.Gender},
		bson.E{Key: "Age", Value: u.Age}, bson.E{Key: "Id", Value: u.Id}}

	_, err = uc.client.Database("go_webdev").Collection("users").InsertOne(context.Background(), doc)

	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
	}

	json.NewDecoder(r.Body).Decode(&u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	collection := uc.client.Database("go_webdev").Collection("users")
	filter := bson.D{bson.E{Key: "Id", Value: id}}

	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if !cur.Next(context.Background()) {
		http.Error(w, fmt.Sprintf("User %s not found", id), http.StatusNotFound)
		return
	}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s deleted successfully", id)
}
