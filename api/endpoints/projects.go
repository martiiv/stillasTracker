package endpoints

import (
	"fmt"
	"net/http"
)
import (
	"stillasTracker/api/Database"
)

/**
Class projects
This class will contain all data formating and handling of the clients projects
Class contains the following functions:
	- getProject:         The function returns information regarding an active project
	- createProject:      The function lets a user create a project and assign scaffolding units as well as a geofence
	- deleteProject:      The function deletes a project from the system
	- getStorageFacility: The function returns the state of the storage facility (amount of scaffolding equipment)

Version 0.1
Last modified Martin Iversen
*/

/**
Main function to switch between the different request types.
*/
func projectRequest(w http.ResponseWriter, r *http.Request) {

	requestType := r.Method
	switch requestType {
	case "GET":
		fmt.Println("Hello World")
	case "POST":
		createProject(w, r)
	case "PUT":
	case "DELETE":
	}
}

func storageRequest(w http.ResponseWriter, r *http.Request) {

}

func createProject(w http.ResponseWriter, r *http.Request) {
	documentPath := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc("1")

	project := map[string]interface{}{
		"capital": true,
	}

	Database.AddDocument(documentPath, project)
}
