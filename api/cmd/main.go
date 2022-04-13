package main

import (
	"fmt"
	"stillasTracker/api/apiTools"
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
	fmt.Println("Starting API")

	database.DatabaseConnection()
	fmt.Println("Started database")

	fmt.Println("Initializing MQTT Server")
	apiTools.InitializeMQTTClient()

	endpoints.Handle()
}
