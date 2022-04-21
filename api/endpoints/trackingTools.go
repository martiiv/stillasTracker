package endpoints

import (
	"encoding/hex"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"io/ioutil"
	"net/http"
	"os"
	tool "stillasTracker/api/apiTools"
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
		if m := igs.Parse(payloadList[i]); m != nil {
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

func updateRegistered(gatewayList []*igs.Message, beaconID string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	project, err := http.NewRequest(http.MethodGet, "http://10.212.138.205:8080/stillastracking/v1/api/project?id="+beaconID, nil)
	if err != nil {
		tool.HandleError(tool.INVALIDREQUEST, w)
	}
	//TODO Hente ut prosjekt fra databasen og få tak i typen stillasdeler som skal oppdateres
	response, err := ioutil.ReadAll(project.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
	}

	print(response)

	for _, v := range gatewayList {
		tagID := v.Beacon()
		print(tagID)
	}
}

func addIDtoPart(m *igs.Message) {

}

func printFilteredGatewayInfo(gatewayList []*igs.Message, tagList []*ibs.Payload) {
	var printList []string
	for i := 0; i < len(tagList); i++ {
		tagInfo := gatewayList[i].Beacon()
		battery, _ := tagList[i].BatteryVoltage()
		printList = append(printList, "Tag ID:"+tagInfo+" battery voltage:"+strconv.FormatFloat(float64(battery), 'E', -1, 32)+"\n")
	}

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(tagList))
	fmt.Printf("List of tags:\n %v", printList)
	fmt.Printf("-----------------------------------------------------\n")
}
