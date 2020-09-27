package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type category struct {
	Title string `json:"title"`
	Words []word `json:"words"`
}

type word struct {
	Word        string   `json:"word"`
	Difficulty  int      `json:"difficulty"`
	Description string   `json:"description"`
	Leads       []string `json:"leads"`
}

func getCategories() (categories []category) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("hangman").Find(ctx, bson.M{})
	var category category
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&category)
		categories = append(categories, category)
	}
	defer client.Disconnect(ctx)
	return categories
}
