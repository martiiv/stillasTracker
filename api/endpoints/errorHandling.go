package endpoints

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

//Error interface, will be used when handling err object
type Error interface {
	error
	Status() int
}

func getErrorMessage(error error) {

}
