package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"net/url"
	"stillasTracker/api/Database"
	tool "stillasTracker/api/apiTools"
	_struct "stillasTracker/api/struct"
	"strings"
)

func createPath(segments []string) string {
	var finalPath string
	for _, s := range segments {
		finalPath += s + "/"
	}
	finalPath = strings.TrimRight(finalPath, "/")
	return finalPath
}

func getQuery(r *http.Request) url.Values {
	query := r.URL.Query()
	if len(query) != 1 {
		return nil
	}
	switch true {
	case query.Has("name"),
		query.Has("id"):
		return query
	}
	return nil
}

func getScaffoldingInput(w http.ResponseWriter, r *http.Request) ([]_struct.Scaffolding, _struct.InputScaffoldingWithID) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return nil, _struct.InputScaffoldingWithID{}
	}
	ok := checkTransaction(data)
	if !ok {
		http.Error(w, "body invalid", http.StatusBadRequest)
		return nil, _struct.InputScaffoldingWithID{}
	}

	var inputScaffolding _struct.InputScaffoldingWithID
	err = json.Unmarshal(data, &inputScaffolding)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return nil, _struct.InputScaffoldingWithID{}
	}
	var scaffolds []_struct.Scaffolding
	for i := range inputScaffolding.InputScaffolding {

		quantity := _struct.Quantity{
			Expected:   inputScaffolding.InputScaffolding[i].Quantity,
			Registered: 0,
		}

		scaffolding := _struct.Scaffolding{
			Category: inputScaffolding.InputScaffolding[i].Type,
			Quantity: quantity,
		}

		scaffolds = append(scaffolds, scaffolding)
	}
	return scaffolds, inputScaffolding
}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProjects(id int, name string) (*firestore.DocumentRef, tool.ErrorStruct) {
	var documentReference *firestore.DocumentRef
	collection := projectCollection.Collections(Database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		var document *firestore.DocumentIterator
		if name != "" {
			document = projectCollection.Collection(collRef.ID).Where("projectName", "==", name).Documents(Database.Ctx)
		} else {
			document = projectCollection.Collection(collRef.ID).Where("projectID", "==", id).Documents(Database.Ctx)
		}
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			documentReference = documentRef.Ref
		}
	}

	if documentReference != nil {
		return documentReference, tool.ErrorStruct{}
	}
	return nil, tool.NODOCUMENTWITHID
}

//checkStateBody will check if the body is of correct format, and if the values are correct datatypes.
func checkStateBody(body []byte) bool {
	var dat map[string]interface{}
	err := json.Unmarshal(body, &dat)
	if err != nil {
		return false
	}
	_, stateBool := dat["state"]
	_, idBool := dat["id"]
	_, isFloat := dat["id"].(float64)

	var correctValues bool
	if stateBool && idBool && isFloat {
		correctValues = checkState(dat["state"].(string))
	}
	return correctValues
}

//checkProjectBody function that will verify the correct format of project struct
func checkProjectBody(body []byte) bool {
	var project map[string]interface{}
	err := json.Unmarshal(body, &project)
	if err != nil {
		return false
	}

	period := project["period"]
	correctPeriod := checkPeriod(period)
	costumer := project["customer"]
	correctCustomer := checkCustomer(costumer)
	geoFence := project["geofence"]
	correctGeofence := checkGeofence(geoFence)
	_, longitudeFloat := project["longitude"].(float64)
	_, latitudeFloat := project["latitude"].(float64)
	_, sizeFloat := project["size"].(float64)
	_, projectID := project["projectID"].(float64)
	_, projectName := project["projectName"].(string)
	_, state := project["state"].(string)

	if !state {
		return false
	}
	validState := checkState(project["state"].(string))
	correctFormat := validState && longitudeFloat && latitudeFloat && sizeFloat &&
		projectID && correctGeofence && correctCustomer && correctPeriod && projectName

	return correctFormat

}

//checkPeriod function that will verify the correct format of period struct
func checkPeriod(period interface{}) bool {
	periodByte, err := json.Marshal(period)
	if err != nil {
		fmt.Println(err.Error())
	}
	var periods map[string]interface{}
	err = json.Unmarshal(periodByte, &periods)
	if err != nil {
		return false
	}

	nestedPeriod := []string{"startDate", "endDate"}
	for _, period := range nestedPeriod {
		_, ok := periods[period]
		if !ok {
			return false
		}
	}

	return true
}

//checkCustomer function that will verify the correct format of customer struct
func checkCustomer(customer interface{}) bool {
	periodByte, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err.Error())
	}
	var customerMap map[string]interface{}
	err = json.Unmarshal(periodByte, &customerMap)
	if err != nil {
		return false
	}
	nestedPeriod := []string{"name", "email", "number"}
	for _, period := range nestedPeriod {
		_, ok := customerMap[period]
		if !ok {
			return false
		}
	}

	_, numberFloat := customerMap["number"].(float64)
	if !numberFloat {
		return false
	}

	return true
}

//checkGeofence function that will verify the correct format of geofence struct
func checkGeofence(fence interface{}) bool {
	periodByte, err := json.Marshal(fence)
	if err != nil {
		fmt.Println(err.Error())
	}
	var geofenceMap map[string]interface{}
	err = json.Unmarshal(periodByte, &geofenceMap)
	if err != nil {
		return false
	}

	nestedPeriod := []string{"w-position", "x-position", "y-position", "z-position"}
	for _, period := range nestedPeriod {
		_, ok := geofenceMap[period]
		if !ok {
			return false
		} else {
			correctCoordinate := checkGeofenceCoordinates(geofenceMap[period])
			if !correctCoordinate {
				return false
			}
		}
	}

	return true
}

//checkGeofenceCoordinates function that will verify the correct format of geofence position struct
func checkGeofenceCoordinates(location interface{}) bool {
	periodByte, err := json.Marshal(location)
	if err != nil {
		fmt.Println(err.Error())
	}
	var locationMap map[string]interface{}
	err = json.Unmarshal(periodByte, &locationMap)
	if err != nil {
		return false
	}

	coordinates := []string{"longitude", "latitude"}
	for _, period := range coordinates {
		_, ok := locationMap[period]
		if !ok {
			return false
		}
	}

	_, longitudeFloat := locationMap["longitude"].(float64)
	_, latitudeFloat := locationMap["latitude"].(float64)

	if !latitudeFloat || !longitudeFloat {
		return false
	}

	return true
}

//checkState will check the value of the body, to ensure that the user has selected the correct state.
func checkState(input string) bool {
	state := []string{"Active", "Inactive", "Upcoming"}

	var correctValues bool
	for _, states := range state {
		if input == states {
			correctValues = true
			break
		}
	}
	return correctValues
}

func checkTransaction(body []byte) bool {

	var inputScaffolding map[string]interface{}
	err := json.Unmarshal(body, &inputScaffolding)
	if err != nil {
		return false
	}

	_, toProject := inputScaffolding["toProjectID"].(float64)
	_, fromProject := inputScaffolding["fromProjectID"].(float64)
	_, scaffold := inputScaffolding["scaffold"]
	scaffoldingInput := checkScaffoldingBody(inputScaffolding["scaffold"])

	return (toProject && fromProject && scaffold && scaffoldingInput)

}

//checkScaffoldingBody will check the body, to ensure the required fields are filled
func checkScaffoldingBody(scaffold interface{}) bool {

	periodByte, err := json.Marshal(scaffold)
	if err != nil {
		return false
	}
	var scaffoldMap []map[string]interface{}
	err = json.Unmarshal(periodByte, &scaffoldMap)
	if err != nil {
		return false
	}

	for i, m := range scaffoldMap {
		fmt.Println(i, m)
		_, scaffoldingOk := scaffoldMap[i]["quantity"]
		if !scaffoldingOk {
			return false
		}
		_, typeOk := scaffoldMap[i]["type"]
		if !typeOk {
			return false
		}
	}

	return true
}
