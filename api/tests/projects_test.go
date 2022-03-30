package tests

import (
	"github.com/steinfletcher/apitest"
	"net/http"
	"stillasTracker/api/endpoints"
	"testing"
)

/**
Test_ProjectAPITEST
Function for testing the project endpoint
*/
func Test_PorjectAPITEST(t *testing.T) {
	dataBaseTestConnection()
	handler := http.HandlerFunc(endpoints.ProjectRequest)

	t.Run("Add new Projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get all projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get project by something", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get individual project", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Move scaffolding parts to new project", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Put("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Delete all projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Delete("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})
}
