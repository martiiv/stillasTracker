package main

import (
	"log"
	"stillasTracker/api/database"
	"stillasTracker/api/endpoints"
)

/**
Class main
Will run the api
Version 0.1
Last update 08.03.2022 Martin Iversen
*/
var (
	InfoLogger     *log.Logger
	DatabaseLogger *log.Logger
)

func main() {
	endpoints.InitLog()
	InfoLogger.Printf("Starting API")

	database.DatabaseConnection()
	DatabaseLogger.Printf("Started database")

	endpoints.Handle()
}
