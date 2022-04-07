package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"net/http"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
	"strconv"
	"strings"
)

/**
Class projects
This class will contain all data formatting and handling of the clients projects
Class contains the following functions:
	getProject:         The function returns information regarding an active project
	createProject:      The function lets a user create a project and assign scaffolding units as well as a geofence
	deleteProject:      The function deletes a project from the system
	getStorageFacility: The function returns the state of the storage facility (amount of scaffolding equipment)

Version 0.2
Last modified Aleksander Aaboen
TODO Find alternative for strings.Title since the function is deprecated
TODO If possible modularize the unmarshalling and encoding of lists since there is a lot of duplicate code doing this
*/

//ProjectCollection reference to a firestore document
var ProjectCollection *firestore.DocumentRef

/*
ProjectRequest
Function will redirect the user request to the appropriate endpoint
All communication with the project side of the API goes through this function
*/
func ProjectRequest(w http.ResponseWriter, r *http.Request) {
	//Sets headers to communicate with mobile and web application
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Defines the path to the project Documents in the firestore database
	ProjectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)
	requestType := r.Method

	//Redirects the request to the appropriate function based on the type of request
	switch requestType {
	case http.MethodGet:
		getProject(w, r)
	case http.MethodPost:
		createProject(w, r)
	case http.MethodPut:
		putRequest(w, r)
	case http.MethodDelete:
		deleteProject(w, r)
	}
}

//storageRequest will return all the scaffolding parts in the selected storage location.
func storageRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Predefines the necessary variables and structs for the function
	var storageArray []_struct.Scaffolding
	var storageArr []*firestore.DocumentIterator
	var storage *firestore.DocumentIterator

	query := r.URL.Query()

	//Checks if the query contains a scaffolding type
	if query.Has(constants.S_type) { //If it does, it gets the amount of scaffolding of that type from the database
		storage = database.Client.Collection(constants.P_LocationCollection+"/"+constants.P_StorageDocument+"/"+constants.P_Inventory).Where(constants.S_type, "==", query.Get(constants.S_type)).Documents(database.Ctx)
		storageArr = append(storageArr, storage)
	} else { //If not it gets all the scaffolding types in the storage section of the database
		storage = database.Client.Collection(constants.P_LocationCollection + "/" + constants.P_StorageDocument + "/" + constants.P_Inventory).Documents(database.Ctx)
	}

	for { //The function will iterate through the documents defined by the database path above
		doc, err := storage.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}

		var storageStruct _struct.Scaffolding  //Defines a struct with the structure of the Scaffolding type in the storage section of the database
		bytes, err := json.Marshal(doc.Data()) //Formats the data of the database document
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		err = json.Unmarshal(bytes, &storageStruct) //Unmarshal the document into the predefined struct on line 94
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}
		storageArray = append(storageArray, storageStruct) //Appends the storageStruct from line 94 to the array of the same type
	}

	err := json.NewEncoder(w).Encode(storageArray) //"Sends" the storageArray to the user
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}
}

/**
getProject will guide to the requested function.
The user will be redirected to either getProjectCollection or getProjectWithID.
If the user made an invalid request, the user will be redirected to invalidRequest.
*/
func getProject(w http.ResponseWriter, r *http.Request) {
	query, err := tool.GetQueryProject(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	switch true {
	case !query.Has(constants.P_idURL) && !query.Has(constants.P_nameURL) && !query.Has(constants.P_State):
		getProjectCollection(w, r) //If the query has keywords specific to the state of the project it ends up here
		break
	case query.Has(constants.P_idURL) || query.Has(constants.P_nameURL) || query.Has(constants.P_State):
		getProjectWithID(w, r) //If the query has keywords specific to the state of the project it ends up here
		break
	default:
		tool.InvalidRequest(w, r) //If the query has neither the request is invalid
		break
	}
}

/*
getProjectCollection will fetch every project in the database.
Uses getScaffoldingStruct in order to fetch all scaffolding parts associated with a project
*/
func getProjectCollection(w http.ResponseWriter, r *http.Request) {
	//Defines the necessary variables
	var projects []_struct.GetProject //Defines a list containing multiple GetProject structs
	collectionIterator := ProjectCollection.Collections(database.Ctx)
	queryMap := r.URL.Query()

	for { //For loop iterates through all the documents in the project collection and appends them
		collRef, err := collectionIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			tool.HandleError(tool.COLLECTIONITERATORERROR, w)
			return
		}
		document := ProjectCollection.Collection(collRef.ID).Documents(database.Ctx) //Gets the project collection from the database
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			var projectNew _struct.NewProject //Defines the struct of a new project
			doc, err := database.GetDocumentData(documentRef.Ref)
			if err != nil {
				tool.HandleError(tool.NODOCUMENTWITHID, w)
				return
			}

			projectByte, err := json.Marshal(doc)          //Formats the project data
			err = json.Unmarshal(projectByte, &projectNew) //Unmarshal the data into the struct defined on 165
			if err != nil {
				tool.HandleError(tool.UNMARSHALLERROR, w)
				return
			}

			var project _struct.GetProject  //Defines one instance of the GetProject struct, we have the project details but need scaffolding parts
			project.NewProject = projectNew //Defines the project section of the GetProject struct

			if queryMap.Has(constants.P_scaffolding) && queryMap.Get(constants.P_scaffolding) == "true" {
				scaffold, err := getScaffoldingStruct(documentRef.Ref) //If the request asks for scaffolding parts we get the scaffolding struct
				if err != nil {
					tool.HandleError(tool.COULDNOTFINDDATA, w)
					return
				}
				project.ScaffoldingArray = scaffold //We define the section of the GetProject struct which contains scaffolding parts and insert scaffolding parts
			}
			projects = append(projects, project)
		}
	}

	err := json.NewEncoder(w).Encode(projects)
	if err != nil {
		tool.HandleError(tool.NEWENCODERERROR, w)
		return
	}
}

/*
getProjectWithID will fetch a project based on the id
Uses IterateProjects to search through the database for a project containing the passed in ID
Uses getScaffoldingStruct in order to fetch all scaffolding parts associated with a project
*/
func getProjectWithID(w http.ResponseWriter, r *http.Request) {
	var documentReference []*firestore.DocumentRef
	var projects []_struct.GetProject
	var err error

	queryMap, _ := tool.GetQueryProject(r)
	if queryMap.Has(constants.P_idURL) { //If the query contains a project ID we convert it to int and store it for later use
		intID, err := strconv.Atoi(queryMap.Get(constants.P_idURL))
		if err != nil {
			tool.HandleError(tool.INVALIDBODY, w)
			return
		}

		documentReference, err = IterateProjects(intID, "", "") //Finds the project with the defined ID

	} else if queryMap.Has(constants.P_nameURL) { //If the url contains a project name we will use that instead
		documentReference, err = IterateProjects(0, queryMap.Get(constants.P_nameURL), "")
	} else { //If the url contains a state we will use that instead
		//Function strings.Title is deprecated however formatting makes usages of cases.Title impossible
		documentReference, err = IterateProjects(0, "", strings.Title(queryMap.Get(constants.P_State)))
	}
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}
	if documentReference == nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	for _, ref := range documentReference { //Iterates through the projects found in the if statement on line 214-228

		data, err := database.GetDocumentData(ref)
		if err != nil {
			tool.HandleError(tool.NODOCUMENTWITHID, w)
			return
		}

		jsonStr, err := json.Marshal(data)
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		var projectNew _struct.NewProject
		err = json.Unmarshal(jsonStr, &projectNew)
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		var project _struct.GetProject
		project.NewProject = projectNew
		if queryMap.Has(constants.P_scaffolding) && queryMap.Get(constants.P_scaffolding) == "true" {
			scaffold, err := getScaffoldingStruct(ref)
			if err != nil {
				tool.HandleError(tool.COULDNOTFINDDATA, w)
				return
			}
			project.ScaffoldingArray = scaffold

		}
		projects = append(projects, project)
	}

	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		tool.HandleError(tool.NEWENCODERERROR, w)
		return
	}
}

//deleteProject deletes selected projects from the database.
func deleteProject(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	var dat []map[string]interface{}
	err = json.Unmarshal(bytes, &dat)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	correctBody := true
	for _, m := range dat {
		_, correct := m[constants.P_idBody].(float64)
		if !correct {
			correctBody = false
		}
	}

	if !correctBody {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var deleteID _struct.IDStruct
	batch := database.Client.Batch()

	err = json.Unmarshal(bytes, &deleteID)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	for _, num := range deleteID {
		var correctID []*firestore.DocumentRef

		if num.ID != 0 {
			correctID, err = IterateProjects(num.ID, "", "")
		}

		if correctID == nil {
			tool.HandleError(tool.CouldNotDelete, w)
			return
		}

		subCollection := database.GetCollectionData(correctID[0].Collection(constants.P_StillasType).Documents(database.Ctx))
		if subCollection != nil {
			iter := correctID[0].Collection(constants.P_StillasType).Documents(database.Ctx)

			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					tool.HandleError(tool.DATABASEREADERROR, w)
					return
				}

				scaffoldingType, err := doc.DataAt(constants.S_type)

				scaffoldingParts, err := doc.DataAt(constants.P_QuantityExpected)
				byte, err := json.Marshal(scaffoldingParts)
				if err != nil {
					tool.HandleError(tool.MARSHALLERROR, w)
					return
				}

				var restScaffold int
				err = json.Unmarshal(byte, &restScaffold)
				if err != nil {
					tool.HandleError(tool.UNMARSHALLERROR, w)
					return
				}

				if restScaffold != 0 {
					typeScaffold := fmt.Sprint(scaffoldingType)

					move := _struct.InputScaffoldingWithID{
						ToProjectID:   0,
						FromProjectID: num.ID,
						InputScaffolding: _struct.InputScaffolding{
							{
								Type:     strings.Title(strings.ToLower(typeScaffold)),
								Quantity: restScaffold,
							},
						},
					}

					moveByte, err := json.Marshal(move)
					if err != nil {
						tool.HandleError(tool.MARSHALLERROR, w)
						return
					}

					read := strings.NewReader(string(moveByte))
					req, _ := http.NewRequest(http.MethodPut, "/stillastracking/v1/api/project/scaffolding/", read)
					transferProject(w, req)

				}

				batch.Delete(doc.Ref)
			}
		}

		batch.Delete(correctID[0])

	}
	_, err = batch.Commit(database.Ctx)
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

}

//createProject will create a Project and add it to the database
func createProject(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}
	correctBody := CheckStateBody(bytes)
	if !correctBody {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}
	var project _struct.NewProject
	err = json.Unmarshal(bytes, &project)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	id := strconv.Itoa(project.ProjectID)

	var correctID []*firestore.DocumentRef

	if project.ProjectID != 0 {
		correctID, err = IterateProjects(project.ProjectID, "", "")
	}

	if correctID != nil {
		tool.HandleError(tool.CouldNotAddSameID, w)
		return
	}

	state := project.State
	documentPath := ProjectCollection.Collection(state).Doc(id)

	var firebaseInput map[string]interface{}

	bytes, err = json.Marshal(project)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return
	}

	err = json.Unmarshal(bytes, &firebaseInput)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	batch := database.Client.Batch()
	batch.Create(documentPath, firebaseInput)
	err = addScaffolding(documentPath, batch)
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	} else {
		tool.HandleError(tool.ADDED, w)
	}
}

/*putRequest will guide to the requested function.
The user will be redirected to either updateState or transferProject.
If the user made an invalid request, the user will be redirected to invalidRequest.
*/
func putRequest(w http.ResponseWriter, r *http.Request) {
	lastElement := tool.GetLastUrlElement(r)

	switch true {
	case constants.P_scaffolding == lastElement:
		transferProject(w, r)
	case constants.P_projectURL == lastElement:
		updateState(w, r)
	default:
		tool.InvalidRequest(w, r)
	}

}

/*updateState will change the state of the project. In an atomic operation the project will change state,
be moved into the state collection and deleted form the old state collection.*/
func updateState(w http.ResponseWriter, r *http.Request) {
	batch := database.Client.Batch()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	correctBody := CheckStateBody(data)
	if !correctBody {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var stateStruct _struct.StateStruct
	err = json.Unmarshal(data, &stateStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	documentReference, err := IterateProjects(stateStruct.ID, "", "")
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	project, err := database.GetDocumentData(documentReference[0])
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	newPath := ProjectCollection.Collection(stateStruct.State).Doc(strconv.Itoa(stateStruct.ID))
	batch.Create(newPath, project)

	scaffolding, err := getScaffoldingStruct(documentReference[0])
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	scaffoldingMap, err := structToMap(scaffolding)

	for _, s := range scaffoldingMap {
		scaffoldingType := strings.Title((s[constants.P_typeField].(string)))
		batch.Create(newPath.Collection(constants.P_StillasType).Doc(scaffoldingType), s)
		batch.Delete(documentReference[0].Collection(constants.P_StillasType).Doc(scaffoldingType))
	}

	batch.Delete(documentReference[0])
	update := firestore.Update{
		Path:  constants.P_State,
		Value: stateStruct.State,
	}
	var updates []firestore.Update
	updates = append(updates, update)

	batch.Update(newPath, updates)

	_, err = batch.Commit(database.Ctx)
	if err != nil {
		tool.HandleError(tool.CHANGESWERENOTMADE, w)
		return
	}

}

func structToMap(input interface{}) ([]map[string]interface{}, error) {
	output, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var outputMap []map[string]interface{}
	err = json.Unmarshal(output, &outputMap)
	if err != nil {
		return nil, err
	}

	return outputMap, nil
}

/**
transferProject will move a project from one collection of a given state to another (active, inactive, upcoming).
This function will use batched writes to ensure integrity.
*/
func transferProject(w http.ResponseWriter, r *http.Request) {
	batch := database.Client.Batch()
	_, inputScaffolding, err := GetScaffoldingInput(r)
	if err != nil {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var fromLocation []interface{}
	var newLocation []interface{}
	var newPath []*firestore.DocumentRef
	var fromPaths []*firestore.DocumentRef

	switch inputScaffolding.FromProjectID {
	case 0:
		fromLocation, fromPaths = getScaffoldingFromStorage(inputScaffolding.InputScaffolding)
	default:
		fromLocation, fromPaths = getScaffoldingFromProject(inputScaffolding.FromProjectID, inputScaffolding.InputScaffolding)
	}
	if fromLocation == nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	switch inputScaffolding.ToProjectID {
	case 0:
		newLocation, newPath = getScaffoldingFromStorage(inputScaffolding.InputScaffolding)
	default:
		newLocation, newPath = getScaffoldingFromProject(inputScaffolding.ToProjectID, inputScaffolding.InputScaffolding)
	}
	if newLocation == nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	if len(fromLocation) != len(newLocation) {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	var sub = map[string]interface{}{}
	var add = map[string]interface{}{}

	for i := range fromLocation {
		quantity := inputScaffolding.InputScaffolding[i].Quantity

		intVar, err := tool.InterfaceToInt(fromLocation[i])
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		if quantity > intVar {
			tool.HandleError(tool.CANNOTTRANSFERESCAFFOLDS, w)
			return
		}

		inputQuantity, err := tool.InterfaceToInt(inputScaffolding.InputScaffolding[i].Quantity)
		fromLoc, err := tool.InterfaceToInt(fromLocation[i])
		newLoc, err := tool.InterfaceToInt(newLocation[i])
		if err != nil {
			tool.HandleError(tool.READALLERROR, w)
			return
		}

		sub[constants.P_Quantity] = map[string]interface{}{}
		sub[constants.P_Quantity].(map[string]interface{})[constants.P_Expected] = fromLoc - inputQuantity

		add[constants.P_Quantity] = map[string]interface{}{}
		add[constants.P_Quantity].(map[string]interface{})[constants.P_Expected] = newLoc + inputQuantity

		batch.Set(newPath[i], add, firestore.MergeAll)
		batch.Set(fromPaths[i], sub, firestore.MergeAll)
	}

	_, err = batch.Commit(database.Ctx)
	if err != nil {
		tool.HandleError(tool.CHANGESWERENOTMADE, w)
		return
	}
}

/**
getScaffoldingFromProject will get the selected scaffolding from a project, and return the number of scaffolding parts
and the documentref where the scaffolding are placed.
*/
func getScaffoldingFromProject(input int, scaffold _struct.InputScaffolding) ([]interface{}, []*firestore.DocumentRef) {

	var fromPath *firestore.DocumentRef
	var fromPaths []*firestore.DocumentRef
	var scaffoldingArray []interface{}

	newPath, _ := IterateProjects(input, "", "")
	if newPath == nil {
		return nil, nil
	}

	documentPath := tool.CreatePath(strings.Split(newPath[0].Path, "/")[5:])

	for _, s := range scaffold {
		iter := database.Client.Doc(documentPath).Collection(constants.P_StillasType).Where(constants.P_Type, "==", strings.ToLower(s.Type)).Documents(database.Ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, nil
			}

			scaffoldingFrom, err := doc.DataAt(constants.P_QuantityExpected)
			if err != nil {
				return nil, nil
			}
			fromPath = doc.Ref
			fromPaths = append(fromPaths, fromPath)
			scaffoldingArray = append(scaffoldingArray, scaffoldingFrom)
		}
	}

	return scaffoldingArray, fromPaths

}

/**
getScaffoldingFromStorage will get the selected scaffolding from a project, and return the number of scaffolding parts
and the documentref where the scaffolding are placed.
This function is used in
*/
func getScaffoldingFromStorage(scaffold _struct.InputScaffolding) ([]interface{}, []*firestore.DocumentRef) {

	var fromPath *firestore.DocumentRef
	var fromPaths []*firestore.DocumentRef
	var scaffoldingArray []interface{}

	for _, s := range scaffold {
		fromPath = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_StorageDocument + "/" + constants.P_Inventory + "/" + s.Type)
		fromPaths = append(fromPaths, fromPath)
	}

	for _, path := range fromPaths {
		doc, err := path.Get(database.Ctx)
		if err != nil {
			return nil, nil
		}

		scaffoldingFrom, err := doc.DataAt(constants.P_QuantityExpected)
		if err != nil {
			return nil, nil
		}

		scaffoldingArray = append(scaffoldingArray, scaffoldingFrom)

	}
	return scaffoldingArray, fromPaths

}

func addScaffolding(documentPath *firestore.DocumentRef, batch *firestore.WriteBatch) error {
	scaffoldingCollection := documentPath.Collection(constants.P_StillasType)

	for _, scaffoldingType := range constants.ScaffoldingTypes {
		scaffoldingTypeDocument := scaffoldingCollection.Doc(scaffoldingType)

		scaffoldingStruct := _struct.Scaffolding{
			Type: strings.ToLower(scaffoldingType),
			Quantity: _struct.Quantity{
				Expected:   0,
				Registered: 0,
			},
		}

		scaffoldByte, err := json.Marshal(scaffoldingStruct)
		if err != nil {
			return err
		}

		var scaffoldingInput map[string]interface{}
		err = json.Unmarshal(scaffoldByte, &scaffoldingInput)
		if err != nil {
			return err
		}

		batch.Set(scaffoldingTypeDocument, scaffoldingInput)
	}

	_, err := batch.Commit(database.Ctx)

	return err
}

//getScaffoldingStruct function will take a project document and return all the scaffolding parts associated with the project
func getScaffoldingStruct(document *firestore.DocumentRef) (_struct.ScaffoldingArray, error) {
	scaffolding := document.Collection(constants.P_StillasType).Documents(database.Ctx)
	var scaffoldArr _struct.ScaffoldingArray

	for { //Iterates through the section of the project document containing the associated scaffolding parts
		documentRef, err := scaffolding.Next()
		if err == iterator.Done {
			break
		}

		var scaffold _struct.Scaffolding //Defines format for individual scaffolding part
		doc, err := database.GetDocumentData(documentRef.Ref)
		if err != nil {
			return _struct.ScaffoldingArray{}, errors.New("no document")
		}

		projectByte, err := json.Marshal(doc)
		err = json.Unmarshal(projectByte, &scaffold)
		if err != nil {
			return _struct.ScaffoldingArray{}, errors.New("no document")
		}

		scaffoldArr = append(scaffoldArr, scaffold) //Appends data to the scaffolding array

	}
	return scaffoldArr, nil //Returns an array containing the list of scaffolding parts associated with a project
}
