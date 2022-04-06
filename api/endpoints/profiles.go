package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
	"strconv"
	"strings"
	"time"
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

func ProfileRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	baseCollection = database.Client.Doc(constants.U_UsersCollection + "/" + constants.U_Employee)
	lastElement := tool.GetLastUrlElement(r)
	if lastElement != constants.U_User {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	requestType := r.Method
	switch requestType {
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

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	batch := database.Client.Batch()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	ok := checkDeleteBody(bytes)
	if !ok {
		tool.HandleError(tool.INVALIDBODY, w)
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
			tool.HandleError(tool.CouldNotDelete, w)
			return
		}

		batch.Delete(document[0])
	}
	_, err = batch.Commit(database.Ctx)
	if err != nil {
		tool.HandleError(tool.CouldNotDelete, w)
	}
}

func checkDeleteBody(bytes []byte) bool {

	var deleteID []map[string]interface{}
	err := json.Unmarshal(bytes, &deleteID)
	if err != nil {
		return false
	}

	for _, m := range deleteID {
		_, errDelete := m[constants.U_idURL].(float64)
		if !errDelete {
			return false
		}
	}

	return true
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}
	batch := database.Client.Batch()

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

	employee := employeeStruct[constants.U_employeeID].(float64)

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

	for _, ref := range documentReference {
		batch.Update(ref, updates)
	}
	_, err = batch.Commit(database.Ctx)
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

	_, err = iterateProfiles(employee.EmployeeID, "")
	if err == nil {
		tool.HandleError(tool.CouldNotAddSameID, w)
		return
	}

	state := employee.Role
	documentPath := baseCollection.Collection(state).Doc(id)

	var firebaseInput map[string]interface{}
	err = json.Unmarshal(bytes, &firebaseInput)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	err = database.AddDocument(documentPath, firebaseInput)
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	} else {
		tool.HandleError(tool.ADDED, w)
	}
}

//getProfile will fetch the profile based on employeeID or role.
func getProfile(w http.ResponseWriter, r *http.Request) {
	query, err := tool.GetQueryProfile(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}
	switch true {
	case query.Has(constants.U_Role):
		getUsersByRole(w, r)
	case query.Has(constants.U_idURL) || query.Has(constants.U_nameURL):
		getIndividualUser(w, r)
	default:
		getAll(w, r)
	}

}

func getAll(w http.ResponseWriter, r *http.Request) {
	var employees []_struct.Employee

	collection := baseCollection.Collections(database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		document := baseCollection.Collection(collRef.ID).Documents(database.Ctx)
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			var employee _struct.Employee
			doc, _ := database.GetDocumentData(documentRef.Ref)
			projectByte, err := json.Marshal(doc)
			err = json.Unmarshal(projectByte, &employee)
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

func getUsersByRole(w http.ResponseWriter, r *http.Request) {
	queryValue := getQueryCustomer(w, r)
	documentPath := baseCollection.Collection(queryValue).Documents(database.Ctx)
	var employees []_struct.Employee
	for {
		documentRef, err := documentPath.Next()
		if err == iterator.Done {
			break
		}

		var employee _struct.Employee
		doc, _ := database.GetDocumentData(documentRef.Ref)
		projectByte, err := json.Marshal(doc)
		err = json.Unmarshal(projectByte, &employee)
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

func getIndividualUser(w http.ResponseWriter, r *http.Request) {
	query, err := tool.GetQueryProfile(r)
	if !err {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}
	switch true {
	case query.Has(constants.U_name):
		getUserByName(w, r)
	case query.Has(constants.U_idURL):
		getIndividualUserByID(w, r)

	}
}

func getUserByName(w http.ResponseWriter, r *http.Request) {
	var documentReference []*firestore.DocumentRef
	var employees []_struct.Employee
	var err error
	queryMap := r.URL.Query()

	documentReference, err = iterateProfiles(0, queryMap.Get(constants.U_nameURL))
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	for _, ref := range documentReference {
		data, _ := database.GetDocumentData(ref)

		jsonStr, err := json.Marshal(data)
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
			return
		}

		var employee _struct.Employee
		err = json.Unmarshal(jsonStr, &employee)
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

func getIndividualUserByID(w http.ResponseWriter, r *http.Request) {

	var documentReference []*firestore.DocumentRef
	var err error
	queryMap := r.URL.Query()

	intID, err := strconv.Atoi(queryMap.Get(constants.U_idURL))
	if err != nil {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}
	documentReference, err = iterateProfiles(intID, "")
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	data, _ := database.GetDocumentData(documentReference[0])

	jsonStr, err := json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return
	}

	var employee _struct.Employee
	err = json.Unmarshal(jsonStr, &employee)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	if employee.EmployeeID == 0 {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		return
	}
}

func getQueryCustomer(w http.ResponseWriter, r *http.Request) string {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	_, ok := m[constants.U_Role]
	if ok {
		validRoles := []string{constants.U_Admin, constants.U_Storage, constants.U_Installer}
		for _, role := range validRoles {
			if m[constants.U_Role][0] == strings.ToLower(role) {
				return m[constants.U_Role][0]
			}
		}
	}
	http.Error(w, "no valid query", http.StatusBadRequest)
	return ""

}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProfiles(id int, name string) ([]*firestore.DocumentRef, error) {
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
		constants.U_email, constants.U_admin, constants.U_employeeID, constants.U_phone}
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

func checkStruct(body []byte) bool {
	var userMap map[string]interface{}
	err := json.Unmarshal(body, &userMap)
	if err != nil {
		return false
	}

	_, idFormat := userMap[constants.U_employeeID].(float64)
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
