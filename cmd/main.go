package main

import (
	"log"

	"github.com/codedbypm/mongo-driver/mongo"
)

func main() {
	// Read secret
	mongoURI, err := mongo.GenerateURI("agora-polis")

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(mongoURI)
}
