package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const baseURL = "/stillastracking/v1/api"

/**
Class API handler
Will initialize all the endpoints and set ports
Version 0.1
Last edit 08.03.2022 by Martin Iversen
*/
func Handle() {
	fmt.Println("Listening on port" + getPort())

	//Scaffolding endpoints
	http.HandleFunc(baseURL+"/unit", scaffoldingRequest) //country endpoint
	//Project endpoint
	http.HandleFunc(baseURL+"/project", projectRequest) //DELETE, POST, GET
	http.HandleFunc(baseURL+"/storage", storageRequest)
	//Profile endpoint
	http.HandleFunc(baseURL+"/user/", profileRequest)
	log.Println(http.ListenAndServe(getPort(), nil))

}

/*
Function used for setting the port for the application
We use localhost 8080 for testing
Takes no parameters
Returns the port the software is listening on
*/
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
