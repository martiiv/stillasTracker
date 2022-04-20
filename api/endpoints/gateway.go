package endpoints

import (
	"encoding/hex"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"io/ioutil"
	"net/http"
	"os"
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

	fmt.Println("\nString converted payload")
	payload, _ := ioutil.ReadAll(r.Body)
	convertedPayload := string(payload)
	payloadList := strings.Split(convertedPayload, "$")

	for _, v := range payloadList {
		print(v)
		fmt.Println("\nBeacon payload:")
		if m := igs.Parse(v); m != nil {
			fmt.Printf("Type:    %v\n", m.MsgType())
			fmt.Printf("Beacon:  %v\n", m.Beacon())
			fmt.Printf("Gateway: %v\n", m.Gateway())
			fmt.Printf("RSSI:    %v\n", m.RSSI())
			fmt.Printf("Payload: %v\n", m.Payload())

			if bytes, err := hex.DecodeString(m.Payload()); err == nil {
				p := ibs.Parse(bytes)
				fmt.Printf("Payload(parsed): %v\n", p)
			}
		} else {
			fmt.Println("Error: Invalid input message")
			fmt.Println(os.Args[1])
		}
	}
}
