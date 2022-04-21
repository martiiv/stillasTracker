package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

/**
Class APIHandler.go
Class forwards requests to the appropriate endpoint and assigns the port of the program
Last modified by martiiv@stud.ntnu.no
Date: 06.04.2022
Version 0.8
*/

const baseURL = "/stillastracking/v1/api"

//Handle Function starts when launching program, function forwards the request to the appropriate endpoint
func Handle() {
	router := mux.NewRouter()

	//router.HandleFunc(baseURL+"/unit", ScaffoldingRequest) //DELETE, POST, GET

	//Scaffolding endpoint
	router.Path(baseURL+"/unit").HandlerFunc(ScaffoldingRequest).Queries("type", "{type}").Queries("id", "{id}") //GET POST PUT DELETE
	router.Path(baseURL+"/unit").HandlerFunc(ScaffoldingRequest).Queries("type", "{type}")                       //GET POST PUT DELETE
	router.Path(baseURL + "/unit").HandlerFunc(ScaffoldingRequest)                                               //GET POST PUT DELETE

	//Project endpoint
	router.HandleFunc(baseURL+"/project", ProjectRequest).Queries("id", "{id}").Queries("scaffolding", "{scaffolding}")     //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project", ProjectRequest).Queries("name", "{name}").Queries("scaffolding", "{scaffolding}") //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project", ProjectRequest).Queries("id", "{id}")                                             //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project", ProjectRequest).Queries("name", "{name}")
	router.HandleFunc(baseURL+"/project", ProjectRequest).Queries("scaffolding", "{scaffolding}") //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/scaffolding", ProjectRequest)                             //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project", ProjectRequest)                                         //DELETE, POST, GET

	//Storage endpoint
	router.HandleFunc(baseURL+"/storage", storageRequest)

	//Profile endpoint
	router.HandleFunc(baseURL+"/user", ProfileRequest).Queries("id", "{id}")
	router.HandleFunc(baseURL+"/user", ProfileRequest).Queries("role", "{role}")
	router.HandleFunc(baseURL+"/user", ProfileRequest)

	//Gateway endpoint
	router.HandleFunc(baseURL+"/gateway", GatewayRequest).Queries("id", "{id}")
	router.HandleFunc(baseURL+"/gateway", GatewayRequest).Queries("projectName", "{projectName}")
	router.HandleFunc(baseURL+"/gateway", GatewayRequest).Queries("projectID", "{projectID}")
	router.HandleFunc(baseURL+"/gateway", GatewayRequest)

	//Gateway POST request endpoint (Only used for registering tags)
	router.HandleFunc(baseURL+"/gateway/input", UpdatePosition)

	http.Handle("/", router)
	fmt.Println("MQTT Server initializing...")
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
