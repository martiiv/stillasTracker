package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
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
	err = json.Unmarshal(bytes, &deleteID)
	if err != nil {
		return
	}

	for _, num := range deleteID {
		_, err := iterateProfiles(num.Id).Delete(Database.Ctx)
		if err != nil {
			return
		}

		fmt.Println(num.Id)
	}
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	batch := Database.Client.Batch()

	var employeeStruct map[string]interface{}
	err = json.Unmarshal(data, &employeeStruct)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	if !checkUpdate(employeeStruct) {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	employee := employeeStruct["employeeID"].(float64)

	documentReference := iterateProfiles(int(employee))

	var updates []firestore.Update

	for s, i := range employeeStruct {
		update := firestore.Update{
			Path:  s,
			Value: i,
		}
		updates = append(updates, update)
	}

	batch.Update(documentReference, updates)
	_, err = batch.Commit(Database.Ctx)
	if err != nil {
		http.Error(w, "could not save changes", http.StatusConflict)
		return
	}

}

func createProfile(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if !checkStruct(bytes) {
		http.Error(w, "body invalid", http.StatusBadRequest)
		return
	}

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

func checkUpdate(update map[string]interface{}) bool {
	fmt.Println(update)
	var counter int
	_, employeeID := update["employeeID"]
	_, name := update["name"]

	if name {
		var ok bool
		ok = checkName(update["name"])
		if !ok {
			return false
		}
	}

	fields := []string{"name", "role", "phone", "email", "admin", "employeeID"}
	if employeeID {
		for _, field := range fields {
			for f := range update {
				if field == f {
					counter++
					break
				}
			}
		}
		if len(update) > (counter) {
			return false
		}

	}

	return true
}

func checkName(name interface{}) bool {
	var counter int
	nameByte, err := json.Marshal(name)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var nameMap map[string]interface{}
	err = json.Unmarshal(nameByte, &nameMap)
	if err != nil {
		return false
	}

	validKeys := []string{"firstName", "lastName"}
	for s := range nameMap {
		for _, key := range validKeys {
			if s == key {
				counter++
				break
			}
		}
	}

	if len(nameMap) > counter {
		return false
	}
	fmt.Println(name)
	return true
}

func checkStruct(body []byte) bool {
	var userMap map[string]interface{}
	err := json.Unmarshal(body, &userMap)
	if err != nil {
		return false
	}

	_, idFormat := userMap["employeeID"].(float64)
	_, phoneFormat := userMap["phone"].(float64)
	roles := []string{"admin", "installer", "storage"}
	var roleFormat bool
	for _, role := range roles {
		if userMap["role"] == role {
			roleFormat = true
			break
		}
	}
	_, emailFormat := userMap["email"].(string)
	_, adminFormat := userMap["admin"].(bool)
	_, dateFormat := userMap["dateOfBirth"].(string)
	name := checkNameFormat(userMap["name"])

	validFormat := idFormat && phoneFormat && roleFormat && emailFormat && adminFormat && name && dateFormat

	if !validFormat {
		return false
	}

	return true
}

func checkNameFormat(name interface{}) bool {
	periodByte, err := json.Marshal(name)
	if err != nil {
		fmt.Println(err.Error())
	}

	var nameMap map[string]interface{}
	err = json.Unmarshal(periodByte, &nameMap)
	if err != nil {
		return false
	}

	_, firstName := nameMap["firstName"].(string)
	_, lastName := nameMap["lastName"].(string)

	if !firstName || !lastName {
		return false
	}
	return true
}
