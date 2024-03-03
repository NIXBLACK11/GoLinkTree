package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(username string, password string) (string, error){
	client, err := ConnectToMongoDB();

	if(err != nil) {
		return "", err
	}
	defer client.Disconnect(context.Background())
	
	// Get the Users collection
	usersCollection := client.Database("yourdbname").Collection("Users")

	// Find the user by username and password
	var user bson.M
	err = usersCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User does not exist
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	// User exists
	return "user exists", nil

}
