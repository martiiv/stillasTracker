package endpoints

import (
	"fmt"
	"os"
)

const baseURL = "/stillastracking/v1/api"

/**
Class API handler
Will initialize all the endpoints and set ports
Version 0.1
Last edit 08.03.2022 by Martin Iversen
*/
func handle() {
	fmt.Println("Listening on port" + getPort())
	//TODO Legg inn alle endpoints her se POSTMAN for dokumentasjon

	//Scaffolding endpoints
	//http.HandleFunc(baseURL + "/unit") //Scaffolding unit: GET,Post and DELETE
	//r.HandleFunc(baseURL+"/unit/{id}/history", ) //country endpoint
	//r.HandleFunc(baseURL+"/unit/{id}", ) //country endpoint
	//Project endpoint
	//r.HandleFunc(baseURL+"project/{id}&scaffolding=true", )
	//r.HandleFunc(baseURL+"project", ) //DELETE, POST, GET
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
