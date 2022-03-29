package tests

import (
	"github.com/steinfletcher/apitest"
	"net/http"
	"stillasTracker/api/endpoints"
	"testing"
)

/**
Test_getScaffolding
Function for testing the get scaffolding endpoint
Creates a server and router before sending the request and getting the desired statuscode and output
*/
func Test_ScaffoldingAPITEST(t *testing.T) {
	dataBaseTestConnection()
	handler := http.HandlerFunc(endpoints.ScaffoldingRequest)

	t.Run("Add list of Scaffoldingparts", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/unit/").
			Body(`[ { "id": 1, "type": "Spire", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address":null } }, { "id": 2, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 3, "type": "Short-Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 4, "type": "Staircase", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 5, "type": "Bottom-Screw", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 6, "type": "Diagonals", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 7, "type": "Beam1", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 8, "type": "Beam2", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 9, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 10, "type": "Flooring", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 11, "type": "Spire", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } }, { "id": 12, "type": "Railing", "batteryLevel": 100, "location": { "longitude": null, "latitude": null, "address": null } } ]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Get all Scaffolding parts", func(t *testing.T) {

	})

	t.Run("Get Scaffolding by type", func(t *testing.T) {

	})

	t.Run("Get Individual Scaffolding part", func(t *testing.T) {

	})
}
