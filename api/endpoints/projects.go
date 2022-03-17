package endpoints

import "net/http"

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
		getProjectDetails(w, r)
	case "POST":
	case "PUT":
	case "DELETE":
	}
}

func storageRequest(w http.ResponseWriter, r *http.Request) {

}

func getProjectDetails(w http.ResponseWriter, r *http.Request) {

}
