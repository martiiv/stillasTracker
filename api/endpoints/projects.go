package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"log"
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
This class will contain all data formating and handling of the clients projects
Class contains the following functions:
	- getProject:         The function returns information regarding an active project
	- createProject:      The function lets a user create a project and assign scaffolding units as well as a geofence
	- deleteProject:      The function deletes a project from the system
	- getStorageFacility: The function returns the state of the storage facility (amount of scaffolding equipment)

Version 0.2
Last modified Aleksander Aaboen
*/
var projectCollection *firestore.DocumentRef

/**
ProjectRequest
Main function to switch between the different request types.
*/
func ProjectRequest(w http.ResponseWriter, r *http.Request) {

	projectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)
	requestType := r.Method

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

/*
getLastUrlElement will split the url and return the last element.
*/
func getLastUrlElement(r *http.Request) string {
	url := r.URL.Path
	trimmedURL := strings.TrimRight(url, "/")
	splittedURL := strings.Split(trimmedURL, "/")
	lastElement := splittedURL[len(splittedURL)-1]
	return lastElement
}

//storageRequest will return all the scaffolding parts in the selected storage location.
func storageRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storageArray []_struct.Scaffolding
	var storageArr []*firestore.DocumentIterator
	query := getQuery(r)
	var storage *firestore.DocumentIterator
	if query.Has(constants.S_type) {
		storage = database.Client.Collection(constants.P_LocationCollection+"/"+constants.P_StorageDocument+"/"+constants.P_Inventory).Where(constants.S_type, "==", query.Get(constants.S_type)).Documents(database.Ctx)
		storageArr = append(storageArr, storage)
	} else {
		storage = database.Client.Collection(constants.P_LocationCollection + "/" + constants.P_StorageDocument + "/" + constants.P_Inventory).Documents(database.Ctx)

	}
	for {
		doc, err := storage.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}

		var storageStruct _struct.Scaffolding
		bytes, err := json.Marshal(doc.Data())
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		err = json.Unmarshal(bytes, &storageStruct)
		if err != nil {
			log.Println("Her det fucker seg?")
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		storageArray = append(storageArray, storageStruct)
	}

	err := json.NewEncoder(w).Encode(storageArray)
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
	w.Header().Set("Content-Type", "application/json")
	query, err := tool.GetQueryProject(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	switch true {
	case !query.Has(constants.P_idURL) && !query.Has(constants.P_nameURL) && !query.Has(constants.P_State):
		getProjectCollection(w, r)
		break
	case query.Has(constants.P_idURL) || query.Has(constants.P_nameURL) || query.Has(constants.P_State):
		getProjectWithID(w, r)
		break
	default:
		invalidRequest(w, r)
		break
	}
}

/*
getProjectCollection will fetch every projects in the database.
*/
func getProjectCollection(w http.ResponseWriter, r *http.Request) {
	var projects []_struct.GetProject
	queryMap := getQuery(r)
	collectionIterator := projectCollection.Collections(database.Ctx)
	for {
		collRef, err := collectionIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			tool.HandleError(tool.COLLECTIONITERATORERROR, w)
			return
		}
		document := projectCollection.Collection(collRef.ID).Documents(database.Ctx)
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			var projectNew _struct.NewProject
			doc, err := database.GetDocumentData(documentRef.Ref)
			if err != nil {
				tool.HandleError(tool.NODOCUMENTWITHID, w)
				return
			}

			projectByte, err := json.Marshal(doc)
			err = json.Unmarshal(projectByte, &projectNew)
			if err != nil {
				tool.HandleError(tool.UNMARSHALLERROR, w)
				return
			}

			var project _struct.GetProject
			project.NewProject = projectNew
			if queryMap.Has(constants.P_scaffolding) && queryMap.Get(constants.P_scaffolding) == "true" {
				scaffold, err := getScaffoldingStruct(documentRef.Ref)
				if err != nil {
					tool.HandleError(tool.COULDNOTFINDDATA, w)
					return
				}
				project.ScaffoldingArray = scaffold

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
*/
func getProjectWithID(w http.ResponseWriter, r *http.Request) {
	queryMap, _ := tool.GetQueryProject(r)
	var documentReference []*firestore.DocumentRef
	var projects []_struct.GetProject
	var err error
	if queryMap.Has(constants.P_idURL) {
		intID, err := strconv.Atoi(queryMap.Get(constants.P_idURL))
		if err != nil {
			tool.HandleError(tool.INVALIDBODY, w)
			return
		}
		documentReference, err = iterateProjects(intID, "", "")

	} else if queryMap.Has(constants.P_nameURL) {
		documentReference, err = iterateProjects(0, queryMap.Get(constants.P_nameURL), "")
	} else {
		documentReference, err = iterateProjects(0, "", strings.Title(queryMap.Get(constants.P_State)))
	}
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	if documentReference == nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	for _, ref := range documentReference {

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

//invalidRequest
func invalidRequest(w http.ResponseWriter, r *http.Request) {
	tool.HandleError(tool.INVALIDREQUEST, w)
	return
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
			correctID, err = iterateProjects(num.ID, "", "")
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
	correctBody := checkProjectBody(bytes)
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
		correctID, err = iterateProjects(project.ProjectID, "", "")
	}

	if correctID != nil {
		tool.HandleError(tool.CouldNotAddSameID, w)
		return
	}

	state := project.State
	documentPath := projectCollection.Collection(state).Doc(id)

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
	err = addScaffolding(w, documentPath, batch)
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
	lastElement := getLastUrlElement(r)

	switch true {
	case constants.P_scaffolding == lastElement:
		transferProject(w, r)
	case constants.P_projectURL == lastElement:
		updateState(w, r)
	default:
		invalidRequest(w, r)
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

	correctBody := checkStateBody(data)
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

	documentReference, err := iterateProjects(stateStruct.ID, "", "")
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	project, err := database.GetDocumentData(documentReference[0])
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	newPath := projectCollection.Collection(stateStruct.State).Doc(strconv.Itoa(stateStruct.ID))
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
	_, inputScaffolding, err := getScaffoldingInput(w, r)
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
		fromLocation, fromPaths = getScaffoldingFromStorage(inputScaffolding.FromProjectID, inputScaffolding.InputScaffolding)
	default:
		fromLocation, fromPaths = getScaffoldingFromProject(inputScaffolding.FromProjectID, inputScaffolding.InputScaffolding)
	}
	if fromLocation == nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return
	}

	switch inputScaffolding.ToProjectID {
	case 0:
		newLocation, newPath = getScaffoldingFromStorage(inputScaffolding.ToProjectID, inputScaffolding.InputScaffolding)
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

		intVar, err := interfaceToInt(fromLocation[i])
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		if quantity > intVar {
			tool.HandleError(tool.CANNOTTRANSFERESCAFFOLDS, w)
			return
		}

		inputQuantity, err := interfaceToInt(inputScaffolding.InputScaffolding[i].Quantity)
		fromLoc, err := interfaceToInt(fromLocation[i])
		newLoc, err := interfaceToInt(newLocation[i])
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

	newPath, _ := iterateProjects(input, "", "")
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
func getScaffoldingFromStorage(input int, scaffold _struct.InputScaffolding) ([]interface{}, []*firestore.DocumentRef) {

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

func addScaffolding(w http.ResponseWriter, documentPath *firestore.DocumentRef, batch *firestore.WriteBatch) error {
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

		byte, err := json.Marshal(scaffoldingStruct)
		if err != nil {
			return err
		}

		var scaffoldingInput map[string]interface{}
		err = json.Unmarshal(byte, &scaffoldingInput)
		if err != nil {
			return err
		}

		batch.Set(scaffoldingTypeDocument, scaffoldingInput)
	}

	_, err := batch.Commit(database.Ctx)

	return err
}

func getScaffoldingStruct(document *firestore.DocumentRef) (_struct.ScaffoldingArray, error) {

	scaffolding := document.Collection(constants.P_StillasType).Documents(database.Ctx)
	var scaffoldArr _struct.ScaffoldingArray

	for {
		documentRef, err := scaffolding.Next()
		if err == iterator.Done {
			break
		}

		var scaffold _struct.Scaffolding
		doc, err := database.GetDocumentData(documentRef.Ref)
		if err != nil {
			return _struct.ScaffoldingArray{}, errors.New("no document")
		}

		projectByte, err := json.Marshal(doc)
		err = json.Unmarshal(projectByte, &scaffold)
		if err != nil {
			return _struct.ScaffoldingArray{}, errors.New("no document")
		}

		scaffoldArr = append(scaffoldArr, scaffold)

	}
	return scaffoldArr, nil
}
