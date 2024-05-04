// package models

// import (
// 	"context"
// 	"errors"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func CheckUserExists(username string, password string) (bool, error){
// 	client, err := ConnectToMongoDB()

// 	if(err != nil) {
// 		return false, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	usersCollection := client.Database("GoLinkTree").Collection("Users")

// 	var user bson.M
// 	err = usersCollection.FindOne(context.TODO(), bson.D{
// 		{Key: "username", Value: username},
// 		{Key: "password", Value: password},
// 	}).Decode(&user)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return false, nil
// 		}
// 		return false, nil
// 	}

// 	return true, nil
// }

// func ShowUserLinks(username string) ([]map[string]string, error) {
// 	client, err := ConnectToMongoDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	usersCollection := client.Database("GoLinkTree").Collection("Users")

// 	var user UserLinks

// 	filter := bson.M{"username": username}

// 	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, errors.New("no user found with the provided username")
// 		}
// 		return nil, err
// 	}

// 	return user.Links, nil
// }

// func InsertLink(username string, link Link) (bool, error) {
// 	client, err := ConnectToMongoDB()
// 	if err != nil {
// 		return false, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	usersCollection := client.Database("GoLinkTree").Collection("Users")

// 	// Define filter to identify the document to update
// 	filter := bson.M{"username": username}

// 	// Define update operation
// 	update := bson.M{"$push": bson.M{"Links": bson.M{link.Name: link.URL}}}

// 	// Perform update operation
// 	_, err = usersCollection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func DeleteLink(username string, Link RemLink) (bool, error) {
// 	client, err := ConnectToMongoDB()
// 	if err != nil {
// 		return false, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	usersCollection := client.Database("GoLinkTree").Collection("Users")

// 	// Define filter to identify the document to update
// 	filter := bson.M{"username": username}

// 	linkName := Link.Name

// 	// Define update operation to remove the link from the array
// 	update := bson.M{"$pull": bson.M{"Links": bson.M{linkName: bson.M{"$exists": true}}}}

// 	// Perform update operation
// 	result, err := usersCollection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return false, err
// 	}

// 	// Check if any documents matched the filter criteria
// 	if result.ModifiedCount == 0 {
// 		// If no documents were modified, it means the link wasn't found
// 		return false, nil
// 	}

// 	return true, nil
// }

// func CreateUser(user User) (bool, error) {
// 	client, err := ConnectToMongoDB()
// 	if err != nil {
// 		return false, err
// 	}
// 	defer client.Disconnect(context.Background())

// 	usersCollection := client.Database("GoLinkTree").Collection("Users")

// 	_, err = usersCollection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(username string, password string) (bool, error) {
	if MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := MongoDBClient.Database("GoLinkTree").Collection("Users")

	var user bson.M
	err := usersCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, nil
	}

	return true, nil
}

func ShowUserLinks(username string) ([]map[string]string, error) {
	if MongoDBClient == nil {
		return nil, errors.New("MongoDB client is not initialized")
	}

	usersCollection := MongoDBClient.Database("GoLinkTree").Collection("Users")

	var user UserLinks

	filter := bson.M{"username": username}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no user found with the provided username")
		}
		return nil, err
	}

	return user.Links, nil
}

func InsertLink(username string, link Link) (bool, error) {
	if MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := MongoDBClient.Database("GoLinkTree").Collection("Users")

	// Define filter to identify the document to update
	filter := bson.M{"username": username}

	// Define update operation
	update := bson.M{"$push": bson.M{"Links": bson.M{link.Name: link.URL}}}

	// Perform update operation
	_, err := usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteLink(username string, Link RemLink) (bool, error) {
	if MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := MongoDBClient.Database("GoLinkTree").Collection("Users")

	// Define filter to identify the document to update
	filter := bson.M{"username": username}

	linkName := Link.Name

	// Define update operation to remove the link from the array
	update := bson.M{"$pull": bson.M{"Links": bson.M{linkName: bson.M{"$exists": true}}}}

	// Perform update operation
	result, err := usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	// Check if any documents matched the filter criteria
	if result.ModifiedCount == 0 {
		// If no documents were modified, it means the link wasn't found
		return false, nil
	}

	return true, nil
}

func CreateUser(user User) (bool, error) {
	if MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := MongoDBClient.Database("GoLinkTree").Collection("Users")

	_, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return false, err
	}

	return true, nil
}
