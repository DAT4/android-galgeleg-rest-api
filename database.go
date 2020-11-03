package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type HighScore struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Word   Word               `json:"word"`
	Player string             `json:"player"`
	Time   int                `json:"time"`
	Hints  int                `json:"hints"`
	Wrongs int                `json:"wrongs"`
}

type Category struct {
	Title string `json:"title"`
	Words []Word `json:"words"`
}

type Word struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Word        string             `json:"word"`
	Difficulty  int                `json:"difficulty"`
	Description string             `json:"description"`
	Hint1       string             `json:"hint1"`
	Hint2       string             `json:"hint2"`
	Hint3       string             `json:"hint3"`
	Category    string             `json:"category"`
}

func getHighScores() (scores []HighScore) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("hangmanHighScore").Find(ctx, bson.M{})
	var highScore HighScore
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&highScore)
		scores = append(scores, highScore)
	}
	defer client.Disconnect(ctx)
	return scores
}

func addHighScore(highScore HighScore) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_, err = client.Database("dtu").Collection("hangmanHighScore").InsertOne(ctx, highScore)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteHighScore(score HighScore) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": score.ID}
	_, err = client.Database("dtu").Collection("hangmanHighScore").DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
}

func getWords() (words []Word) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("hangman").Find(ctx, bson.M{})
	var word Word
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&word)
		words = append(words, word)
	}
	defer client.Disconnect(ctx)
	return words
}

func getCategories() (categories []Category) {
	var words []Word
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("hangman").Find(ctx, bson.M{})
	var word Word
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&word)
		words = append(words, word)
	}
	defer client.Disconnect(ctx)
	for _, word := range words {
		is := false
		for i, category := range categories {
			if word.Category == category.Title {
				is = true
				categories[i].Words = append(categories[i].Words, word)
			}
		}
		if !is {
			categories = append(categories, Category{
				Title: word.Category,
				Words: []Word{word},
			})
		}
	}
	return categories
}

func createWord(word Word) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_, err = client.Database("dtu").Collection("hangman").InsertOne(ctx, word)
	if err != nil {
		fmt.Println(err)
	}
}

func updateWord(word Word) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": word.ID}
	update := bson.M{"$set": word}
	_, err = client.Database("dtu").Collection("hangman").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
}
func deleteWord(word Word) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": word.ID}
	_, err = client.Database("dtu").Collection("hangman").DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
}
