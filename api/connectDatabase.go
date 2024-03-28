package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Gubaz311/myOwnUnila_server/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Connect() {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	namadatabase := os.Getenv("DB_NAME")
	fmt.Println("We are getting the env values with", namadatabase)

	server.Start(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Info(":8080")
}
