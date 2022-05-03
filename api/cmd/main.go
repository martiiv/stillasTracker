package main

import (
	"log"
	"os"
	"stillasTracker/api/database"
	"stillasTracker/api/endpoints"
)

/**
Class main
Will run the api
Version 0.1
Last update 08.03.2022 Martin Iversen
*/

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Printf("Starting API")

	database.DatabaseConnection()
	log.Printf("Started database")

	endpoints.Handle()
}
