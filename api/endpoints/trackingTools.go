package endpoints

import (
	"cloud.google.com/go/firestore"
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
	var gatewayList []*igs.Message
	var beaconList []*ibs.Payload

	payload, _ := ioutil.ReadAll(r.Body)
	convertedPayload := string(payload)
	payloadList := strings.Split(convertedPayload, "\n")

	for i := 0; i < len(payloadList)-1; i++ {
		if m := igs.Parse(payloadList[i]); m != nil {
			gatewayList = append(gatewayList, m)
			if bytes, err := hex.DecodeString(m.Payload()); err == nil {
				p := ibs.Parse(bytes)
				beaconList = append(beaconList, p)
			}
			idList, batteryList := getTagLists(gatewayList, beaconList)
			updateAmountProject(gatewayList, gatewayList[i].Gateway(), w, idList, batteryList)

			idList = append(idList, "Tag ID:"+idList[i]+" battery voltage:"+strconv.FormatFloat(float64(batteryList[idList[i]]), 'E', -1, 32)+"\n")

			fmt.Printf("\n-----------------------------------------------------")
			fmt.Println("\nBeacon payload:")
			fmt.Printf("Time of POST: %v \n", time.Now())
			fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
			fmt.Printf("Amount of tags registered: %v \n", len(idList))
			fmt.Printf("List of tags:\n %v", idList)
			fmt.Printf("-----------------------------------------------------\n")
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

	scaffoldingLIst := getProjectInfo(w, beaconID)
	updatedProject := updateRegistered(w, scaffoldingLIst, idList)
	_, err := database.Client.Doc(constants.P_LocationCollection+"/"+constants.P_ProjectDocument).Collection(constants.P_Active).Doc(string(rune(updatedProject.ProjectID))).Set(database.Ctx, updatedProject, firestore.MergeAll)
	if err != nil {
		tool.HandleError(tool.DATABASEADDERROR, w)
	}
}

func getTags(w http.ResponseWriter) {

}

func getProjectInfo(w http.ResponseWriter, beaconID string) _struct.GetProject {
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

func getTagLists(gatewayList []*igs.Message, tagList []*ibs.Payload) ([]string, map[string]float32) {
	var tagIDList []string
	batteryList := make(map[string]float32)

	for i := 0; i < len(tagList); i++ {
		tagInfo := gatewayList[i].Beacon()
		runePayload := []rune(tagInfo)
		tagID := string(runePayload[6:12])
		battery, _ := tagList[i].BatteryVoltage()

		tagIDList = append(tagIDList, tagID)
		batteryList[tagID] = battery
	}
	return tagIDList, batteryList
}

func updateRegistered(w http.ResponseWriter, tagList _struct.GetProject, idList []string) _struct.GetProject {
	var updatedProject _struct.GetProject
	updatedProject.NewProject = tagList.NewProject

	resultList := getTagTypes(w, tagList, idList)
	for i := range tagList.ScaffoldingArray {
		scaffoldingType := tagList.ScaffoldingArray[i].Type
		newAmount := resultList[scaffoldingType]
		expected := tagList.ScaffoldingArray[i].Expected

		updatedProject.ScaffoldingArray[i].Type = scaffoldingType
		updatedProject.ScaffoldingArray[i].Expected = expected
		updatedProject.ScaffoldingArray[i].Registered = newAmount
	}
	return updatedProject
}

/*

 */
func getTagTypes(w http.ResponseWriter, projectList _struct.GetProject, idList []string) map[string]int {
	var typeList map[string]int
	var resultList map[string]int
	for i := range projectList.ScaffoldingArray {
		typeList[projectList.ScaffoldingArray[i].Type] = projectList.ScaffoldingArray[i].Registered
	}

	for j, _ := range typeList {
		counter := 0
		for i := range idList {
			objectPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(j).Doc(idList[i])

			_, err := database.GetDocumentData(objectPath)
			if err != nil {
				tool.HandleError(tool.DATABASEREADERROR, w)
				return nil
			}
			//TODO oppdater batterinivå her et sted
			counter = counter + 1
		}
		resultList[j] = counter
	}

	return resultList
}

func updateBatteryOnTag(w http.ResponseWriter, battery float32, scaffoldingRef *firestore.DocumentRef) {

}

func updateProjectOnTag() {

}
