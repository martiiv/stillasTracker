package endpoints

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Got payload")
	payload, _ := json.Marshal(r.Body)
	payloadString := string(payload[:])
	print("Payload:\n")
	print(payload)
	print("\n")
	print("Converted payload:\n")
	print(payloadString)
	/*
		if payloadBytes, err := hex.DecodeString(payloadString); err == nil {
			payload := ibs.Parse(payloadBytes)
			fmt.Println(payload)
		} else {
			fmt.Printf("Invalid hex string: %v", payloadString)
			fmt.Println(err)
		}
	*/
}
