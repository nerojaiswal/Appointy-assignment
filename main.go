package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID string `json:"userId" bson:"userId" binding: "reqired"`
	Name string   `json:"name," bson:"name"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	Email string   `json:"email" bson:"email"`
	TimeStamp string `json:"timeStamp" bson: "timeStamp"`
	DateOfBirth string `json:"dob,omitempty" bson: "dob,omitempty"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Users
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}


func main() {
		//Init Router
	r := mux.NewRouter()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx,clientOptions)
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", nil))
}
