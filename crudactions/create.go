package crudactions

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"source.cloud.google.com/agora-262523/gcloud-secret-manager/secretaccess"
)

// Create inserts a document in the given collection and database
func Create(dbName string, collectionName string, document interface{}) (interface{}, error) {

	mongoUser, err := secretaccess.GetSecret()

	// Create Mongo connection
	mongoContext, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	mongoClient, mongoError := mongo.Connect(mongoContext, options.Client().ApplyURI("mongodb://localhost:27017"))
	mongoError = mongoClient.Ping(mongoContext, readpref.Primary())

	if mongoError != nil {
		return nil, fmt.Errorf("Error: could not connect to Mongo (%s)", mongoError)
	}

	// Get collection
	collection := mongoClient.Database(dbName).Collection(collectionName)

	// Insert data
	insertContext := context.Background()
	insertContext, insertCancel := context.WithTimeout(insertContext, 5*time.Second)
	defer insertCancel()

	insertedDocument, insertError := collection.InsertOne(insertContext, document)

	if insertError != nil {
		return nil, fmt.Errorf("Error: could not insert document %s, (%s)", document, mongoError)
	}

	return insertedDocument, nil
}
