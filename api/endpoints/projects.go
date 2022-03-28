package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"stillasTracker/api/Database"
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

Version 0.1
Last modified Aleksander Aaboen
*/
var projectCollection *firestore.DocumentRef

func CheckIDFromURL(r *http.Request) (string, error) {
	url := strings.Split(r.RequestURI, "/")
	lastUrlSegment := url[len(url)-1]
	matched, _ := regexp.MatchString(`\d`, lastUrlSegment)
	if matched {
		return lastUrlSegment, nil
	}
	return "", errors.New("not a valid ID")
}

/**
Main function to switch between the different request types.
*/
func projectRequest(w http.ResponseWriter, r *http.Request) {

	projectCollection = Database.Client.Doc("Location/Project")

	requestType := r.Method
	switch requestType {
	case "GET":
		getProject(w, r)
	case "POST":
		createProject(w, r)
	case "PUT":
		transactionAdd(w, r)
	case "DELETE":
		deleteProject(w, r)

	}
}

//storageRequest will return all the scaffolding parts in the selected storage location.
func storageRequest(w http.ResponseWriter, r *http.Request) {

}

/**
getProject will fetch the information from the selected project.
*/
func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := CheckIDFromURL(r)
	if err != nil {
		var projects []_struct.Project

		collection := projectCollection.Collections(Database.Ctx)
		for {
			collRef, err := collection.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				break
			}
			document := projectCollection.Collection(collRef.ID).Documents(Database.Ctx)
			for {
				documentRef, err := document.Next()
				if err == iterator.Done {
					break
				}

				var project _struct.Project
				doc, _ := Database.GetDocumentData(documentRef.Ref)
				projectByte, err := json.Marshal(doc)
				err = json.Unmarshal(projectByte, &project)
				if err != nil {
					fmt.Println(err.Error())
				}

				projects = append(projects, project)
			}
		}

		err := json.NewEncoder(w).Encode(projects)
		if err != nil {
			return
		}

	} else {

		intID, err := strconv.Atoi(id)

		documentReference := iterateProjects(intID)
		data, _ := Database.GetDocumentData(documentReference)

		if err != nil {
			fmt.Println(err)
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		var project _struct.Project
		err = json.Unmarshal(jsonStr, &project)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.NewEncoder(w).Encode(project)
		if err != nil {
			return
		}
	}
}

//deleteProject deletes selected projects from the database.
func deleteProject(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var deleteID _struct.IDStruct
	json := json.Unmarshal(bytes, &deleteID)
	fmt.Println(json)

	for _, num := range deleteID {

		id := strconv.Itoa(num.ID)
		_, err := projectCollection.Collection("Active").Doc(id).Delete(Database.Ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}
		fmt.Println(num.ID)

	}

}

//createProject will create a Project and add it to the database
func createProject(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	correctBody := checkProjectBody(bytes)
	if !correctBody {
		http.Error(w, "Body is not correct", http.StatusBadRequest)
		return
	}

	var project _struct.NewProject

	err = json.Unmarshal(bytes, &project)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	id := strconv.Itoa(project.ProjectID)
	state := project.State
	documentPath := projectCollection.Collection(state).Doc(id)

	var firebaseInput map[string]interface{}
	bytes, err = json.Marshal(project)
	json.Unmarshal(bytes, &firebaseInput)

	Database.AddDocument(documentPath, firebaseInput)
}

/*updateState will change the state of the project. In an atomic operation the project will change state,
be moved into the state collection and deleted form the old state collection.*/
func updateState(w http.ResponseWriter, r *http.Request) {
	batch := Database.Client.Batch()
	data, err := ioutil.ReadAll(r.Body)

	correctBody := checkStateBody(data)
	if !correctBody {
		http.Error(w, "Body is not correct", http.StatusBadRequest)
		return
	}

	var stateStruct _struct.StateStruct
	err = json.Unmarshal(data, &stateStruct)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	documentReference := iterateProjects(stateStruct.ID)

	project, err := Database.GetDocumentData(documentReference)
	if err != nil {
		fmt.Println(err.Error())
	}

	newPath := projectCollection.Collection(stateStruct.State).Doc(strconv.Itoa(stateStruct.ID))
	batch.Create(newPath, project)

	batch.Delete(documentReference)
	update := firestore.Update{
		Path:  "state",
		Value: stateStruct.State,
	}
	var updates []firestore.Update
	updates = append(updates, update)

	batch.Update(newPath, updates)

	batch.Commit(Database.Ctx)

}

func getScaffoldingInput(w http.ResponseWriter, r *http.Request) ([]_struct.Scaffolding, _struct.InputScaffoldingWithID) {

	data, err := ioutil.ReadAll(r.Body)

	var inputScaffolding _struct.InputScaffoldingWithID
	err = json.Unmarshal(data, &inputScaffolding)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return nil, _struct.InputScaffoldingWithID{}
	}
	var scaffolds []_struct.Scaffolding
	for i := range inputScaffolding.InputScaffolding {

		quantity := _struct.Quantity{
			Expected:   inputScaffolding.InputScaffolding[i].Quantity,
			Registered: 0,
		}

		scaffolding := _struct.Scaffolding{
			Category: inputScaffolding.InputScaffolding[i].Type,
			Quantity: quantity,
		}

		scaffolds = append(scaffolds, scaffolding)
	}
	return scaffolds, inputScaffolding
}

func transactionAdd(w http.ResponseWriter, r *http.Request) {

	scaffolds, inputScaffolding := getScaffoldingInput(w, r)

	var fromPath *firestore.DocumentRef

	switch inputScaffolding.FromProjectID {
	case 0:
		fromPath = Database.Client.Doc("Location/Storage/Inventory/Spire")
	default:
		fromPath = iterateProjects(inputScaffolding.FromProjectID)
	}

	newPath := iterateProjects(inputScaffolding.ToProjectID)
	err := Database.Client.RunTransaction(Database.Ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(fromPath)
		if err != nil {
			return err
		}

		newPath.Path = newPath.Path + "/StillasType/Spire"

		to, err := tx.Get(newPath)
		if err != nil {
			return err
		}

		scaffoldingFrom, err := doc.DataAt("Quantity.expected")
		if err != nil {
			return err
		}

		scaffoldingTo, err := to.DataAt("Quantity.expected")
		if err != nil {
			return err
		}

		//numb:= scaffoldingFrom.(int64)

		if scaffolds[0].Quantity.Expected > 100000 {
			return nil
		}

		var sub = map[string]interface{}{}
		sub["Quantity"] = map[string]interface{}{}
		sub["Quantity"].(map[string]interface{})["expected"] = scaffoldingFrom.(int64) - int64(scaffolds[0].Expected)

		err = tx.Set(fromPath, sub, firestore.MergeAll)

		var add = map[string]interface{}{}
		add["Quantity"] = map[string]interface{}{}
		add["Quantity"].(map[string]interface{})["expected"] = scaffoldingTo.(int64) + int64(scaffolds[0].Expected)

		err = tx.Set(newPath, add, firestore.MergeAll)

		return err
	})
	if err != nil {
		// Handle any errors appropriately in this section.
		log.Printf("An error has occurred: %s", err)
	}
}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProjects(id int) *firestore.DocumentRef {
	var documentReference *firestore.DocumentRef
	collection := projectCollection.Collections(Database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		document := projectCollection.Collection(collRef.ID).Documents(Database.Ctx)
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			if documentRef.Ref.ID == strconv.Itoa(id) {
				fmt.Printf("Found ID in  collection: %s\n", collRef.ID)
				documentReference = documentRef.Ref
				break
			}
		}
	}
	return documentReference
}

//checkStateBody will check if the body is of correct format, and if the values are correct datatypes.
func checkStateBody(body []byte) bool {
	var dat map[string]interface{}
	err := json.Unmarshal(body, &dat)
	if err != nil {
		return false
	}
	_, stateBool := dat["state"]
	_, idBool := dat["id"]
	_, isFloat := dat["id"].(float64)

	var correctValues bool
	if stateBool && idBool && isFloat {
		correctValues = checkState(dat["state"].(string))
	}
	return correctValues
}

//checkProjectBody function that will verify the correct format of project struct
func checkProjectBody(body []byte) bool {
	var project map[string]interface{}
	err := json.Unmarshal(body, &project)
	if err != nil {
		return false
	}

	period := project["period"]
	correctPeriod := checkPeriod(period)
	costumer := project["customer"]
	correctCustomer := checkCustomer(costumer)
	geoFence := project["geofence"]
	correctGeofence := checkGeofence(geoFence)
	_, longitudeFloat := project["longitude"].(float64)
	_, latitudeFloat := project["latitude"].(float64)
	_, sizeFloat := project["size"].(float64)
	_, projectID := project["projectID"].(float64)
	_, projectName := project["projectName"].(string)

	validState := checkState(project["state"].(string))
	correctFormat := validState && longitudeFloat && latitudeFloat && sizeFloat &&
		projectID && correctGeofence && correctCustomer && correctPeriod && projectName

	return correctFormat

}

//checkPeriod function that will verify the correct format of period struct
func checkPeriod(period interface{}) bool {
	periodByte, err := json.Marshal(period)
	if err != nil {
		fmt.Println(err.Error())
	}
	var periods map[string]interface{}
	err = json.Unmarshal(periodByte, &periods)
	if err != nil {
		return false
	}

	nestedPeriod := []string{"startDate", "endDate"}
	for _, period := range nestedPeriod {
		_, ok := periods[period]
		if !ok {
			return false
		}
	}

	return true
}

//checkCustomer function that will verify the correct format of customer struct
func checkCustomer(customer interface{}) bool {
	periodByte, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err.Error())
	}
	var customerMap map[string]interface{}
	err = json.Unmarshal(periodByte, &customerMap)
	if err != nil {
		return false
	}
	nestedPeriod := []string{"name", "email", "number"}
	for _, period := range nestedPeriod {
		_, ok := customerMap[period]
		if !ok {
			return false
		}
	}

	_, numberFloat := customerMap["number"].(float64)
	if !numberFloat {
		return false
	}

	return true
}

//checkGeofence function that will verify the correct format of geofence struct
func checkGeofence(fence interface{}) bool {
	periodByte, err := json.Marshal(fence)
	if err != nil {
		fmt.Println(err.Error())
	}
	var geofenceMap map[string]interface{}
	err = json.Unmarshal(periodByte, &geofenceMap)
	if err != nil {
		return false
	}

	nestedPeriod := []string{"w-position", "x-position", "y-position", "z-position"}
	for _, period := range nestedPeriod {
		_, ok := geofenceMap[period]
		if !ok {
			return false
		} else {
			correctCoordinate := checkGeofenceCoordinates(geofenceMap[period])
			if !correctCoordinate {
				return false
			}
		}
	}

	return true
}

//checkGeofenceCoordinates function that will verify the correct format of geofence position struct
func checkGeofenceCoordinates(location interface{}) bool {
	periodByte, err := json.Marshal(location)
	if err != nil {
		fmt.Println(err.Error())
	}
	var locationMap map[string]interface{}
	err = json.Unmarshal(periodByte, &locationMap)
	if err != nil {
		return false
	}

	coordinates := []string{"longitude", "latitude"}
	for _, period := range coordinates {
		_, ok := locationMap[period]
		if !ok {
			return false
		}
	}

	_, longitudeFloat := locationMap["longitude"].(float64)
	_, latitudeFloat := locationMap["latitude"].(float64)

	if !latitudeFloat || !longitudeFloat {
		return false
	}

	return true
}

//checkState will check the value of the body, to ensure that the user has selected the correct state.
func checkState(input string) bool {
	state := []string{"Active", "Inactive", "Upcoming"}

	var correctValues bool
	for _, states := range state {
		if input == states {
			correctValues = true
			break
		}
	}
	return correctValues
}
