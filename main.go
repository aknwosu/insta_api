package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aknwosu/insta_api/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// mongo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("testing").Collection("numbers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Fprint(w, id)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/createPost", posts.CreatePost)
	http.ListenAndServe(":8080", nil)
}
