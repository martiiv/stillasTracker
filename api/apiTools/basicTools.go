package apiTools

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"stillasTracker/api/constants"
	"strings"
)

func CreatePath(segments []string) string {
	var finalPath string
	for _, s := range segments {
		finalPath += s + "/"
	}
	finalPath = strings.TrimRight(finalPath, "/")
	return finalPath
}

//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
func GetQueryProject(r *http.Request) (url.Values, bool) {
	query := r.URL.Query()

	allowedQuery := map[string]bool{constants.P_idURL: true, constants.P_nameURL: true, constants.P_scaffolding: true, constants.P_State: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	if query.Has(constants.P_scaffolding) {
		if !(query.Get(constants.P_scaffolding) == "true" || query.Get(constants.P_scaffolding) == "false") {
			return nil, false
		}
	}

	return query, true
}

//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
func GetQueryProfile(r *http.Request) (url.Values, bool) {
	query := r.URL.Query()

	//Defines the allowed parts of the url
	allowedQuery := map[string]bool{constants.U_nameURL: true, constants.U_Role: true, constants.U_idURL: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	//Checks that the URL only contains the allowed roles
	if query.Has(constants.U_Role) {
		if !(query.Get(constants.U_Role) == constants.U_admin || query.Get(constants.U_Role) == strings.ToLower(constants.U_Installer) || query.Get(constants.U_Role) == strings.ToLower(constants.U_Storage)) {
			return nil, false
		}
	}

	return query, true
}

func GetQueryScaffolding(r *http.Request) (url.Values, bool) {
	query := r.URL.Query()

	allowedQuery := map[string]bool{constants.S_id: true, constants.S_type: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	if query.Has(constants.S_type) {
		for _, i := range constants.ScaffoldingTypes {
			if i == query.Get(constants.S_type) {
				return nil, true
			}
		}
	}

	return query, true
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

func GetQueries(w http.ResponseWriter, r *http.Request) url.Values {
	w.Header().Set("Content-Type", "application/json")
	lastElement := GetLastUrlElement(r)

	switch true {
	case "unit" == lastElement:
		query, err := GetQueryScaffolding(r)
		if !err {
			HandleError(INVALIDREQUEST, w)
		}
		return query

	case "project" == lastElement:
		query, err := GetQueryProject(r)
		if !err {
			HandleError(INVALIDREQUEST, w)
		}
		return query

	case "profile" == lastElement:
		query, err := GetQueryProfile(r)
		if !err {
			HandleError(INVALIDREQUEST, w)
		}
		return query
	}

	return r.URL.Query()
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
