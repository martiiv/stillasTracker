package endpoints

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

func genericRequest(w http.ResponseWriter, r *http.Request, request string) []byte {
	w.Header().Set("Content-Type", "application/json")
	client := &http.Client{}

	apiRequest, err := http.NewRequest(http.MethodGet, request, nil)
	if err != nil {
		getErrorMessage(err)
	}

	response, err := client.Do(apiRequest)
	if err != nil {
		getErrorMessage(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		getErrorMessage(err)
	}

	return body
}

func getErrorMessage() {

}
