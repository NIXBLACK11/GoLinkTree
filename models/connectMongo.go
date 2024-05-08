package models

import (
    "GoLinkTree/custom"
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    // "github.com/joho/godotenv"
)

var MongoDBClient *mongo.Client

func InitMongoDB() error {
    // err := godotenv.Load()
    // if err != nil {
    //     return err
    // }

    DATABASE_URL := os.Getenv("DATABASE_URL")
    
    if DATABASE_URL == "" {
        err := custom.MyError("DATABASE URL not found in env")
        return err
    }

    clientOptions := options.Client().ApplyURI(DATABASE_URL)

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }
    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    MongoDBClient = client

    fmt.Println("Connected to MongoDB!")
    return nil
}
