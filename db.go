package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/game-req")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}
	fmt.Println("Connected to MongoDB!")
}

func SaveGameRequirements(game GameRequirements) error {
	collection := client.Database("gameDB").Collection("requirements")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"FinalGameName":        game.FinalGameName,
		"MinCPU":               game.MinCPU,
		"MinRAM":               game.MinRAM,
		"MinVideoCard":         game.MinVideoCard,
		"MinDedicatedVideoRAM": game.MinDedicatedVideoRAM,
		"MinDiskSpace":         game.MinDiskSpace,
		"MinOS":                game.MinOS,
	})
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	fmt.Println("Game requirements saved successfully!")
	return nil
}
