package endpoints

import (
	"encoding/json"
	"google.golang.org/api/iterator"
	"net/http"
	"stillasTracker/api/Database"
	"stillasTracker/api/struct"
	"strconv"
	"strings"
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
		deletePart(w, r)
	case http.MethodPut:

	}
}

/**
getPart function gets all scaffolding parts, some parts or one part
a user can search based on projects, id or type
*/
func getPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.Path //Defining the url and splitting it on /
	splitUrl := strings.Split(url, "/")

	print(len(splitUrl))

	switch len(splitUrl) {
	case 8: //Case 5 means that only an id is passed in the URL, we return one spesific scaffolding part with the id
		objectPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(splitUrl[4]).Doc(splitUrl[5])

		for i := 0; i < len(splitUrl); i++ {
			err := json.NewEncoder(w).Encode(splitUrl[i])
			if err != nil {
			}
		}

		part, err := Database.GetDocumentData(objectPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	case 7: //Case 4 means that a type of scaffolding is wanted however, not a specific one since no ID is passed in
		objectPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(splitUrl[4]).Documents(Database.Ctx)
		partList := Database.GetCollectionData(objectPath)

		err := json.NewEncoder(w).Encode(partList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	case 6: //Case 3 means that the user wants all the scaffolding parts int the database
		partPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collections(Database.Ctx)
		for {
			scaffoldingType, err := partPath.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusNoContent)
				break
			}
			document := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(scaffoldingType.ID).Documents(Database.Ctx)
			for {
				partRef, err := document.Next()
				if err == iterator.Done {
					break
				}

				part, err := Database.GetDocumentData(partRef.Ref)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNoContent)
				}

				err = json.NewEncoder(w).Encode(part)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
			}
		}

	}
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

func deletePart(w http.ResponseWriter, r *http.Request) {
	var deleteList _struct.DeleteScaffolding
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&deleteList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for i := range deleteList {
		objectPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(deleteList[i].Type).Doc(strconv.Itoa(deleteList[i].Id))
		err := Database.DeleteDocument(objectPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
