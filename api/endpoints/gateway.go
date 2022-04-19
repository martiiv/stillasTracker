package endpoints

import (
	"encoding/json"
	"net/http"
)

/**
Class gateway
The class wil handle all information regarding the cellular gateways in the system
The class will contain the following functions:
	- getCentrals
	- addCentral
	- deleteCentral
	- fetchConnections

Version 0.1
Last modified Martin Iversen
*/

func GatewayRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	json.NewEncoder(w).Encode("Request sent to the gateway endpoint")
	json.NewEncoder(w).Encode(r.Body)
}
