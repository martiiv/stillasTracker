package endpoints

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"log"
	"net/http"
	"strings"
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

	messageByte, _ := json.Marshal(r.Body)
	if bytes, err := hex.DecodeString(string(messageByte)); err == nil {
		pkt := ibs.Parse(bytes)
		fmt.Println(pkt)

		if model, ok := pkt.ProductModel(); ok && strings.HasPrefix(model, "iBS") {
			fmt.Println(pkt)
		}
	} else {
		log.Printf("invalid payload: %s: %s", err, r.Body)
	}
}
