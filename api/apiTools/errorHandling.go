package apiTools

import "net/http"

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

type ErrorStruct struct {
	message string
	code    int
}

var MARSHALLERROR = ErrorStruct{
	message: "could not json marshall",
	code:    http.StatusInternalServerError,
}

var NODOCUMENTSINDATABASE = ErrorStruct{
	message: "unable to find document",
	code:    http.StatusInternalServerError,
}

var NODOCUMENTWITHID = ErrorStruct{
	message: "no document with selected ID",
	code:    http.StatusNoContent,
}

var INVALIDREQUEST = ErrorStruct{
	message: "invalid request",
	code:    http.StatusInternalServerError,
}

var UNMARSHALLERROR = ErrorStruct{
	message: "could not jsonunmarshall",
	code:    http.StatusInternalServerError,
}

var NEWENCODERERROR = ErrorStruct{
	message: "could not encode data",
	code:    http.StatusInternalServerError,
}

var COLLECTIONITERATORERROR = ErrorStruct{
	message: "could not go through collection",
	code:    http.StatusInternalServerError,
}

var READALLERROR = ErrorStruct{
	message: "could not read input",
	code:    http.StatusBadRequest,
}

var INVALIDBODY = ErrorStruct{
	message: "invalid body",
	code:    http.StatusBadRequest,
}

var COULDNOTADDDOCUMENT = ErrorStruct{
	message: "could not add document",
	code:    http.StatusInternalServerError,
}

var CHANGESWERENOTMADE = ErrorStruct{
	message: "changes were not made",
	code:    http.StatusInternalServerError,
}

var COULDNOTFINDDATA = ErrorStruct{
	message: "could not find data in database",
	code:    http.StatusInternalServerError,
}

var CANNOTTRANSFERESCAFFOLDS = ErrorStruct{
	message: "cannot transfer the amount of scaffolding",
	code:    http.StatusInternalServerError,
}

func HandleError(err ErrorStruct, w http.ResponseWriter) {
	http.Error(w, err.message, err.code)

}
