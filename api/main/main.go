package main

import (
	"fmt"
	"stillasTracker/api/Database"
	"stillasTracker/api/endpoints"
)

/**
Class main
Will run the api
Version 0.1
Last update 08.03.2022 Martin Iversen
*/
func main() {
	Database.DatabaseConnection()
	fmt.Println("Starting API")
	//startTime = time.Now()
	fmt.Println("initialized handler")
	endpoints.Handle()

}
