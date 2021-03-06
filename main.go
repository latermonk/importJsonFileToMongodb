/*
https://kb.objectrocket.com/mongo-db/how-to-insert-mongodb-documents-from-json-using-the-golang-driver-457
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"reflect"
	"time"
)

type MongoFields struct {
	//Key string `json:"key,omitempty"`
	ID        int    `json:"_id"`
	FieldStr  string `json:"Field Str"`
	FieldInt  int    `json:"Field Int"`
	FieldBool bool   `json:"Field Bool"`
}

func main() {

	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://myUserAdmin:abc123@10.252.49.56:27017")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// Access a MongoDB collection through a database
	col := client.Database("JSON_docs").Collection("JSON Collection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "n")

	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("./docs.json")
	if err != nil {
		// Print any IO errors with the .json file
		fmt.Println("ioutil.ReadFile ERROR:", err)
	} else {
		// Print the values of the JSON docs, and insert them if no error
		fmt.Println("ioutil.ReadFile byteValues TYPE:", reflect.TypeOf(byteValues))
		fmt.Println("byteValues:", byteValues, "n")
		fmt.Println("byteValues:", string(byteValues))

		// Declare an empty slice for the MongoFields docs
		var docs []MongoFields

		// Unmarshal the encoded JSON byte string into the slice
		err = json.Unmarshal(byteValues, &docs)

		// Print MongoDB docs object type
		fmt.Println("nMongoFields Docs:", reflect.TypeOf(docs))

		// Iterate the slice of MongoDB struct docs
		for i := range docs {

			// Put the document element in a new variable
			doc := docs[i]
			fmt.Println("ndoc _id:", doc.ID)
			fmt.Println("doc Field Str:", doc.ID)

			// Call the InsertOne() method and pass the context and doc objects
			result, insertErr := col.InsertOne(ctx, doc)

			// Check for any insertion errors
			if insertErr != nil {
				fmt.Println("InsertOne ERROR:", insertErr)
			} else {
				fmt.Println("InsertOne() API result:", result)
			}
		}
	}
}
