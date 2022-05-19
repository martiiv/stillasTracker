package apiTools

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"stillasTracker/api/constants"
	"strings"
)

/*
Class basicTools contains functions for assisting API endpoints
Last update 19.05.2022 Martin Ivesren
@version 1.0
*/

func CreatePath(segments []string) string {
	var finalPath string
	for _, s := range segments {
		finalPath += s + "/"
	}
	finalPath = strings.TrimRight(finalPath, "/")
	return finalPath
}

/*
GetQueryScaffolding function checks that the queries are valid in the scaffolding requests
Code inspired by the following stackoverflow issue:
//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
*/
func GetQueryScaffolding(r *http.Request) (map[string]string, bool) {
	query := mux.Vars(r)

	allowedQuery := map[string]bool{constants.S_id: true, constants.S_type: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}
	valid := true
	if query[constants.S_type] != "" {
		for i := range constants.ScaffoldingTypes {
			if !(query[constants.S_type] == constants.ScaffoldingTypes[i]) {
				valid = false
			} else {
				valid = true
				break
			}
		}
		if valid == false {
			return nil, valid
		}
	}
	return query, true
}

/*
GetQueryProject function checks that the queries are valid in the project requests
Code inspired by the following stackoverflow issue:
//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
*/
func GetQueryProject(r *http.Request) (map[string]string, bool) {
	query := mux.Vars(r)

	allowedQuery := map[string]bool{constants.P_idURL: true, constants.P_nameURL: true, constants.P_scaffolding: true, constants.P_State: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	if query[constants.P_scaffolding] != "" {
		if !(query[constants.P_scaffolding] == "true" || query[constants.P_scaffolding] == "false") {
			return nil, false
		}
	}
	return query, true
}

/*
GetQueryProfile function checks that the queries are valid in the profile requests
Code inspired by the following stackoverflow issue:
//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
*/
func GetQueryProfile(r *http.Request) (map[string]string, bool) {
	query := mux.Vars(r)

	//Defines the allowed parts of the url
	allowedQuery := map[string]bool{constants.U_nameURL: true, constants.U_Role: true, constants.U_idURL: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	//Checks that the URL only contains the allowed roles
	if query[constants.U_Role] != "" {
		if !(query[constants.U_Role] == constants.U_admin || query[constants.U_Role] == strings.ToLower(constants.U_Installer) || query[constants.U_Role] == strings.ToLower(constants.U_Storage)) {
			return nil, false
		}
	}

	return query, true
}

/*
GetQueryCustomer Function returns a query list containing the queries specific to the profile endpoint
*/
func GetQueryCustomer(w http.ResponseWriter, r *http.Request) string {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queries := mux.Vars(r)

	validRoles := []string{constants.U_Admin, constants.U_Storage, constants.U_Installer}

	for _, role := range validRoles {
		if queries[constants.U_Role] == strings.ToLower(role) {
			return queries[constants.U_Role]
		}
	}
	return "Invalid query"

}

/*
GetLastUrlElement will split the url and return the last element.
*/
func GetLastUrlElement(r *http.Request) string {
	url := r.URL.Path
	trimmedURL := strings.TrimRight(url, "/")
	splittedURL := strings.Split(trimmedURL, "/")
	lastElement := splittedURL[len(splittedURL)-1]
	return lastElement
}

func InterfaceToInt(input interface{}) (int, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return 0, errors.New("cannot marshal")
	}

	var returnInt int
	err = json.Unmarshal(bytes, &returnInt)
	if err != err {
		return 0, errors.New("cannot unmarshal")
	}

	return returnInt, nil
}

//https://freshman.tech/snippets/go/check-if-slice-contains-element/
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//InvalidRequest
func InvalidRequest(w http.ResponseWriter, r *http.Request) {
	HandleError(INVALIDREQUEST, w)
	return
}

func StructToMap(input interface{}) ([]map[string]interface{}, error) {
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
