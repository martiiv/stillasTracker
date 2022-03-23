package endpoints

import (
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

func deleteProject(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	var deleteID _struct.IDStruct
	json := json.Unmarshal(b, &deleteID)
	fmt.Println(json)

	for _, num := range deleteID {

		id := strconv.Itoa(num.ID)

		_, err := Database.Client.Collection("Location").Doc("Project").Collection("Active").Doc(id).Delete(Database.Ctx)
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}
		fmt.Println(num.ID)

	}

}

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
