package api

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Connect() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
