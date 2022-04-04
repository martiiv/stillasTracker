package endpoints

import (
	"encoding/json"
	"google.golang.org/api/iterator"
	"net/http"
	"net/url"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	"stillasTracker/api/struct"
	"strconv"
	"strings"
)

/**
Class scaffolding
This class will contain all functions used for the handling of scaffolding units
The class contains the following functions:

	ScaffoldingRequest Function routes the request to the appropriate function
	getPart Handles all the get requests
	createPart Handles post requests
	deletePart Handles delete requests

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
	lastElement := getLastUrlElement(r)
	query := tool.GetQueryScaffolding(r)

	switch true {
	case "unit" == lastElement && len(query) > 1: //URL is on the following format: /stillastracking/v1/api/unit?type=""&id=""
		getIndividualScaffoldingPart(w, query)

	case "unit" == lastElement && len(query) == 1: //URL is on the following format: /stillastracking/v1/api/unit?type=""
		getScaffoldingByType(w, query)

	case "unit" == lastElement && len(query) == 0: //URL is on the following format: /stillastracking/v1/api/unit/
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
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	//Prints the amount of scaffolding parts added to the system
	err = json.NewEncoder(w).Encode(strconv.Itoa(len(scaffoldList)) + " new scaffolding units added to the system the following units were added:")
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}

	for i := range scaffoldList { //For loop iterates through the list of new scaffolding parts

		newPartPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldList[i].Type).Doc(strconv.Itoa(scaffoldList[i].ID))

		var firebasePart map[string]interface{} //Defines the database structure for the new part

		part, err := json.Marshal(scaffoldList[i]) //Marshalls te body of the request into the right data format (byte)
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		err = json.Unmarshal(part, &firebasePart) //Unmarshals the part object into the firebase part structure
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		err = database.AddDocument(newPartPath, firebasePart) //Adds the part to the database
		if err != nil {
			tool.HandleError(tool.DATABASEADDERROR, w)
			return
		}

		storagePath := database.Client.Collection(constants.P_LocationCollection).Doc(constants.P_StorageDocument).Collection(constants.P_Inventory).Doc(scaffoldList[i].Type)
		data, err := database.GetDocumentData(storagePath)
		if err != nil {
			tool.HandleError(tool.DATABASEREADERROR, w)
			return
		}

		oldQuantity, oldExpected := getQuantity(w, data)
		_, err = storagePath.Set(database.Ctx, map[string]interface{}{
			"type": scaffoldList[i].Type,
			"Quantity": map[string]interface{}{
				"expected":   oldQuantity + 1,
				"registered": oldExpected,
			}})
	}

	err = json.NewEncoder(w).Encode(scaffoldList)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
	}
}

func deletePart(w http.ResponseWriter, r *http.Request) {
	var deleteList _struct.DeleteScaffolding
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&deleteList)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
	}

	for i := range deleteList {
		objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(deleteList[i].Type).Doc(strconv.Itoa(deleteList[i].Id))
		err := database.DeleteDocument(objectPath)
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}
	}

	err = json.NewEncoder(w).Encode("All parts deleted successfully, number of parts deleted:")
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}

	err = json.NewEncoder(w).Encode(len(deleteList))
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}

}

//getQuantity Function takes an object of scaffolding type in storage and returns the expected amount and registered quantity
func getQuantity(w http.ResponseWriter, object map[string]interface{}) (int, int) {
	marshalled, err := json.Marshal(object)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return 0, 0
	}

	err = json.Unmarshal(marshalled, &_struct.Scaffolding{})
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return 0, 0
	}
	quantityString := string(marshalled)
	splitString := strings.Split(quantityString, ":")
	oldQuantity := strings.Split(splitString[2], ",")
	quantity, _ := strconv.Atoi(oldQuantity[0])

	oldExpected := strings.Split(splitString[0], "expected")
	expected, _ := strconv.Atoi(oldExpected[0])

	return quantity, expected
}

/**
getIndividualScaffoldingPart
Function takes the url and uses the passed in type and id to fetch a specific part from the database
URL Format: /stillastracking/v1/api/unit?type=""&?id=""
*/
func getIndividualScaffoldingPart(w http.ResponseWriter, query url.Values) {
	objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(query.Get("type")).Doc(query.Get("id"))

	part, err := database.GetDocumentData(objectPath)
	if err != nil {
		tool.HandleError(tool.DATABASEREADERROR, w)
		return
	}

	err = json.NewEncoder(w).Encode(part)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}
}

/**
getScaffoldingByType
Function takes the request URL, connects to the database and gets all the scaffolding parts
in the database with the passed in type
The url: /stillastracking/v1/api/unit/type=""
*/
func getScaffoldingByType(w http.ResponseWriter, query url.Values) {
	objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(query.Get("type")).Documents(database.Ctx)
	partList := database.GetCollectionData(objectPath)

	err := json.NewEncoder(w).Encode(partList)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
	}
}

/**
getAllScaffoldingParts
Function connects to the database and fetches all the parts in the database
URL format: /stillastracking/v1/api/unit/
*/
func getAllScaffoldingParts(w http.ResponseWriter, r *http.Request) {
	partPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collections(database.Ctx)
	for {
		scaffoldingType, err := partPath.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			tool.HandleError(tool.COLLECTIONITERATORERROR, w)
			return
		}

		document := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldingType.ID).Documents(database.Ctx)
		for {
			partRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			part, err := database.GetDocumentData(partRef.Ref)
			if err != nil {
				tool.HandleError(tool.DATABASEREADERROR, w)
				return
			}

			err = json.NewEncoder(w).Encode(part)
			if err != nil {
				tool.HandleError(tool.ENCODINGERROR, w)
				return
			}
		}
	}
}
