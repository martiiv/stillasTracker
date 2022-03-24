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

	case http.MethodPost:
		createPart(w, r)

	case http.MethodDelete:

	case http.MethodPut:

	}
}

func createPart(w http.ResponseWriter, r *http.Request) {
	var scaffoldList _struct.AddScaffolding

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&scaffoldList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	print("%v new scaffolding units added to the system \n "+"the following units were added:", len(scaffoldList))

	for i := range scaffoldList {
		newPartPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(scaffoldList[i].Type).Doc(strconv.Itoa(scaffoldList[i].ID))

		var firebasePart map[string]interface{}
		part, err := json.Marshal(scaffoldList[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.Unmarshal(part, &firebasePart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		Database.AddDocument(newPartPath, firebasePart)
		print(scaffoldList[i].Type + "\n")
	}
}
