package endpoints

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"io/ioutil"
	"net/http"
	"os"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
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

func UpdatePosition(w http.ResponseWriter, r *http.Request) {
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

	scaffoldingLIst := getTagList(w, beaconID)

	print(scaffoldingLIst)

	for _, v := range gatewayList {
		tagID := v.Beacon()
		print(tagID)
	}
}

func getTags(w http.ResponseWriter) {

}

func getTagList(w http.ResponseWriter, beaconID string) _struct.ScaffoldingArray {
	ProjectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)
	project, err := http.NewRequest(http.MethodGet, "http://10.212.138.205:8080/stillastracking/v1/api/gateway?id="+beaconID, nil)
	if err != nil {
		tool.HandleError(tool.INVALIDREQUEST, w)
	}
	var responseStruct _struct.Gateway

	response, err := ioutil.ReadAll(project.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return nil
	}

	err = json.Unmarshal(marshal, &responseStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return nil
	}
	projectRef, err := http.NewRequest(http.MethodGet, "http://10.212.138.205:8080/stillastracking/v1/api/project?id="+strconv.Itoa(responseStruct.ProjectID)+"&scaffolding=true", nil)
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return nil
	}
	data, err := ioutil.ReadAll(projectRef.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return nil
	}

	var projectStruct _struct.GetProject

	marshal, err = json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return nil
	}
	err = json.Unmarshal(marshal, &projectStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return nil
	}

	scaffoldingList := projectStruct.ScaffoldingArray
	return scaffoldingList
}

func addIDtoPart(m *igs.Message) {

}

func printFilteredGatewayInfo(gatewayList []*igs.Message, tagList []*ibs.Payload) {
	var printList []string
	for i := 0; i < len(tagList); i++ {
		tagInfo := gatewayList[i].Beacon()
		runedPayload := []rune(tagInfo)
		tagID := string(runedPayload[6:12])
		battery, _ := tagList[i].BatteryVoltage()

		printList = append(printList, "Tag ID:"+tagID+" battery voltage:"+strconv.FormatFloat(float64(battery), 'E', -1, 32)+"\n")
	}

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(tagList))
	fmt.Printf("List of tags:\n %v", printList)
	fmt.Printf("-----------------------------------------------------\n")
}
