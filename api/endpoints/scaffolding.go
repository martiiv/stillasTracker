package endpoints

import (
	"net/http"
	"stillasTracker/api/apiTools"
)

/**
Class scaffolding
This class will contain all functions used for the handling of scaffolding units
The class contains the following functions:
	- addScaffolding:    Function lets a user add a scaffolding part to the system
	- deleteScaffolding: Function removes a scaffolding unit from the system
	- moveScaffold:      Function lets a user move scaffolding parts to a new project
	- getScaffoldingUnit Function returns information about a scaffolding part
	- getUnitHistory     Function returns the history of a scaffolding part

Version 0.1
Last modified Martin Iversen
*/
func scaffoldingRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := apiTools.GetRequestURL(w, r)
	switch url {

	}

}
