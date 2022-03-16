package apiTools

import (
	"log"
	"net/http"
)

/**
Class error handling, this class will contain all code concerning
handling of error messages
Code inspired by:
	https://blog.questionable.services/article/http-handler-error-handling-revisited/
	Page authored by Matt Silverlock from questionable serviecs
	Last visit: 08.03.2022

version 0.1
Last edited 08.03.2022 by Martin Iversen
*/

// Error Handled error. Object with both the error interface
//and the http status code
type Error interface {
	error
	Status() int
}

//StatusError represents error with HTTP code
type StatusError struct {
	Code int
	Err  error
}

/**
Function which lets an statuserror interface act as an error interface
Essentially only formatting
*/
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status
/**
Function returns StatusError status code
*/
func (se StatusError) Status() int {
	return se.Code
}

//getErrorMessage
/**
Function will return an appropriate error message when possible
Takes in the http Responsewriter (In order to give a response to the one doing the request)
And the potential error message
The method returns nothing however, appropriate error message will be printed in console and returned to the user
*/
func getErrorMessage(w http.ResponseWriter, err error) {
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
