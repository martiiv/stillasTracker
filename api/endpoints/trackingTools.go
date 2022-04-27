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
			idList, batteryList := getTagLists(gatewayList, beaconList)
			print(idList, batteryList)

		} else {
			fmt.Println("Error: Invalid input message")
			fmt.Println(os.Args[1])
		}

	}
}

func updateAmountProject(gatewayList []*igs.Message, beaconID string, w http.ResponseWriter, idList []string, batteryList map[string]float32) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	scaffoldingLIst := getTagList(w, beaconID)
	updateRegistered(scaffoldingLIst, idList, batteryList)

	for _, v := range gatewayList {
		tagID := v.Beacon()
		print(tagID)
	}
}

func getTags(w http.ResponseWriter) {

}

func getTagList(w http.ResponseWriter, beaconID string) _struct.GetProject {
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
		return _struct.GetProject{}
	}

	err = json.Unmarshal(marshal, &responseStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return _struct.GetProject{}
	}
	projectRef, err := http.NewRequest(http.MethodGet, "http://10.212.138.205:8080/stillastracking/v1/api/project?id="+strconv.Itoa(responseStruct.ProjectID)+"&scaffolding=true", nil)
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return _struct.GetProject{}
	}
	data, err := ioutil.ReadAll(projectRef.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return _struct.GetProject{}
	}

	var projectStruct _struct.GetProject

	marshal, err = json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return _struct.GetProject{}
	}
	err = json.Unmarshal(marshal, &projectStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return _struct.GetProject{}
	}

	return projectStruct
}

func addIDtoPart(m *igs.Message) {

}

func getTagLists(gatewayList []*igs.Message, tagList []*ibs.Payload) ([]string, map[string]float32) {
	var printList []string
	var tagIDList []string
	var batteryList map[string]float32

	for i := 0; i < len(tagList); i++ {
		tagInfo := gatewayList[i].Beacon()
		runedPayload := []rune(tagInfo)
		tagID := string(runedPayload[6:12])
		battery, _ := tagList[i].BatteryVoltage()

		printList = append(printList, "Tag ID:"+tagID+" battery voltage:"+strconv.FormatFloat(float64(battery), 'E', -1, 32)+"\n")
		tagIDList = append(tagIDList, tagID)
		batteryList[tagID] = battery
	}

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(tagList))
	fmt.Printf("List of tags:\n %v", printList)
	fmt.Printf("-----------------------------------------------------\n")

	return tagIDList, batteryList
}

func updateRegistered(tagList _struct.GetProject, idList []string, batteryList map[string]float32) {

	//TODO Update the project with information from the batteryList and the IDList
	for i := range tagList.ScaffoldingArray {

	}
}

func getTagTypes(idList []string) map[string]int {
	//TODO use the list of tag ID's to create a map with the scaffoldingtype and the amount of that type

}
