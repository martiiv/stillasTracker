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
	"net/url"
	"stillasTracker/api/Database"
	tool "stillasTracker/api/apiTools"
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
	case http.MethodGet:
		getIndividualUser(w, r)
	case http.MethodPost:
		createProfile(w, r)
	case http.MethodPut:
		updateProfile(w, r)
	case http.MethodDelete:
		deleteProfile(w, r)
	}
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	var deleteID _struct.ProfileDelete
	err = json.Unmarshal(bytes, &deleteID)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	for _, num := range deleteID {
		document, err := iterateProfiles(num.Id, "")
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}
		_, err = document.Delete(Database.Ctx)
		if err != nil {
			tool.HandleError(tool.DELETE, w)
			return
		}
	}
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}
	batch := Database.Client.Batch()

	var employeeStruct map[string]interface{}
	err = json.Unmarshal(data, &employeeStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	if !checkUpdate(employeeStruct) {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	employee := employeeStruct["employeeID"].(float64)

	documentReference, err := iterateProfiles(int(employee), "")
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

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
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	}

}

func createProfile(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if !checkStruct(bytes) {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var employee _struct.Employee

	err = json.Unmarshal(bytes, &employee)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	id := strconv.Itoa(employee.EmployeeID)
	state := employee.Role
	documentPath := baseCollection.Collection(state).Doc(id)

	var firebaseInput map[string]interface{}
	err = json.Unmarshal(bytes, &firebaseInput)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	//Todo sjekk om id ikke er tatt
	err = Database.AddDocument(documentPath, firebaseInput)
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	}
}

//getProfile will fetch the profile based on employeeID or role.
func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := getLastUrlElement(r)
	var documentPath *firestore.DocumentIterator
	var employees []_struct.Employee

	if r.URL.Query().Has("role") {
		queryValue := getQueryCustomer(w, r)
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
		getIndividualUser(w, r)
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

func getIndividualUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var documentReference *firestore.DocumentRef
	var err error
	queryMap := getQuery(r)

	if queryMap.Has("id") {
		intID, err := strconv.Atoi(queryMap.Get("id"))
		if err != nil {
			tool.HandleError(tool.INVALIDREQUEST, w)
			return
		}
		documentReference, err = iterateProfiles(intID, "")
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}
	} else if queryMap.Has("name") {
		documentReference, err = iterateProfiles(0, queryMap.Get("name"))
		if err != nil {
			tool.HandleError(tool.COULDNOTFINDDATA, w)
			return
		}
	} else {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	data, _ := Database.GetDocumentData(documentReference)

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
}

func getQueryCustomer(w http.ResponseWriter, r *http.Request) string {
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
func iterateProfiles(id int, name string) (*firestore.DocumentRef, error) {
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
		var document *firestore.DocumentIterator
		if name != "" {
			document = baseCollection.Collection(collRef.ID).Where("name.lastName", "==", name).Documents(Database.Ctx)
		} else {
			document = baseCollection.Collection(collRef.ID).Where("employeeID", "==", id).Documents(Database.Ctx)
		}

		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			documentReference = documentRef.Ref

		}
	}

	if documentReference != nil {
		return documentReference, nil
	} else {
		return nil, errors.New("could not find document")
	}

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

	} else {
		return false
	}

	return employeeID
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
	//Todo sjekk om dato er skrevet på riktig format
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
