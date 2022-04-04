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
func Test_ProjectAPITEST(t *testing.T) {
	dataBaseTestConnection()
	handler := http.HandlerFunc(endpoints.ProjectRequest)
	input := `{"projectID":2,"projectName":"CCGjovik","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"size":980,"state":"Upcoming","address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}}}`
	t.Run("Add new Projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/project/").
			Body(input).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusCreated).
			End()
	})

	input2 := `{"projectID":7,"projectName":"CCGjovik","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"size":980,"state":"Upcoming","address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}}}`
	t.Run("Add new Projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/project/").
			Body(input2).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusCreated).
			End()
	})

	input3 := `{"projectID":3,"projectName":"Fjellhallen","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"size":980,"state":"Upcoming","address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}}}`
	t.Run("Add new Projects, with existing id", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/project/").
			Body(input3).
			Expect(t).
			Body("id is already in use\n"). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	wrongInput := `{"ds":7,"projectme":"NTNU","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18.03.2022","endDate":"30.05.2023"},"size":980,"state":"Upcoming","address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}}}`
	t.Run("Add new Projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Post("/stillastracking/v1/api/project/").
			Body(wrongInput).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	outputAll := `[{"projectID":3,"projectName":"Fjellhallen","size":980,"state":"Upcoming","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}},"scaffolding":null},{"projectID":7,"projectName":"CCGjovik","size":980,"state":"Upcoming","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}},"scaffolding":null}]`
	t.Run("Get all projects", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project/").
			Body(""). //Fill in request body here
			Expect(t).
			Body(outputAll). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	output7 := `[{"projectID":7,"projectName":"CCGjovik","size":980,"state":"Upcoming","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}},"scaffolding":null}]`
	t.Run("Get project by id", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project").Query("id", "7").
			Body(""). //Fill in request body here
			Expect(t).
			Body(output7). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get project by id, wrong id", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project").Query("id", "4213").
			Body(""). //Fill in request body here
			Expect(t).
			Body("document does not exist\n"). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	output3 := `[{"projectID":7,"projectName":"CCGjovik","size":980,"state":"Upcoming","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}},"scaffolding":null}]`
	t.Run("Get individual project", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("stillastracking/v1/api/project").Query("name", "CCGjovik").
			Body(""). //Fill in request body here
			Expect(t).
			Body(output3). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get project by name, wrong name", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project").Query("name", "mjosa").
			Body(""). //Fill in request body here
			Expect(t).
			Body("document does not exist\n"). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Get project by query, no valid query", func(t *testing.T) {
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project").Query("size", "432").
			Body(""). //Fill in request body here
			Expect(t).
			Body("invalid request\n"). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Move scaffolding parts from storage to new project", func(t *testing.T) {
		inputBodyMove1 := `{"toProjectID": 3,"fromProjectID": 0,"scaffold":[{"type": "Spire","quantity": 200},{"type": "Bunnskrue","quantity": 300}]}`
		apitest.New().
			HandlerFunc(handler).
			Put("/stillastracking/v1/api/project/scaffolding/").
			Body(inputBodyMove1). //Fill in request body here
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Get project with updated scaffolding", func(t *testing.T) {
		scaffoldingHolding := `[{"projectID":3,"projectName":"Fjellhallen","size":980,"state":"Upcoming","latitude":60.7905060889568,"longitude":10.681777071532371,"period":{"startDate":"18-03-2022","endDate":"30-05-2023"},"address":{"street":"Jernbanesvingen6","zipcode":"2821","municipality":"Gjovik","county":"Innlandet"},"customer":{"name":"CCGjovik","number":61130410,"email":"gjovik@cc.no"},"geofence":{"w-position":{"latitude":60.79077759591496,"longitude":10.683249543160402},"x-position":{"latitude":60.79015256651516,"longitude":10.684424851812308},"y-position":{"latitude":60.789159847696716,"longitude":10.68094413003551},"z-position":{"latitude":60.78963782726421,"longitude":10.680160590934236}},"scaffolding":[{"type":"bunnskrue","Quantity":{"expected":300,"registered":0}},{"type":"diagonalstang","Quantity":{"expected":0,"registered":0}},{"type":"enrørsbjelke","Quantity":{"expected":0,"registered":0}},{"type":"gelender","Quantity":{"expected":0,"registered":0}},{"type":"lengdebjelke","Quantity":{"expected":0,"registered":0}},{"type":"plank","Quantity":{"expected":0,"registered":0}},{"type":"rekkverksramme","Quantity":{"expected":0,"registered":0}},{"type":"spire","Quantity":{"expected":200,"registered":0}},{"type":"stillaslem","Quantity":{"expected":0,"registered":0}},{"type":"trapp","Quantity":{"expected":0,"registered":0}}]}]`
		apitest.New().
			HandlerFunc(handler).
			Get("/stillastracking/v1/api/project").Query("id", "3").Query("scaffolding", "true").
			Body(""). //Fill in request body here
			Expect(t).
			Body(scaffoldingHolding). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Delete all project", func(t *testing.T) {
		deleteBody := `[{"id": 3}, {"id": 7}]`
		apitest.New().
			HandlerFunc(handler).
			Delete("/stillastracking/v1/api/project/").
			Body(deleteBody).
			Expect(t).
			Body(""). //Fill in response body here in raw format
			Status(http.StatusOK).
			End()
	})

	t.Run("Delete a non existing project", func(t *testing.T) {
		deleteBody := `[{"id": 64}]`
		apitest.New().
			HandlerFunc(handler).
			Delete("/stillastracking/v1/api/project/").
			Body(deleteBody).
			Expect(t).
			Body("invalid id, could not delete\n"). //Fill in response body here in raw format
			Status(http.StatusBadRequest).
			End()
	})
}
