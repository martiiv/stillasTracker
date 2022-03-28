package tests

import (
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
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
	r := http.NewServeMux()
	r.HandleFunc("http://localhost:8080/stillastracking/v1/api/unit/", endpoints.ScaffoldingRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get("http://localhost:8080/stillastracking/v1/api/unit/Flooring/321/")
		if err != nil {
			t.Errorf("Expected %d, recieved %d", http.StatusOK, res.StatusCode)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}

func Test_getScaffoldingAPITEST(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("http://localhost:8080/stillastracking/v1/api/unit/Flooring/321/", endpoints.ScaffoldingRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("not found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("http://localhost:8080/stillastracking/v1/api/unit/Flooring/321/").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("http://localhost:8080/stillastracking/v1/api/unit/Flooring/321/").
			Expect(t).
			Assert(jsonpath.Equal(`batteryLevel`, `100`)).
			Status(http.StatusOK).
			End()
	})
}
