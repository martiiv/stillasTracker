package endpoints

import "net/http"

func GatewayRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Defines data type
	w.Header().Set("Access-Control-Allow-Origin", "*") //Allows mobile and web application to access the api
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

}
