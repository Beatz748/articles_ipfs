package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getArticleFromMongoDB(id string, coll *mongo.Collection, article *Article) error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find the article in MongoDB
	if err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(article); err != nil {
		return fmt.Errorf("Article not found in MongoDB: %v", err)
	}

	return nil
}

func addArticleToMongoDB(coll *mongo.Collection, article *Article) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert into MongoDB
	_, err := coll.InsertOne(ctx, article)
	if err != nil {
		return fmt.Errorf("Failed to insert article into MongoDB: %v", err)
	}

	return nil
}
