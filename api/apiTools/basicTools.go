package apiTools

import (
	"io/ioutil"
	"net/http"
)

/**
Class basicTools
This class will contain all basic REST API functions which will be used throughout the API
This class contains the following functions:
	-

Version 0.1
Last modified Martin Iversen
*/

//genericRequest
/**
Generic request function
Takes in a writer and reader in addition to the API request url
The function handles the request and returns the body of the request
*/
func genericRequest(w http.ResponseWriter, r *http.Request, request string) []byte {
	w.Header().Set("Content-Type", "application/json")
	client := &http.Client{}

	apiRequest, err := http.NewRequest(http.MethodGet, request, nil)
	if err != nil {
		getErrorMessage(w, err)
	}

	response, err := client.Do(apiRequest)
	if err != nil {
		getErrorMessage(w, err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		getErrorMessage(w, err)
	}

	return body
}
