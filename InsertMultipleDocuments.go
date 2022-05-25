package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {

	//client
	clientOptions := options.Client().ApplyURI("mongodb://myUserAdmin:abc123@10.252.49.56:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}

	coll := client.Database("myDB").Collection("favorite_books")
	docs := []interface{}{
		bson.D{{"title", "My Brilliant Friend"}, {"author", "Elena Ferrante"}, {"year_published", 2012}},
		bson.D{{"title", "Lucy"}, {"author", "Jamaica Kincaid"}, {"year_published", 2002}},
		bson.D{{"title", "Cat's Cradle"}, {"author", "Kurt Vonnegut Jr."}, {"year_published", 1998}},
	}

	result, err := coll.InsertMany(context.TODO(), docs)
	list_ids := result.InsertedIDs
	fmt.Printf("Documents inserted: %v\n", len(list_ids))

	for _, id := range list_ids {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}
