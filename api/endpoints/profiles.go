package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"net/http"
	"net/url"
	"stillasTracker/api/Database"
	_struct "stillasTracker/api/struct"
	"strconv"
)

/**
Class profiles
This class will contain all data formating and modification regarding the users of the system
Class contains the following functions:
	- getProfiles: The function returns all the active profiles in the system
	- updateProfile: The function lets a user modify their profile
	- createProfile: THe function lets the admin create a new profile
	- deleteProfile: The function deletes a user profile

Version 0.1
Last modified Aleksander Aaboen
*/

var baseCollection *firestore.DocumentRef

func profileRequest(w http.ResponseWriter, r *http.Request) {
	baseCollection = Database.Client.Doc("Users/Employee")

	requestType := r.Method
	switch requestType {
	case "GET":
		getProfile(w, r)
	case "POST":
		createProfile(w, r)
	case "PUT":
		updateProfile(w, r)
	case "DELETE":
		deleteProfile(w, r)
	}
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var deleteID _struct.ProfileDelete
	json.Unmarshal(bytes, &deleteID)

	for _, num := range deleteID {
		_, err := iterateProfiles(num.Id).Delete(Database.Ctx)
		if err != nil {
			return
		}

		fmt.Println(num.Id)
	}
}

func updateProfile(w http.ResponseWriter, r *http.Request) {

}

func createProfile(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)

	var employee _struct.Employee

	err = json.Unmarshal(bytes, &employee)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	id := strconv.Itoa(employee.EmployeeID)
	state := employee.Role
	documentPath := baseCollection.Collection(state).Doc(id)

	var firebaseInput map[string]interface{}
	json.Unmarshal(bytes, &firebaseInput)

	Database.AddDocument(documentPath, firebaseInput)
}

//getProfile will fetch the profile based on employeeID or role.
func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := CheckIDFromURL(r)
	var documentPath *firestore.DocumentIterator
	var employees []_struct.Employee

	if r.URL.Query().Has("role") {
		queryValue := getQuery(w, r)
		documentPath = baseCollection.Collection(queryValue).Documents(Database.Ctx)
		for {
			documentRef, err := documentPath.Next()
			if err == iterator.Done {
				break
			}

			var employee _struct.Employee
			doc, _ := Database.GetDocumentData(documentRef.Ref)
			projectByte, err := json.Marshal(doc)
			err = json.Unmarshal(projectByte, &employee)
			if err != nil {
				fmt.Println(err.Error())
			}

			employees = append(employees, employee)
		}
		err := json.NewEncoder(w).Encode(employees)
		if err != nil {
			http.Error(w, "error when decoding", http.StatusInternalServerError)
			return
		}
	} else if id != "" {

		intID, err := strconv.Atoi(id)

		documentReference := iterateProfiles(intID)
		data, _ := Database.GetDocumentData(documentReference)

		if err != nil {
			fmt.Println(err)
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}

		var employee _struct.Employee
		err = json.Unmarshal(jsonStr, &employee)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.NewEncoder(w).Encode(employee)
		if err != nil {
			return
		}
	} else {

		collection := baseCollection.Collections(Database.Ctx)
		for {
			collRef, err := collection.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				break
			}
			document := baseCollection.Collection(collRef.ID).Documents(Database.Ctx)
			for {
				documentRef, err := document.Next()
				if err == iterator.Done {
					break
				}

				var employee _struct.Employee
				doc, _ := Database.GetDocumentData(documentRef.Ref)
				projectByte, err := json.Marshal(doc)
				err = json.Unmarshal(projectByte, &employee)
				if err != nil {
					fmt.Println(err.Error())
				}

				employees = append(employees, employee)
			}
		}

		err := json.NewEncoder(w).Encode(employees)
		if err != nil {
			return
		}
	}

}

func getQuery(w http.ResponseWriter, r *http.Request) string {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, ok := m["role"]
	if ok {
		validRoles := []string{"Admin", "Storage", "installer"}
		for _, role := range validRoles {
			if m["role"][0] == role {
				return m["role"][0]
			}
		}
	}
	http.Error(w, "no valid query", http.StatusBadRequest)
	return ""

}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProfiles(id int) *firestore.DocumentRef {
	var documentReference *firestore.DocumentRef
	collection := baseCollection.Collections(Database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		document := baseCollection.Collection(collRef.ID).Documents(Database.Ctx)
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
