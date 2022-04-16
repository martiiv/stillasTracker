package main

import (
	"fmt"
	"stillasTracker/api/database"
	"stillasTracker/api/endpoints"
	"stillasTracker/api/mqtt"
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
	mqtt.InitializeMQTTClient()

	endpoints.Handle()
}
