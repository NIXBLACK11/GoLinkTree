package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(username string, password string) (bool, error){
	client, err := ConnectToMongoDB();

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
