package endpoints

import (
	"encoding/hex"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"net/http"
	"os"
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

	for _, payloadHex := range os.Args[1:] {
		if payloadBytes, err := hex.DecodeString(payloadHex); err == nil {
			payload := ibs.Parse(payloadBytes)
			fmt.Println(payload)
		} else {
			fmt.Printf("Invalid hex string: %v", payloadHex)
			fmt.Println(err)
		}
	}
}
