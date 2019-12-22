package main

import (
	"fmt"

	"source.cloud.google.com/agora-262523/agora-git/functions/getChallenge"
	mongoDriver "source.cloud.google.com/agora-262523/mongo-driver"
)

func main() {
	var challenge getChallenge.Challenge

	const db = "auth"
	const collection = "challenges"
	secret, err := mongoDriver.Create(db, challenges, challenges)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secret)
}
