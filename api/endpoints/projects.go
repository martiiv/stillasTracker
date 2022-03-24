package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
Last modified Martin Iversen
*/

var start int = 3

func counter() int {
	start++
	return start - 1
}

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

	requestType := r.Method
	switch requestType {
	case "GET":
		storageRequest(w, r)
	case "POST":
		createProject()
	case "PUT":
		updateState(w, r)
	case "DELETE":
		deleteProject(w, r)

	}
}

func storageRequest(w http.ResponseWriter, r *http.Request) {
	id, err := CheckIDFromURL(r)
	if err != nil {
		collection := Database.Client.Collection("Location").Doc("Project").Collection("Active").Documents(Database.Ctx)
		data := Database.GetCollectionData(collection)

		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(jsonStr))

	} else {
		document, err := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(id).Get(Database.Ctx)
		if err != nil {
			fmt.Println(err)
		}
		data := document.Data()
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprint(w, string(jsonStr))

	}

}

/**
getProject will fetch the information from the selected project.
*/
func getProject(w http.ResponseWriter, r *http.Request) {

}

//deleteProject deletes selected projects from the database.
func deleteProject(w http.ResponseWriter, r *http.Request) {
	//TODO make this to a standalone function
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var deleteID _struct.IDStruct
	json := json.Unmarshal(bytes, &deleteID)
	fmt.Println(json)

	for _, num := range deleteID {

		id := strconv.Itoa(num.ID)

		_, err := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(id).Delete(Database.Ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}
		fmt.Println(num.ID)

	}

}

//createProject will create a Project and add it to the database
//TODO read struct from body
func createProject() {
	project := _struct.Project{
		ProjectID:   3,
		ProjectName: "GjovikRaadhus",
		Latitude:    60.79497726217587,
		Longitude:   10.692896676125931,
		Size:        430,
		State:       "Active",
		Period: _struct.Period{
			StartDate: "2020-05-09T22:00:00Z",
			EndDate:   "2020-02-19T23:00:00Z",
		},
		Address: _struct.Address{
			Street:       "Kauffeldts Plass 1",
			Zipcode:      "2815",
			Municipality: "Gjovik",
			County:       "Innlandet",
		},
		Customer: _struct.Customer{
			Name:   "Ola",
			Number: 932818193,
			Email:  "sjka@sosi.com",
		},
		Geofence: _struct.Geofence{},
	}

	id := strconv.Itoa(project.ProjectID)
	documentPath := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(id)

	var firebaseInput map[string]interface{}
	data, _ := json.Marshal(project)
	json.Unmarshal(data, &firebaseInput)

	fmt.Println(firebaseInput)

	Database.AddDocument(documentPath, firebaseInput)
}

//copyDocumentProject will get a document, and save it inside a struct.
func copyDocumentProject(documentPath *firestore.DocumentRef) _struct.Project {
	document, err := documentPath.Get(Database.Ctx)
	if err != nil {
		fmt.Println(err)
	}

	data := document.Data()
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	var project _struct.Project
	err = json.Unmarshal(jsonStr, &project)
	if err != nil {
		fmt.Println(err.Error())
	}

	return project
}

func updateState(w http.ResponseWriter, r *http.Request) {

	var stateStruct _struct.StateStruct
	json.NewDecoder(r.Body).Decode(&stateStruct)

	document := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(strconv.Itoa(stateStruct.ID))

	update := firestore.Update{
		Path:  "state",
		Value: stateStruct.State,
	}

	var updates []firestore.Update
	updates = append(updates, update)

	//Database.UpdateDocument(document, updates)

	batch := Database.Client.Batch()

	batch.Update(document, updates)

	batch.Commit(Database.Ctx)

}
