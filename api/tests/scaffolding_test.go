package tests

import (
	"github.com/steinfletcher/apitest"
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
	router.HandleFunc("/stillastracking/v1/api/unit/", endpoints.ScaffoldingRequest)
	ts := httptest.NewServer(router)
	defer ts.Close()

	apitest.New().Handler(router).Get("/stillastracking/v1/api/unit/").
		Expect(t).
		Status(http.StatusOK).
		End()
}
