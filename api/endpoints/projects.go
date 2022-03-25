package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
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

/*const project = _struct.Project{
		ProjectID:   4,
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
*/

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
		getProject(w, r)
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
	w.Header().Set("Content-Type", "application/json")
	id, err := CheckIDFromURL(r)
	if err != nil {

		//collection := Database.Client.Collection("Location").Doc("Project").Collection("Active").Documents(Database.Ctx)

		var projects []_struct.Project

		collection := Database.Client.Collection("Location").Doc("Project").Collections(Database.Ctx)
		for {
			collRef, err := collection.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				break
			}
			document := Database.Client.Collection("Location").Doc("Project").Collection(collRef.ID).Documents(Database.Ctx)
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

	/*id := strconv.Itoa(project.ProjectID)
	documentPath := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(id)

	var firebaseInput map[string]interface{}
	data, _ := json.Marshal(project)
	json.Unmarshal(data, &firebaseInput)

	fmt.Println(firebaseInput)

	Database.AddDocument(documentPath, firebaseInput)*/
}

//updateState will change the state of the project. In an atomic operation the project will change state,
//be moved into the state collection and deleted form the old state collection.
func updateState(w http.ResponseWriter, r *http.Request) {
	batch := Database.Client.Batch()

	var stateStruct _struct.StateStruct
	err := json.NewDecoder(r.Body).Decode(&stateStruct)
	if err != nil {
		return
	}

	documentReference := iterateProjects(stateStruct.ID)

	project, err := Database.GetDocumentData(documentReference)
	if err != nil {
		fmt.Println(err.Error())
	}

	newPath := Database.Client.Collection("Location").Doc("Project").Collection(stateStruct.State).Doc(strconv.Itoa(stateStruct.ID))
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

func iterateProjects(id int) *firestore.DocumentRef {
	var documentReference *firestore.DocumentRef
	collection := Database.Client.Collection("Location").Doc("Project").Collections(Database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		document := Database.Client.Collection("Location").Doc("Project").Collection(collRef.ID).Documents(Database.Ctx)
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

func checkStateBody(body io.ReadCloser) {
	file := json.NewDecoder(body)
	err := file.Decode(&_struct.StateStruct{})
	if err != nil {
		return
	}
	fmt.Println(file)
}
