package endpoints

import (
	"encoding/json"
	"net/http"
	"stillasTracker/api/Database"
	"stillasTracker/api/struct"
	"strconv"
)

/**
Class scaffolding
This class will contain all functions used for the handling of scaffolding units
The class contains the following functions:
	- addScaffolding:    Function lets a user add a scaffolding part to the system
	- deleteScaffolding: Function removes a scaffolding unit from the system
	- moveScaffold:      Function lets a user move scaffolding parts to a new project
	- getScaffoldingUnit Function returns information about a scaffolding part
	- getUnitHistory     Function returns the history of a scaffolding part

Version 0.1
Last modified Martin Iversen
*/
func scaffoldingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getPart(w, r) //Function returns all parts, one part or a selection of parts

	case http.MethodPost:
		createPart(w, r) //Function for adding new scaffolding parts to the system

	case http.MethodDelete:

	case http.MethodPut:

	}
}

/**
getPart function gets all scaffolding parts, some parts or one part
a user can search based on projects, id or type

*/
func getPart(w http.ResponseWriter, r *http.Request) {
	//var scaffoldingPart _struct.ScaffoldingType

	w.Header().Set("Content-Type", "application/json")

}

/**
createPart
function adds a list of scaffolding parts to the database
responds to a POST request with a body containing new scaffolding parts
*/
func createPart(w http.ResponseWriter, r *http.Request) {
	var scaffoldList _struct.AddScaffolding //Defines the structure of the body

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&scaffoldList) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//Prints the amount of scaffolding parts added to the system
	err = json.NewEncoder(w).Encode(strconv.Itoa(len(scaffoldList)) + " new scaffolding units added to the system \n the following units were added: \n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for i := range scaffoldList { //For loop iterates through the list of new scaffolding parts
		newPartPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(scaffoldList[i].Type).Doc(strconv.Itoa(scaffoldList[i].ID))
		var firebasePart map[string]interface{} //Defines the database structure for the new part

		part, err := json.Marshal(scaffoldList[i]) //Marshalls te body of the request into the right data format (byte)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.Unmarshal(part, &firebasePart) //Unmarshals the part object into the firebase part structure
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = Database.AddDocument(newPartPath, firebasePart) //Adds the part to the database
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(scaffoldList[i].Type + "\n")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
