package endpoints

import (
	"github.com/gorilla/mux"
	"net/http"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
)

func GatewayRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Defines data type
	w.Header().Set("Access-Control-Allow-Origin", "*") //Allows mobile and web application to access the api
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	lastElement := tool.GetLastUrlElement(r)
	baseCollection = database.Client.Doc(constants.G_GatewayCollection)

	if lastElement != constants.G_Gateway {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	requestType := r.Method //Defines the method of the request
	switch requestType {    //Forwards the request to the appropriate function
	case http.MethodGet:
		getGateway(w, r)
	case http.MethodPost:
		createGateway(w, r)
	case http.MethodPut:
		updateGateway(w, r)
	case http.MethodDelete:
		deleteGateway(w, r)
	}
}

func getGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := mux.Vars(r) //Gets queries

	switch true { //Forwards the request to the appropriate function based on the passed in query
	case query[constants.G_idURL] != "":
		getGatewayByID(w, r)
	case query[constants.G_ProjectID] != "" || query[constants.G_ProjectName] != "":
		getGatewayByProject(w, r)
	default:
		getAllGateways(w)
	}
}

/*
getAllGateways function gets all the gateways
*/
func getAllGateways(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var gateways []_struct.Gateways
}

func updateGateway(w http.ResponseWriter, r *http.Request) {

}

func createGateway(w http.ResponseWriter, r *http.Request) {

}

func deleteGateway(w http.ResponseWriter, r *http.Request) {

}

func getGatewayByProject(w http.ResponseWriter, r *http.Request) {

}

func getGatewayByID(w http.ResponseWriter, r *http.Request) {

}
