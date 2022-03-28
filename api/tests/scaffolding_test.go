package tests

import (
	"net/http"
	"net/http/httptest"
	"stillasTracker/api/endpoints"
	"testing"
)

/**
Test_getScaffolding
Function for testing the get scaffolding endpoint
Creates a server and router before sending the request and getting the desired statuscode and output
*/
func Test_getScaffolding(t *testing.T) {

	router := http.NewServeMux()
	router.HandleFunc("/stillastracking/v1/api/unit/Flooring/321", endpoints.ScaffoldingRequest)
	ts := httptest.NewServer(router)
	defer ts.Close()

	//TODO sjekke JSON format under her
	//apitest.New().Handler(router).Get("/stillastracking/v1/api/unit/").
}
