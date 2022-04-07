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
	router.Path(baseURL + "/unit").HandlerFunc(ScaffoldingRequest)                                               //GET POST PUT DELETE
	router.Path(baseURL+"/unit").HandlerFunc(ScaffoldingRequest).Queries("type", "{type}")                       //GET POST PUT DELETE
	router.Path(baseURL+"/unit").HandlerFunc(ScaffoldingRequest).Queries("type", "{type}").Queries("id", "{id}") //GET POST PUT DELETE

	//Project endpoint
	router.HandleFunc(baseURL+"/project/", ProjectRequest)                                  //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/", ProjectRequest).Queries()                        //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/", ProjectRequest).Queries().Queries()              //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/{scaffolding}", ProjectRequest)                     //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/{scaffolding}", ProjectRequest).Queries()           //DELETE, POST, GET
	router.HandleFunc(baseURL+"/project/{scaffolding}", ProjectRequest).Queries().Queries() //DELETE, POST, GET
	//Storage endpoint
	router.HandleFunc(baseURL+"/storage/", storageRequest)
	router.HandleFunc(baseURL+"/storage/", storageRequest).Queries()

	//Profile endpoint
	router.HandleFunc(baseURL+"/user/", ProfileRequest)
	router.HandleFunc(baseURL+"/user/", ProfileRequest).Queries()
	router.HandleFunc(baseURL+"/user/", ProfileRequest).Queries()

	http.Handle("/", router)
	log.Println(http.ListenAndServe(getPort(), nil))
}

func oldHandle() {
	fmt.Println("Listening on port" + getPort())
	//Scaffolding endpoint
	http.HandleFunc(baseURL+"/unit/", ScaffoldingRequest) //GET POST PUT DELETE
	//Project endpoint
	http.HandleFunc(baseURL+"/project/", ProjectRequest) //DELETE, POST, GET
	//Storage endpoint
	http.HandleFunc(baseURL+"/storage/", storageRequest)
	//Profile endpoint
	http.HandleFunc(baseURL+"/user/", ProfileRequest)
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
