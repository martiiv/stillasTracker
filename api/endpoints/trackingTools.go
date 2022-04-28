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
	var printlist []string

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
		} else {
			fmt.Println("Error: Invalid input message")
			fmt.Println(os.Args[1])
		}
	}
	idList, _ := getTagLists(gatewayList, beaconList)
	for i := range idList {
		batteryVoltage, _ := beaconList[i].BatteryVoltage()
		battery := strconv.FormatFloat(float64(batteryVoltage), 'E', -1, 32)

		printlist = append(printlist, "Tag id:"+idList[i]+" Battery voltage: "+battery)
	}
	updateAmountProject(gatewayList[0].Gateway(), w, idList)

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(idList))
	fmt.Printf("List of tags:\n %v", printlist)
	fmt.Printf("-----------------------------------------------------\n")
}

func updateAmountProject(beaconID string, w http.ResponseWriter, idList []string) {
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
	fmt.Printf("Succsessfully updated project with gateway id %v", beaconID)
}

func getProjectInfo(w http.ResponseWriter, beaconID string) _struct.GetProject {

	ProjectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)

	documentReference := gatewayCollection.Doc(beaconID)

	data, err := database.GetDocumentData(documentReference)
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		return _struct.GetProject{}
	}

	marshalled, err := json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return _struct.GetProject{}
	}

	var gateway _struct.Gateway
	err = json.Unmarshal(marshalled, &gateway)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return _struct.GetProject{}
	}

	if gateway.GatewayID == "" {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
	}

	project, err := IterateProjects(gateway.ProjectID, "", "")
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
	}
	newProject, _ := database.GetDocumentData(project[0])
	var projectStruct _struct.GetProject

	marshal, err := json.Marshal(newProject)
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
	fmt.Printf("%v", resultList)

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
	typeList := make(map[string]int)
	resultList := make(map[string]int)
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
