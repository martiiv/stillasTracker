package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"net/http"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
	"strings"
	"time"
)

/**
Class profiles
This class will contain all data formatting and modification regarding the users of the system
Version 1.0
Last modified Martin Iversen 07.04.2022
TODO Maybe modularize som functionality the marshall, unmarshall encode routine is repeated often
*/

var baseCollection *firestore.DocumentRef

/*
ProfileRequest Function will redirect all requests from users to the appropriate function
Uses getProfile, createProfile, updateProfile and deleteProfile
*/
func ProfileRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Defines data type
	w.Header().Set("Access-Control-Allow-Origin", "*") //Allows mobile and web application to access the api
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	baseCollection = database.Client.Doc(constants.U_UsersCollection + "/" + constants.U_Employee) //Defines the path to the profiles collection in the database
	lastElement := tool.GetLastUrlElement(r)

	if lastElement != constants.U_User {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	requestType := r.Method //Defines the method of the request
	switch requestType {    //Forwards the request to the appropriate function
	case http.MethodGet:
		getProfile(w, r)
	case http.MethodPost:
		createProfile(w, r)
	case http.MethodPut:
		updateProfile(w, r)
	case http.MethodDelete:
		deleteProfile(w, r)
	}
}

/*
getProfile Function will get a user profile based on role or id
Function uses getUsersByRole and getIndividualUser as well as getAll
*/
func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query, err := tool.GetQueryProfile(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	switch true { //Forwards the request to the appropriate function based on the passed in query
	case query[constants.U_Role] != "":
		getUsersByRole(w, r)
	case query[constants.U_idURL] != "" || query[constants.U_nameURL] != "":
		getIndividualUser(w, r)
	default:
		getAll(w)
	}
}

/*
createProfile function adds new profiles to the database
Function uses iterateProfiles
*/
func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	requestBody, err := io.ReadAll(r.Body) //Reads body
	if !checkStruct(requestBody) {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var employee _struct.Employee                //Defines the appropriate struct
	err = json.Unmarshal(requestBody, &employee) //Unmarshall the requestBody into the struct
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	id := employee.EmployeeID                         //Converts the employee id to string
	_, err = iterateProfiles(employee.EmployeeID, "") //Iterates through the profiles using the id
	if err == nil {
		tool.HandleError(tool.CouldNotAddSameID, w)
		return
	}

	state := employee.Role                                   //Gets the state of the employee
	documentPath := baseCollection.Collection(state).Doc(id) //Defines the path to the profile using state and id
	var firebaseInput map[string]interface{}

	err = json.Unmarshal(requestBody, &firebaseInput) //Formats the requestBody
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	err = database.AddDocument(documentPath, firebaseInput) //Adds the profile to the database
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	} else {
		tool.HandleError(tool.ADDED, w)
	}
}

/*
updateProfile function updates the information about a user in the database
Function uses iterateProfiles and checkUpdate
*/
func updateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	batch := database.Client.Batch() //Defines the database operation

	var employeeStruct map[string]interface{}
	err = json.Unmarshal(data, &employeeStruct) //Unmarshall data into the employeeStruct
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	if !checkUpdate(employeeStruct) { //Checks if the body is in the correct format
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	employee := employeeStruct[constants.U_employeeID].(string) //Defines the employee id

	documentReference, err := iterateProfiles(employee, "") //Finds the profile using the id
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	var updates []firestore.Update

	for s, i := range employeeStruct { //Updates the information about the profiles
		update := firestore.Update{
			Path:  s,
			Value: i,
		}
		updates = append(updates, update)
	}

	for _, ref := range documentReference {
		batch.Update(ref, updates)
	}
	_, err = batch.Commit(database.Ctx) //Commits the database changes if all changes pass
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	}
}

/*
deleteProfile function deletes profiles
Function uses checkDeleteBody and iterateProfiles
*/
func deleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	batch := database.Client.Batch() //Defines the database operation

	requestBody, err := io.ReadAll(r.Body) //Read body of the request
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	ok := checkDeleteBody(requestBody) //Checks format
	if !ok {
		tool.HandleError(tool.INVALIDBODY, w)
		return
	}

	var deleteID _struct.ProfileDelete           //Defines the structure of the request
	err = json.Unmarshal(requestBody, &deleteID) //Unmarshall the body into the defined struct
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	for _, profile := range deleteID { //Iterates through the profiles
		document, err := iterateProfiles(profile.Id, "")
		if err != nil {
			tool.HandleError(tool.CouldNotDelete, w)
			return
		}

		batch.Delete(document[0]) //Deletes the profile
	}
	_, err = batch.Commit(database.Ctx) //Commits the change if all deletes pass
	if err != nil {
		tool.HandleError(tool.CouldNotDelete, w)
	}
}

/*
getAll function gets all the profiles in the database
*/
func getAll(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var employees []_struct.Employee //Defines the list of employees

	collection := baseCollection.Collections(database.Ctx) //Defines the path to the profiles collection
	for {                                                  //Iterates through the collection
		collRef, err := collection.Next()
		if err == iterator.Done || err != nil {
			break
		}

		document := baseCollection.Collection(collRef.ID).Documents(database.Ctx) //Defines the path to the documents within the collection
		for {                                                                     //Iterates through the documents
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			var employee _struct.Employee                       //Defines the employee
			doc, _ := database.GetDocumentData(documentRef.Ref) //Gets the employee
			employeeByte, err := json.Marshal(doc)              //Formatting
			err = json.Unmarshal(employeeByte, &employee)       //Unmarshall data into employee
			if err != nil {
				tool.HandleError(tool.UNMARSHALLERROR, w)
				return
			}
			employees = append(employees, employee)
		}
	}
	err := json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}

/*
getUsersByRole function gets profiles with specific roles
Function uses getQueryCustomer
*/
func getUsersByRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryValue := tool.GetQueryCustomer(w, r)

	documentPath := baseCollection.Collection(queryValue).Documents(database.Ctx)
	var employees []_struct.Employee
	for { //Iterates through the documents with the specified type
		documentRef, err := documentPath.Next()
		if err == iterator.Done {
			break
		}

		var employee _struct.Employee                       //Defines the employee struct
		doc, _ := database.GetDocumentData(documentRef.Ref) //Gets the profile from the database
		projectByte, err := json.Marshal(doc)               //Formats data
		err = json.Unmarshal(projectByte, &employee)        //Unmarshall data
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		employees = append(employees, employee)
	}
	if employees == nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	err := json.NewEncoder(w).Encode(employees)
	if err != nil {
		tool.HandleError(tool.NEWENCODERERROR, w)
		return
	}
}

/*
getIndividualUser function gets a user with a specific profile id
Function uses tool.GetQueryProfile, getUserByName and getIndividualUserByID
*/
func getIndividualUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query, err := tool.GetQueryProfile(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}
	switch true { //Forwards the request based on the passed in constant
	case query[constants.U_name] != "":
		getUserByName(w, r)
	case query[constants.U_idURL] != "":
		getIndividualUserByID(w, r)
	}
}

/*
getUserByName function gets profiles based on their name
Function uses iterateProfiles
*/
func getUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var documentReference []*firestore.DocumentRef
	var employees []_struct.Employee
	var err error

	queryMap := r.URL.Query()

	documentReference, err = iterateProfiles("", queryMap.Get(constants.U_nameURL)) //Gets the profile with the appropriate name
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	for _, ref := range documentReference { //Iterates through documentReference
		data, _ := database.GetDocumentData(ref) //Gets the profile
		jsonStr, err := json.Marshal(data)       //Formats the data
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		var employee _struct.Employee
		err = json.Unmarshal(jsonStr, &employee) //Unmarshalls the document into the struct
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}
		employees = append(employees, employee)
	}

	if employees == nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}

/*
getIndividualUserByID function gets a profile based on a user id
function uses iterateProfiles
*/
func getIndividualUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var documentReference []*firestore.DocumentRef
	var err error
	queryMap := mux.Vars(r)

	id := queryMap[constants.U_idURL] //Converts the query id to int

	documentReference, err = iterateProfiles(id, "") //Gets the documents with the id
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	data, _ := database.GetDocumentData(documentReference[0]) //Gets the profile from the database
	jsonStr, err := json.Marshal(data)                        //Formats the data
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return
	}

	var employee _struct.Employee            //Defines the struct
	err = json.Unmarshal(jsonStr, &employee) //Unmarshall data into the struct
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	if employee.EmployeeID == "" {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		return
	}
}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProfiles(id string, name string) ([]*firestore.DocumentRef, error) {
	var documentReferences []*firestore.DocumentRef

	collection := baseCollection.Collections(database.Ctx)
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
			document = baseCollection.Collection(collRef.ID).Where(constants.U_name+"."+constants.U_LastName, "==", name).Documents(database.Ctx)
		} else {
			document = baseCollection.Collection(collRef.ID).Where(constants.U_employeeID, "==", id).Documents(database.Ctx)
		}

		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			documentReferences = append(documentReferences, documentRef.Ref)

		}
	}

	if documentReferences != nil {
		return documentReferences, nil
	} else {
		return nil, errors.New("could not find document")
	}

}

/*
checkUpdate Function checks the update object
*/
func checkUpdate(update map[string]interface{}) bool {
	var counter int
	_, employeeID := update[constants.U_employeeID]
	_, name := update[constants.U_name]

	if name {
		var ok bool
		ok = checkName(update[constants.U_name])
		if !ok {
			return false
		}
	}

	_, admin := update[constants.U_admin].(bool)
	if !admin {
		return false
	}

	fields := []string{constants.U_name,
		constants.U_email, constants.U_admin, constants.U_Role, constants.U_employeeID, constants.U_phone}
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

/*
checkName Function checks the name object
*/
func checkName(name interface{}) bool {
	var counter int
	nameByte, err := json.Marshal(name)
	if err != nil {
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
	return true
}

/*
checkStruct Function checks the struct objects
*/
func checkStruct(body []byte) bool {
	var userMap map[string]interface{}
	err := json.Unmarshal(body, &userMap)
	if err != nil {
		return false
	}

	_, idFormat := userMap[constants.U_employeeID].(string)
	_, phoneFormat := userMap[constants.U_phone].(float64)
	roles := []string{constants.U_Admin, constants.U_Installer, constants.U_Storage}
	var roleFormat bool
	for _, role := range roles {
		if userMap[constants.U_Role] == strings.ToLower(role) {
			roleFormat = true
			break
		}
	}
	_, emailFormat := userMap[constants.U_email].(string)
	_, adminFormat := userMap[constants.U_admin].(bool)

	date, dateFormat := userMap[constants.U_dateOfBirth].(string)
	_, err = time.Parse("02-01-2006", date)
	if err != nil {
		return false
	}

	name := checkNameFormat(userMap[constants.U_name])
	validFormat := idFormat && phoneFormat && roleFormat && emailFormat && adminFormat && name && dateFormat

	if len(userMap) != 7 {
		return false
	}

	if !validFormat {
		return false
	}

	return true
}

/*
checkNameFormat Function checks the format of the name
*/
func checkNameFormat(name interface{}) bool {
	periodByte, err := json.Marshal(name)
	if err != nil {
		return false
	}

	var nameMap map[string]interface{}
	err = json.Unmarshal(periodByte, &nameMap)
	if err != nil {
		return false
	}

	_, firstName := nameMap[constants.U_FirstName].(string)
	_, lastName := nameMap[constants.U_LastName].(string)

	if !firstName || !lastName {
		return false
	}
	return true
}

/*
checkDeleteBody function checks formatting for the deleteProfile function
*/
func checkDeleteBody(bytes []byte) bool {
	var deleteID []map[string]interface{}
	err := json.Unmarshal(bytes, &deleteID) //Formats the body
	if err != nil {
		return false
	}

	for _, m := range deleteID { //Checks that it is in the appropriate format
		_, errDelete := m[constants.U_idURL].(string)
		if !errDelete {
			return false
		}
	}

	return true
}
