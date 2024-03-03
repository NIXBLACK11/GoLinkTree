package models

import (
	"GoLinkTree/custom"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		err := custom.MyError("DATABASE URL not found in env")
		log.Fatal("Error getting DATABASE_URL:")
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(DATABASE_URL)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
