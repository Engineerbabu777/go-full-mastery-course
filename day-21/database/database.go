package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "string"

const colName = "todo-table"
const collectionName = "todo"

// COLLECTION REFERENCE!
var collection *mongo.Collection

func DatabaseConnect() {

	// CLIENT!
	mongoClient := options.Client().ApplyURI(connectionString)

	// CONNECT TO THE MONGODB DATABASE!
	client, err := mongo.Connect(context.TODO(), mongoClient)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("Connect to DATABASE!")
}
