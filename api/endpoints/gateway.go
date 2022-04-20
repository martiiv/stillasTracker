package endpoints

import (
	"encoding/hex"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

	payload, _ := ioutil.ReadAll(r.Body)
	convertedPayload := string(payload)
	payloadList := strings.Split(convertedPayload, "\n")

	var gatewayList []*igs.Message
	var beaconList []*ibs.Payload

	for i := 0; i < len(payloadList)-1; i++ {
		if m := igs.Parse(payloadList[0]); m != nil {
			gatewayList = append(gatewayList, m)
			if bytes, err := hex.DecodeString(m.Payload()); err == nil {
				p := ibs.Parse(bytes)
				beaconList = append(beaconList, p)
			}
			printFilteredGatewayInfo(gatewayList, beaconList)
		} else {
			fmt.Println("Error: Invalid input message")
			fmt.Println(os.Args[1])
		}
	}

}
func addGateway() {

}

func addIDtoPart(m *igs.Message) {

}

func printFilteredGatewayInfo(beaconList []*igs.Message, tagList []*ibs.Payload) {
	var printList []string
	for i := 0; i < len(tagList); i++ {
		tagInfo := beaconList[i].Beacon()
		battery, _ := tagList[i].BatteryVoltage()
		printList = append(printList, "Tag ID:"+tagInfo+" battery voltage:"+strconv.FormatFloat(float64(battery), 'E', -1, 32)+"\n")
	}

	fmt.Printf("\n------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", beaconList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(tagList))
	fmt.Printf("List of tags: %v", printList)
	fmt.Printf("------------------------------------------\n")
}

func printGatewayInfo(m *igs.Message) {
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Type:    %v\n", m.MsgType())
	fmt.Printf("Beacon:  %v\n", m.Beacon())
	fmt.Printf("Gateway: %v\n", m.Gateway())
	fmt.Printf("RSSI:    %v\n", m.RSSI())
	fmt.Printf("Payload: %v\n", m.Payload())
}
