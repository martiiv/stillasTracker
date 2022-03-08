package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
)

const baseURL = "/myLift/stillasTracker/v1/"

/**
Class API handler
Will initialize all the endpoints and set ports
Version 0.1
Last edit 08.03.2022 by Martin Iversen
*/
func handle() {
	fmt.Println("Listening on port" + getPort())
	r := mux.NewRouter()
	r.HandleFunc(baseURL+"unit", genericResponse) //country endpoint

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
