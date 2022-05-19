package main

import (
	"fmt"
	"stillasTracker/api/database"
	"stillasTracker/api/endpoints"
)

/**
Class main
Will run the api
Version 1.0
Last update 08.03.2022 Martin Iversen
*/

func main() {
	fmt.Printf("Starting API")

	database.DatabaseConnection()
	fmt.Printf("Started database")

	endpoints.Handle()
}
