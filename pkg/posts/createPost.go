package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newPost struct {
	DateAdded string
	PostDate  string
	PostTime  string
	Image     string
	Caption   string
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("holla--")
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t newPost
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

		collection := client.Database("testing").Collection("posts")
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		res, _ := collection.InsertOne(ctx, t)
		id := res.InsertedID

		fmt.Fprint(w, id)
		log.Println(t.Caption)
	}
}
