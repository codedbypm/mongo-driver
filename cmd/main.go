package main

import (
	"log"
	"time"

	"github.com/codedbypm/mongo-driver/mongo"
)

func main() {

	createURI()
	createAuthChallenge()
	readAuthChallenges()

}

func createURI() {
	// create URI
	mongoURI, err := mongo.GenerateURI("agora-polis")
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

	// Decode the JSON
	challenge := struct {
		Id           string
		Content      string
		CreationDate time.Time
	}{}

	challenges, err := mongo.Read("auth", "challenges")
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(challenges)
}
