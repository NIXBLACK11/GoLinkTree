package models

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(username string, password string) (bool, error){
	client, err := ConnectToMongoDB()

	if(err != nil) {
		return false, err
	}
	defer client.Disconnect(context.Background())
	
	usersCollection := client.Database("GoLinkTree").Collection("Users")

	var user bson.M
	err = usersCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, fmt.Errorf("user not found")
		}
		return false, err
	}

	return true, nil
}

func ShowUserLinks(username string) ([]map[string]string, error) {
	client, err := ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	usersCollection := client.Database("GoLinkTree").Collection("Users")

	var user UserLinks

	filter := bson.M{"username": username}

	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no user found with the provided username")
		}
		return nil, err
	}

	return user.Links, nil
}

func InsertLink(username string, link Link) (bool, error) {
	client, err := ConnectToMongoDB()
	if err != nil {
		return false, err
	}
	defer client.Disconnect(context.Background())

	usersCollection := client.Database("GoLinkTree").Collection("Users")

	// Define filter to identify the document to update
	filter := bson.M{"username": username}

	// Define update operation
	update := bson.M{"$push": bson.M{"Links": bson.M{link.Name: link.URL}}}

	// Perform update operation
	_, err = usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}