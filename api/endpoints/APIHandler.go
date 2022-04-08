package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const baseURL = "/stillastracking/v1/api"

// Handle /**
func Handle() {
	//fmt.Println("Listening on port" + getPort())
	/*//Scaffolding endpoints
	http.HandleFunc(baseURL+"/unit/", ScaffoldingRequest) //GET POST PUT DELETE
	//Project endpoint
	http.HandleFunc(baseURL+"/project/", ProjectRequest) //DELETE, POST, GET
	http.HandleFunc(baseURL+"/storage/", storageRequest)
	//Profile endpoint
	http.HandleFunc(baseURL+"/user/", ProfileRequest)
	log.Println(http.ListenAndServe(getPort(), nil))*/

	router := mux.NewRouter()
	router.HandleFunc("/exchange/v1/diag", diagnostics).Queries("limit", "{limit}")
	router.HandleFunc("/exchange/v1/diag/", diagnostics)

	http.Handle("/", router)
	log.Println(http.ListenAndServe(getPort(), router))

}

func diagnostics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryMap := mux.Vars(r)

	countryName := queryMap["limit"]

	//Prints the information
	fmt.Fprintf(w, `{
   	"exchangeratesapi": "HEi",
   "restcountries": "Dette",
   "version": "v1",
   "uptime": "%v" }`, countryName)
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
