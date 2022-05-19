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

Version 1.0
Last modified Martin Iversen 19.05.2022
*/
func UpdatePosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var gatewayList []*igs.Message
	var beaconList []*ibs.Payload
	var printList []string

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
			InfoLogger.Printf("Invalid input message: %v", os.Args[1])
		}
	}
	idList, batteryList := getTagLists(gatewayList, beaconList)

	for i := range idList {

		batteryVoltage, _ := beaconList[i].BatteryVoltage()
		battery := strconv.FormatFloat(float64(batteryVoltage), 'E', -1, 32)

		printList = append(printList, "\nTag id:"+idList[i]+" Battery voltage: "+battery+"\n")
	}
	updateAmountProject(gatewayList[0].Gateway(), w, idList, batteryList)

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now().Format(time.RFC822))
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(idList))
	fmt.Printf("List of tags:\n %v", printList)
	fmt.Printf("\n-----------------------------------------------------\n")
}

func updateAmountProject(beaconID string, w http.ResponseWriter, idList []string, batteryList map[string]float32) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	oldProject := getProjectInfo(w, beaconID)
	updatedProject := updateRegistered(oldProject, idList, batteryList)
	newMap := make(map[string]interface{})

	j, err := json.Marshal(updatedProject)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		ErrorLogger.Printf("Failed to unmarshal object: %v", updatedProject)
		return
	}

	err = json.Unmarshal(j, &newMap)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		ErrorLogger.Printf("Failed to unmarshal object: %v", updatedProject)
		return
	}
	var update bool
	for i := range updatedProject.ScaffoldingArray {
		doc, err := database.Client.Collection(constants.P_LocationCollection).Doc(constants.P_ProjectDocument).Collection(updatedProject.State).Doc(strconv.Itoa(updatedProject.ProjectID)).Collection(constants.P_StillasType).Doc(updatedProject.ScaffoldingArray[i].Type).Set(database.Ctx,
			newMap["scaffolding"].([]interface{})[i],
			firestore.MergeAll)

		if err != nil {
			tool.HandleError(tool.DATABASEADDERROR, w)
			ErrorLogger.Printf("Database ADD error on object %v", doc)
			update = false
		} else {
			update = true
		}
	}
	if update == true {
		DatabaseLogger.Printf("Successfully updated project with gateway id %v\n", beaconID)
	} else {
		DatabaseLogger.Printf("Unsuccessfully updated project with gateway id %v\n", beaconID)
	}
}

func getProjectInfo(w http.ResponseWriter, beaconID string) _struct.GetProject {
	ProjectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)
	gatewayCollection = database.Client.Collection(constants.G_GatewayCollection)

	data, err := database.GetDocumentData(gatewayCollection.Doc(beaconID))
	if err != nil {
		tool.HandleError(tool.NODOCUMENTWITHID, w)
		ErrorLogger.Printf("Could not get document %v from the database", beaconID)
		return _struct.GetProject{}
	}

	marshalled, err := json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		ErrorLogger.Printf("Marshall error on object %v", data)
		return _struct.GetProject{}
	}
	var gateway _struct.Gateway
	err = json.Unmarshal(marshalled, &gateway)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		ErrorLogger.Printf("Unmarshall error on object %v", marshalled)
		return _struct.GetProject{}
	}

	project := ProjectCollection.Collection(constants.P_Active).Doc(strconv.Itoa(gateway.ProjectID))
	newProject, err := database.GetDocumentData(project)
	if err != nil {
		tool.HandleError(tool.DATABASEREADERROR, w)
		ErrorLogger.Printf("Error occurred when fetching document %v", gateway.ProjectID)
	}
	var projectStruct _struct.NewProject

	marshal, err := json.Marshal(newProject)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		ErrorLogger.Printf("Error occurred when marshalling object: %v", newProject)
		return _struct.GetProject{}
	}
	err = json.Unmarshal(marshal, &projectStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		ErrorLogger.Printf("Error occurred when unmarshalling object: %v", marshal)
		return _struct.GetProject{}
	}

	var completeProject _struct.GetProject
	completeProject.NewProject = projectStruct

	scaffoldingParts, err := getScaffoldingStruct(project)
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		ErrorLogger.Printf("Could not find scaffolding parts on project:", project)
		return _struct.GetProject{}
	}
	completeProject.ScaffoldingArray = scaffoldingParts
	return completeProject

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

func updateRegistered(oldProject _struct.GetProject, idList []string, batteryList map[string]float32) _struct.GetProject {
	var updatedProject _struct.GetProject

	updatedProject.NewProject = oldProject.NewProject
	updatedProject.ScaffoldingArray = oldProject.ScaffoldingArray
	projectName := updatedProject.NewProject.ProjectName

	resultList := getTagTypes(idList, projectName, oldProject.ScaffoldingArray, batteryList)

	for i := range updatedProject.ScaffoldingArray {
		scaffoldingType := oldProject.ScaffoldingArray[i].Type
		expected := oldProject.ScaffoldingArray[i].Quantity.Expected
		if resultList[scaffoldingType] != 0 {
			updatedProject.ScaffoldingArray[i].Quantity.Registered = resultList[scaffoldingType]
		} else {
			updatedProject.ScaffoldingArray[i].Quantity.Registered = 0
		}
		updatedProject.ScaffoldingArray[i].Type = scaffoldingType
		updatedProject.ScaffoldingArray[i].Quantity.Expected = expected
	}
	return updatedProject
}

/*

 */
func getTagTypes(idList []string, projectName string, scaffoldingArray _struct.ScaffoldingArray, batteryList map[string]float32) map[string]int {
	resultList := make(map[string]int)
	var scaffoldingType string
	for i := range idList {
		for j := range scaffoldingArray {
			scaffoldingType = scaffoldingArray[j].Type
			documentPath := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(scaffoldingType).Doc(idList[i])
			_, err := documentPath.Update(database.Ctx, []firestore.Update{
				{
					Path:  "project",
					Value: projectName,
				}, {
					Path:  "batteryLevel",
					Value: batteryList[idList[i]],
				},
			})
			if err != nil {
				ErrorLogger.Printf("Document with id: %v is not in scaffoldingType collection %n", idList[i], scaffoldingType)
			} else if err == nil {
				DatabaseLogger.Printf("Succsessfully updated scaffolding part: %v", idList[i])
				resultList[scaffoldingType] = resultList[scaffoldingType] + 1
			}
		}
	}
	return resultList
}
