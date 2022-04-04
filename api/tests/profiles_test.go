package tests

import (
	"github.com/steinfletcher/apitest"
	"net/http"
	"stillasTracker/api/endpoints"
	"testing"
)

func Test_ProfileAPTEST(t *testing.T) {
	dataBaseTestConnection()
	handler := http.HandlerFunc(endpoints.ProfileRequest)
	inputUser := `{"employeeID": 32,"name": {"firstName": "Ola", "lastName": "Nordmann"},"dateOfBirth": "28-03-1999","role": "admin","phone": 98326534,"email": "olanordmann@mail.com","admin": true}`
	t.Run("Add new Profile", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/user/").
			Body(inputUser).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusCreated).
			End()
	})

	inputUser2 := `{"employeeID": 1,"name": {"firstName": "Kari", "lastName": "Jensen"},"dateOfBirth": "28-03-1988","role": "installer","phone": 98326534,"email": "karinordmann@mail.com","admin": false}`
	t.Run("Add new Profile", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/user/").
			Body(inputUser2).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusCreated).
			End()
	})

	t.Run("Get profile all", func(t *testing.T) {
		outputAll := `[{"employeeID":32,"name":{"firstName":"Ola","lastName":"Nordmann"},"dateOfBirth":"28-03-1999","role":"admin","phone":98326534,"email":"olanordmann@mail.com","admin":true,"projects":null},{"employeeID":1,"name":{"firstName":"Kari","lastName":"Jensen"},"dateOfBirth":"28-03-1988","role":"installer","phone":98326534,"email":"karinordmann@mail.com","admin":false,"projects":null}]`
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user/").
			Body("").
			Expect(t).
			Body(outputAll). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get profile by name", func(t *testing.T) {
		outputName := `[{"employeeID":32,"name":{"firstName":"Ola","lastName":"Nordmann"},"dateOfBirth":"28-03-1999","role":"admin","phone":98326534,"email":"olanordmann@mail.com","admin":true,"projects":null}]`
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user").Query("name", "Nordmann").
			Body("").
			Expect(t).
			Body(outputName). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get profile by id", func(t *testing.T) {
		outputID := `{"employeeID":1,"name":{"firstName":"Kari","lastName":"Jensen"},"dateOfBirth":"28-03-1988","role":"installer","phone":98326534,"email":"karinordmann@mail.com","admin":false,"projects":null}`
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user").Query("id", "1").
			Body("").
			Expect(t).
			Body(outputID). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get profile by empty id", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user").Query("id", "421").
			Body("").
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusNoContent).
			End()
	})

	t.Run("Get profile by role", func(t *testing.T) {
		outputRole := `[{"employeeID":32,"name":{"firstName":"Ola","lastName":"Nordmann"},"dateOfBirth":"28-03-1999","role":"admin","phone":98326534,"email":"olanordmann@mail.com","admin":true,"projects":null}]`
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user").Query("role", "admin").
			Body("").
			Expect(t).
			Body(outputRole). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get profile by invalid role", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/user").Query("role", "shopper").
			Body("").
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	//Todo add put and delete test

	t.Run("Update profile", func(t *testing.T) {
		update := "{\"employeeID\": 32,\"admin\": false}"
		apitest.New().
			HandlerFunc(handler).
			Put("/stillastracking/v1/api/user/").
			Body("").
			Expect(t).
			Body(update). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

}
