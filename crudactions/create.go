package crudactions

import (
	"context"
	"fmt"
	"time"

	"github.com/codedbypm/gcloud-secret-manager/access"
	"github.com/codedbypm/gcloud-secret-manager/decrypt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Create inserts a document in the given collection and database
func Create(dbName string, collectionName string, document interface{}) (interface{}, error) {

	const mongoUserSecretName = "agora-secret-mongo-user"
	const mongoPassSecretName = "agora-secret-mongo-pass"
	const keyID = "projects/agora-262523/locations/europe-west1/keyRings/agora-key-ring/cryptoKeys/agora-key/cryptoKeyVersions/latest"

	mongoUserSecret, secretError := access.GetSecret("agora-262523", mongoUserSecretName)
	if secretError != nil {
		return nil, fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoUserSecretName, secretError)
	}

	mongoPassSecret, secretError := access.GetSecret("agora-262523", mongoPassSecretName)
	if secretError != nil {
		return nil, fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoPassSecretName, secretError)
	}

	mongoUser, decryptError := decrypt.DecryptSymmetric(keyID, mongoUserSecret.Payload.Data)
	if decryptError != nil {
		return nil, fmt.Errorf("Error: could not decrypt secret %s (%s)", secretError)
	}

	mongoPass, decryptError := decrypt.DecryptSymmetric(keyID, mongoPassSecret.Payload.Data)

	var mongoURI = fmt.Sprint("mongodb+srv://%s:%s@agorapolis-001-ymzlz.gcp.mongodb.net", mongoUser, mongoPass)

	// Create Mongo connection
	mongoContext, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	mongoClient, mongoError := mongo.Connect(mongoContext, options.Client().ApplyURI(mongoURI))
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
