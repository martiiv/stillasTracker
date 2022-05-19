package endpoints

import (
	"encoding/json"
	"google.golang.org/api/iterator"
	"net/http"
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
Version 1.0
Last modified Martin Iversen 07.04.2022
TODO make type non case sensitive
*/

/*
ScaffoldingRequest Function forwards all requests to the appropriate function
Uses getPart, createPart and deletePart
*/
func ScaffoldingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	switch r.Method {
	case http.MethodGet:
		getPart(w, r) //Function returns all parts, one part or a selection of parts

	case http.MethodPost:
		createPart(w, r) //Function for adding new scaffolding parts to the system

	case http.MethodDelete:
		deletePart(w, r) //Function for deleting scaffolding part}
	}
}

/**
getPart function gets all scaffolding parts, some parts or one part
a user can search based on projects, id or type
*/
func getPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	queries, _ := tool.GetQueryScaffolding(r)

	switch true {
	case queries["type"] != "" && queries["id"] != "": //URL is on the following format: /stillastracking/v1/api/unit?type=""&id=""
		getIndividualScaffoldingPart(w, queries["type"], queries["id"])

	case queries["type"] != "": //URL is on the following format: /stillastracking/v1/api/unit?type=""
		getScaffoldingByType(w, queries["type"])

	default: //URL is on the following format: /stillastracking/v1/api/unit/
		getAllScaffoldingParts(w)
	}
}

/**
createPart
function adds a list of scaffolding parts to the database
responds to a POST request with a body containing new scaffolding parts
*/
func createPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var scaffoldList _struct.AddScaffolding //Defines the structure of the body

	err := json.NewDecoder(r.Body).Decode(&scaffoldList) //Decodes the requests body into the structure defined above
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	for i := range scaffoldList { //For loop iterates through the list of new scaffolding parts

		newPartPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldList[i].Type).Doc(scaffoldList[i].Id)

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
		} else {
			tool.HandleError(tool.ADDED, w)
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var deleteList _struct.DeleteScaffolding

	err := json.NewDecoder(r.Body).Decode(&deleteList)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
	}

	for i := range deleteList {
		objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(deleteList[i].Type).Doc(deleteList[i].Id)
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
	//Splits the object into
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
func getIndividualScaffoldingPart(w http.ResponseWriter, scaffoldType string, scaffoldId string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldType).Doc(scaffoldId)

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
func getScaffoldingByType(w http.ResponseWriter, scaffoldingType string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldingType).Documents(database.Ctx)
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
func getAllScaffoldingParts(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var scaffoldList []_struct.ScaffoldingType
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

			var scaffoldPart _struct.ScaffoldingType
			partByte, err := json.Marshal(part)
			if err != nil {
				tool.HandleError(tool.MARSHALLERROR, w)
				return
			}

			err = json.Unmarshal(partByte, &scaffoldPart)
			if err != nil {
				tool.HandleError(tool.UNMARSHALLERROR, w)
				return
			}
			scaffoldList = append(scaffoldList, scaffoldPart)
		}
	}

	err := json.NewEncoder(w).Encode(scaffoldList)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}
}
