package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ingics/ingics-parser-go/ibs"
	"github.com/ingics/ingics-parser-go/igs"
	"google.golang.org/api/iterator"
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

	fmt.Printf("Recieved POST at %v \n", time.Now())

	var gatewayList []*igs.Message
	var beaconList []*ibs.Payload
	var printList []string

	payload, _ := ioutil.ReadAll(r.Body)
	convertedPayload := string(payload)
	payloadList := strings.Split(convertedPayload, "\n")
	fmt.Printf("\n%v\n", payloadList)
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
	idList, batteryList := getTagLists(gatewayList, beaconList)

	for i := range idList {

		batteryVoltage, _ := beaconList[i].BatteryVoltage()
		battery := strconv.FormatFloat(float64(batteryVoltage), 'E', -1, 32)

		printList = append(printList, "Tag id:"+idList[i]+" Battery voltage: "+battery)
	}

	fmt.Printf("\n-----------------------------------------------------")
	fmt.Println("\nBeacon payload:")
	fmt.Printf("Time of POST: %v \n", time.Now())
	fmt.Printf("Gateway: %v\n", gatewayList[0].Gateway())
	fmt.Printf("Amount of tags registered: %v \n", len(idList))
	fmt.Printf("List of tags:\n %v", printList)
	fmt.Printf("\n-----------------------------------------------------\n")

	updateAmountProject(gatewayList[0].Gateway(), w, idList, batteryList)
}

func updateAmountProject(beaconID string, w http.ResponseWriter, idList []string, batteryList map[string]float32) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	oldProject := getProjectInfo(w, beaconID)
	updatedProject := updateRegistered(w, oldProject, idList, batteryList)
	newMap := make(map[string]interface{})

	j, err := json.Marshal(updatedProject)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return
	}

	err = json.Unmarshal(j, &newMap)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}
	var update bool
	for i := range updatedProject.ScaffoldingArray {
		_, err := database.Client.Collection(constants.P_LocationCollection).Doc(constants.P_ProjectDocument).Collection(updatedProject.State).Doc(strconv.Itoa(updatedProject.ProjectID)).Collection(constants.P_StillasType).Doc(updatedProject.ScaffoldingArray[i].Type).Set(database.Ctx,
			newMap["scaffolding"].([]interface{})[i],
			firestore.MergeAll)

		if err != nil {
			tool.HandleError(tool.DATABASEADDERROR, w)
			update = false
		} else {
			update = true
		}
	}
	if update == true {
		fmt.Printf("Succsessfully updated project with gateway id %v\n", beaconID)
	} else {
		fmt.Printf("Unsuccsesful update")
	}
}

func getProjectInfo(w http.ResponseWriter, beaconID string) _struct.GetProject {
	ProjectCollection = database.Client.Doc(constants.P_LocationCollection + "/" + constants.P_ProjectDocument)
	gatewayCollection = database.Client.Collection(constants.G_GatewayCollection)

	data, err := database.GetDocumentData(gatewayCollection.Doc(beaconID))
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

	project := ProjectCollection.Collection(constants.P_Active).Doc(string(rune(gateway.ProjectID)))
	newProject, _ := database.GetDocumentData(project)
	var projectStruct _struct.NewProject

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

	var completeProject _struct.GetProject
	completeProject.NewProject = projectStruct

	scaffoldingParts, err := getScaffoldingStruct(project)
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
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

func updateRegistered(w http.ResponseWriter, oldProject _struct.GetProject, idList []string, batteryList map[string]float32) _struct.GetProject {
	var updatedProject _struct.GetProject

	updatedProject.NewProject = oldProject.NewProject
	updatedProject.ScaffoldingArray = oldProject.ScaffoldingArray
	resultList := getTagTypes(w, idList, updatedProject.ProjectName, batteryList)

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
func getTagTypes(w http.ResponseWriter, idList []string, projectName string, batteryList map[string]float32) map[string]int {
	resultList := make(map[string]int)
	for i := range idList {
		docRef, err := iterateScaffoldingParts(w, idList[i], projectName, batteryList)
		if err != nil {
			tool.HandleError(tool.NODOCUMENTWITHID, w)
		}
		data, err := database.GetDocumentData(docRef[0])
		if err != nil {
			tool.HandleError(tool.DATABASEREADERROR, w)
		}

		marshalled, _ := json.Marshal(data)
		if err != nil {
			tool.HandleError(tool.MARSHALLERROR, w)
		}
		var tagInfo _struct.ScaffoldingType
		err = json.Unmarshal(marshalled, &tagInfo)
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
		}
		resultList[tagInfo.Type] = resultList[tagInfo.Type] + 1
	}
	return resultList
}

func iterateScaffoldingParts(w http.ResponseWriter, scaffoldingID string, projectName string, batteryList map[string]float32) ([]*firestore.DocumentRef, error) {
	var documentReferences []*firestore.DocumentRef
	collection := database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collections(database.Ctx)

	for {
		collRef, err := collection.Next()
		if err == iterator.Done || err != nil {
			break
		}
		var document *firestore.DocumentSnapshot

		document, err = database.Client.Collection(constants.S_TrackingUnitCollection).Doc(constants.S_ScaffoldingParts).Collection(collRef.ID).Doc(scaffoldingID).Get(database.Ctx)
		if err != nil {
			break
		}
		_, err = document.Ref.Set(database.Ctx, map[string]interface{}{
			"project":      projectName,
			"batteryLevel": batteryList[scaffoldingID],
		}, firestore.MergeAll)

		if err != nil {
			tool.HandleError(tool.NODOCUMENTSINDATABASE, w)
		}
		documentReferences = append(documentReferences, document.Ref)
	}
	if documentReferences != nil {
		return documentReferences, nil
	} else {
		return nil, errors.New("could not find document")
	}
}

func convertBattery(batteryVoltage float32) int {

	return 0
}
