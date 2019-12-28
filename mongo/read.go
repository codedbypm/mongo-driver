package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Read fetch a document from the given collection and database
func ReadOne(dbName string, collectionName string, filter interface{}) (interface{}, error) {

	// Create Mongo connection
	mongoContext, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	mongoURI, mongoURIError := GenerateURI()
	if mongoURIError != nil {
		return nil, fmt.Errorf("Error: could not generate Mongo URI (%s)", mongoURIError)
	}

	mongoClient, mongoError := mongo.Connect(mongoContext, options.Client().ApplyURI(mongoURI))
	mongoError = mongoClient.Ping(mongoContext, readpref.Primary())
	if mongoError != nil {
		return nil, fmt.Errorf("Error: could not connect to Mongo (%s)", mongoError)
	}

	// Get collection
	collection := mongoClient.Database(dbName).Collection(collectionName)

	// Insert data
	findOneContext := context.Background()
	findOneContext, insertCancel := context.WithTimeout(findOneContext, 5*time.Second)
	defer insertCancel()

	var document interface{}
	resultError := collection.FindOne(findOneContext, filter).Decode(&document)
	if resultError != nil {
		return nil, fmt.Errorf("Error: could not read document %s, (%s)", document, resultError)
	}
	return document, nil
}
