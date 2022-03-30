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
Class scaffoldingRequest
This class will contain all functions used for the handling of scaffolding units
The class contains the following functions:
	- addScaffolding:    Function lets a user add a scaffolding part to the system
	- deleteScaffolding: Function removes a scaffolding unit from the system
	- moveScaffold:      Function lets a user move scaffolding parts to a new project
	- getScaffoldingUnit Function returns information about a scaffolding part
	- getUnitHistory     Function returns the history of a scaffolding part
TODO Error handle properly
TODO update file head comment
Version 0.1
Last modified Martin Iversen
*/

//ScaffoldingRequest Function redirects the user to different parts of the scaffolding class
func ScaffoldingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getPart(w, r) //Function returns all parts, one part or a selection of parts

	case http.MethodPost:
		createPart(w, r) //Function for adding new scaffolding parts to the system

	case http.MethodDelete:
		deletePart(w, r) //Function for deleting scaffolding part

	case http.MethodPut:

	}
}

/**
getPart function gets all scaffolding parts, some parts or one part
a user can search based on projects, id or type
*/
func getPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := r.URL.Path //Defining the url and splitting it on the symbol: /
	splitUrl := strings.Split(url, "/")

	switch len(splitUrl) {
	case 8: //Case 8 means that the URL is on the following format: /stillastracking/v1/api/unit/?type=""/?id=""/ TODO Formater URL riktig
		getIndividualScaffoldingPart(w, r, splitUrl)

	case 7: //Case 4 means that a type of scaffolding is wanted however, not a specific one since no ID is passed in
		getScaffoldingByType(w, r, splitUrl)

	case 6: //Case 3 means that the user wants all the scaffolding parts in the database
		getAllScaffoldingParts(w, r)
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
	err = json.NewEncoder(w).Encode(strconv.Itoa(len(scaffoldList)) + " new scaffolding units added to the system the following units were added:")
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
	}

	err = json.NewEncoder(w).Encode(scaffoldList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	err = json.NewEncoder(w).Encode("All parts deleted successfully, number of parts deleted:")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(len(deleteList))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/**
getIndividualScaffoldingPart
Function takes the url and uses the passed in type and id to fetch a specific part from the database
URL Format: /stillastracking/v1/api/unit?type=""?id="" TODO Format url correctly
*/
func getIndividualScaffoldingPart(w http.ResponseWriter, r *http.Request, URL []string) {
	objectPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(URL[5]).Doc(URL[6])

	part, err := Database.GetDocumentData(objectPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(part)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

/**
getScaffoldingByType
Function takes the request URL, connects to the database and gets all the scaffolding parts
in the database with the passed in type
The url: /stillastracking/v1/api/unit/type= TODO Configure URL properly with variables
*/
func getScaffoldingByType(w http.ResponseWriter, r *http.Request, URL []string) {
	objectPath := Database.Client.Collection("TrackingUnit").Doc("ScaffoldingParts").Collection(URL[5]).Documents(Database.Ctx)
	partList := Database.GetCollectionData(objectPath)

	err := json.NewEncoder(w).Encode(partList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

/**
getAllScaffoldingParts
Function connects to the database and fetches all the parts in the database
URL format: /stillastracking/v1/api/unit/
*/
func getAllScaffoldingParts(w http.ResponseWriter, r *http.Request) {
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
