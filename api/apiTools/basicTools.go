package apiTools

import (
	"net/http"
	"net/url"
)

/**
Class basicTools
This class will contain all basic REST API functions which will be used throughout the API
This class contains the following functions:
	-

Version 0.1
Last modified Martin Iversen
*/

// GetRequestBody
/**
Generic request function
Takes in a writer and reader in addition to the API request url
The function handles the request and returns the body of the request
*/
func GetRequestBody(w http.ResponseWriter, r *http.Request) {

}

//GetRequestURL function returns the url of the request
func GetRequestURL(w http.ResponseWriter, r *http.Request) *url.URL {
	w.Header().Set("Content-Type", "application/json")

	link := r.URL.Path
	u, err := url.Parse(link)
	if err != nil {
		getErrorMessage(w, err)
	}

	return u
}
