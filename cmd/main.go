package main

import (
	"log"
	"time"

	"github.com/codedbypm/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	createURI()
	createAuthChallenge()
	readAuthChallenges()

}

func createURI() {
	// create URI
	mongoURI, err := mongo.GenerateURI()
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(mongoURI)
}

func createAuthChallenge() {

	// Decode the JSON
	challenge := struct {
		Id           string
		Content      string
		CreationDate time.Time
	}{}

	challenges, err := mongo.Create("auth", "challenges", challenge)
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(challenges)
}

func readAuthChallenges() {

	filter := bson.D{{"content", "paolo"}}

	challenges, err := mongo.ReadOne("auth", "challenges", filter)
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(challenges)
}
