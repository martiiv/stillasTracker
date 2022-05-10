package tests

import (
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"net/http"
	"net/http/httptest"
	"stillasTracker/api/endpoints"
	"testing"
)

/**
Test_getScaffolding
Function for testing the scaffolding endpoint
*/
func Test_ScaffoldingAPITEST(t *testing.T) {
	dataBaseTestConnection()
	r := mux.NewRouter()

	//Add list of Scaffoldingparts which sends a post request and creates 12 scaffolding parts
	t.Run("Add list of Scaffoldingparts", func(t *testing.T) {
		r.HandleFunc("/stillastracking/v1/api/unit/", endpoints.ScaffoldingRequest)
		ts := httptest.NewServer(r)
		defer ts.Close()
		apitest.New().
			Handler(r).
			Post("/stillastracking/v1/api/unit/").
			Body(`[ { "id": 1, "type": "Spire", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address":null } }, { "id": 2, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 3, "type": "Short-Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 4, "type": "Staircase", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 5, "type": "Bottom-Screw", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 6, "type": "Diagonals", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 7, "type": "Beam1", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 8, "type": "Beam2", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 9, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 10, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 11, "type": "Spire", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 12, "type": "Railing", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } } ]`).
			Expect(t).
			Body("\"12 new scaffolding units added to the system the following units were added:\"\n[{\"id\":1,\"type\":\"Spire\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":2,\"type\":\"Flooring\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":3,\"type\":\"Short-Flooring\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":4,\"type\":\"Staircase\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":5,\"type\":\"Bottom-Screw\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":6,\"type\":\"Diagonals\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":7,\"type\":\"Beam1\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":8,\"type\":\"Beam2\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":9,\"type\":\"Flooring\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":10,\"type\":\"Flooring\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":11,\"type\":\"Spire\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}},{\"id\":12,\"type\":\"Railing\",\"batteryLevel\":100,\"location\":{\"longitude\":null,\"latitude\":null,\"address\":null}}]\n").
			Status(http.StatusOK).
			End()
	})

	//Gets all Scaffolding parts
	t.Run("Get all Scaffolding parts", func(t *testing.T) {
		r.HandleFunc("/stillastracking/v1/api/unit/", endpoints.ScaffoldingRequest)
		ts := httptest.NewServer(r)
		defer ts.Close()
		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Get("/stillastracking/v1/api/unit/").
			Expect(t).
			Body("{\"batteryLevel\":100,\"id\":7,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Beam1\"}\n{\"batteryLevel\":100,\"id\":8,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Beam2\"}\n{\"batteryLevel\":100,\"id\":5,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Bottom-Screw\"}\n{\"batteryLevel\":100,\"id\":6,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Diagonals\"}\n{\"batteryLevel\":100,\"id\":10,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"}\n{\"batteryLevel\":100,\"id\":2,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"}\n{\"batteryLevel\":100,\"id\":9,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"}\n{\"batteryLevel\":100,\"id\":12,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Railing\"}\n{\"batteryLevel\":100,\"id\":3,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Short-Flooring\"}\n{\"batteryLevel\":100,\"id\":1,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Spire\"}\n{\"batteryLevel\":100,\"id\":11,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Spire\"}\n{\"batteryLevel\":100,\"id\":4,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Staircase\"}\n").
			Status(http.StatusOK).
			End()
	})

	//Gets all Scaffolding parts by Flooring
	t.Run("Get Scaffolding by type", func(t *testing.T) {
		r.HandleFunc("/stillastracking/v1/api/unit", endpoints.ScaffoldingRequest).Queries()
		ts := httptest.NewServer(r)
		defer ts.Close()
		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Get("/stillastracking/v1/api/unit").
			Query("type", "Flooring").
			Expect(t).
			Body("[{\"batteryLevel\":100,\"id\":10,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"},{\"batteryLevel\":100,\"id\":2,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"},{\"batteryLevel\":100,\"id\":9,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"}]\n").
			Status(http.StatusOK).
			End()

		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Get("/stillastracking/v1/api/unit").
			Query("type", "Beam1").
			Expect(t).
			Body("[{\"batteryLevel\":100,\"id\":7,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Beam1\"}]\n").
			Status(http.StatusOK).
			End()

		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Get("/stillastracking/v1/api/unit").
			Query("type", "Staircase").
			Expect(t).
			Body("[{\"batteryLevel\":100,\"id\":4,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Staircase\"}]\n").
			Status(http.StatusOK).
			End()
	})

	//Gets a specific scaffolding part
	t.Run("Get Individual Scaffolding part", func(t *testing.T) {
		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Get("/stillastracking/v1/api/unit").
			Query("type", "Flooring").
			Query("id", "9").
			Expect(t).
			Body("{\"batteryLevel\":100,\"id\":9,\"location\":{\"address\":null,\"latitude\":null,\"longitude\":null},\"type\":\"Flooring\"}\n").
			Status(http.StatusOK).
			End()
	})

	t.Run("Delete all scaffolding parts", func(t *testing.T) {
		apitest.New().
			HandlerFunc(endpoints.ScaffoldingRequest).
			Delete("/stillastracking/v1/api/unit/").
			Body("[ { \"id\": 7, \"type\": \"Beam1\" }, { \"id\": 8, \"type\": \"Beam2\" }, { \"id\": 5, \"type\": \"Bottom-Screw\" }, { \"id\": 6, \"type\": \"Diagonals\" }, { \"id\": 10, \"type\": \"Flooring\" }, { \"id\": 2, \"type\": \"Flooring\" }, { \"id\": 9, \"type\": \"Flooring\" }, { \"id\": 12, \"type\": \"Railing\" }, { \"id\": 3, \"type\": \"Short-Flooring\" }, { \"id\": 1, \"type\": \"Spire\" }, { \"id\": 11, \"type\": \"Spire\" }, { \"id\": 4, \"type\": \"Staircase\" } ]").
			Expect(t).
			Body("\"All parts deleted successfully, number of parts deleted:\"\n12\n").
			Status(http.StatusOK).
			End()
	})
}
