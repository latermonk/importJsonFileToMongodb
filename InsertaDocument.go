/*
https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/insert/#std-label-golang-insert-guide


*/

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

	//doc
	coll := client.Database("myDB").Collection("favorite_books")
	doc := bson.D{{"title", "Invisible Cities"}, {"author", "Italo Calvino"}, {"year_published", 1974}}

	//result
	result, err := coll.InsertOne(context.TODO(), doc)

	//print
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}
