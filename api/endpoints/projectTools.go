package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
	"time"
)

func interfaceToInt(input interface{}) (int, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return 0, errors.New("cannot marshal")
	}

	var returnInt int
	err = json.Unmarshal(bytes, &returnInt)
	if err != err {
		return 0, errors.New("cannot unmarshal")
	}

	return returnInt, nil
}

func getScaffoldingInput(w http.ResponseWriter, r *http.Request) ([]_struct.Scaffolding, _struct.InputScaffoldingWithID, error) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, _struct.InputScaffoldingWithID{}, errors.New("could not read")
	}

	ok := checkTransaction(data)
	if !ok {
		return nil, _struct.InputScaffoldingWithID{}, errors.New("invalid body")
	}

	var inputScaffolding _struct.InputScaffoldingWithID
	err = json.Unmarshal(data, &inputScaffolding)
	if err != nil {
		return nil, _struct.InputScaffoldingWithID{}, errors.New("could not unmarshal")
	}
	var scaffolds []_struct.Scaffolding
	for i := range inputScaffolding.InputScaffolding {

		quantity := _struct.Quantity{
			Expected:   inputScaffolding.InputScaffolding[i].Quantity,
			Registered: 0,
		}

		scaffolding := _struct.Scaffolding{
			Type:     inputScaffolding.InputScaffolding[i].Type,
			Quantity: quantity,
		}

		scaffolds = append(scaffolds, scaffolding)
	}
	return scaffolds, inputScaffolding, nil
}

//iterateProjects will iterate through every project in active, inactive and upcoming projects.
func iterateProjects(id int, name string, state string) ([]*firestore.DocumentRef, error) {
	var documentReferences []*firestore.DocumentRef

	collection := projectCollection.Collections(database.Ctx)
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
			document = projectCollection.Collection(collRef.ID).Where(constants.P_ProjectName, "==", name).Documents(database.Ctx)
		} else if id != 0 {
			document = projectCollection.Collection(collRef.ID).Where(constants.P_ProjectId, "==", id).Documents(database.Ctx)
		} else {
			document = projectCollection.Collection(collRef.ID).Where(constants.P_State, "==", state).Documents(database.Ctx)
		}
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			documentReferences = append(documentReferences, documentRef.Ref)
		}
	}

	if documentReferences != nil {
		return documentReferences, nil
	} else {
		return nil, errors.New("could not find document")
	}

}

//checkStateBody will check if the body is of correct format, and if the values are correct datatypes.
func checkStateBody(body []byte) bool {
	var dat map[string]interface{}
	err := json.Unmarshal(body, &dat)
	if err != nil {
		return false
	}
	_, stateBool := dat[constants.P_State]
	_, idBool := dat[constants.P_idBody]
	_, isFloat := dat[constants.P_idBody].(float64)

	var correctValues bool
	if stateBool && idBool && isFloat {
		correctValues = checkState(dat[constants.P_State].(string))
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

	period := project[constants.P_Period]
	correctPeriod := checkPeriod(period)
	costumer := project[constants.P_Customer]
	correctCustomer := checkCustomer(costumer)
	geoFence := project[constants.P_Geofence]
	correctGeofence := checkGeofence(geoFence)
	address := project[constants.P_Address]

	correctAddress := checkAddressBody(address)
	_, longitudeFloat := project[constants.P_Longitude].(float64)
	_, latitudeFloat := project[constants.P_Latitude].(float64)
	_, sizeFloat := project[constants.P_Size].(float64)
	_, projectID := project[constants.P_ProjectId].(float64)
	_, projectName := project[constants.P_ProjectName].(string)
	_, state := project[constants.P_State].(string)

	if !state {
		return false
	}
	validState := checkState(project[constants.P_State].(string))
	correctFormat := validState && longitudeFloat && latitudeFloat && sizeFloat &&
		projectID && correctGeofence && correctCustomer && correctPeriod && projectName && correctAddress

	return correctFormat

}

//checkPeriod function that will verify the correct format of period struct
func checkPeriod(period interface{}) bool {
	periodByte, err := json.Marshal(period)
	if err != nil {
		return false
	}
	var periods map[string]interface{}
	err = json.Unmarshal(periodByte, &periods)
	if err != nil {
		return false
	}

	for _, i := range periods {
		_, err = time.Parse("02-01-2006", i.(string))
		if err != nil {
			return false
		}
	}

	nestedPeriod := []string{constants.P_PeriodstartDate, constants.P_PeriodendDate}
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
		return false
	}
	var customerMap map[string]interface{}
	err = json.Unmarshal(periodByte, &customerMap)
	if err != nil {
		return false
	}
	nestedPeriod := []string{constants.P_CustomerName, constants.P_CustomerEmail, constants.P_CustomerNumber}
	for _, period := range nestedPeriod {
		_, ok := customerMap[period]
		if !ok {
			return false
		}
	}

	_, numberFloat := customerMap[constants.P_CustomerNumber].(float64)
	if !numberFloat {
		return false
	}

	return true
}

//checkGeofence function that will verify the correct format of geofence struct
func checkGeofence(fence interface{}) bool {
	periodByte, err := json.Marshal(fence)
	if err != nil {
		return false
	}
	var geofenceMap map[string]interface{}
	err = json.Unmarshal(periodByte, &geofenceMap)
	if err != nil {
		return false
	}

	nestedPeriod := []string{constants.P_GeoW, constants.P_GeoX, constants.P_GeoY, constants.P_GeoZ}
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
		return false
	}
	var locationMap map[string]interface{}
	err = json.Unmarshal(periodByte, &locationMap)
	if err != nil {
		return false
	}

	coordinates := []string{constants.P_Longitude, constants.P_Latitude}
	for _, period := range coordinates {
		_, ok := locationMap[period]
		if !ok {
			return false
		}
	}

	_, longitudeFloat := locationMap[constants.P_Longitude].(float64)
	_, latitudeFloat := locationMap[constants.P_Latitude].(float64)

	if !latitudeFloat || !longitudeFloat {
		return false
	}

	return true
}

//checkState will check the value of the body, to ensure that the user has selected the correct state.
func checkState(input string) bool {
	state := []string{constants.P_Active, constants.P_Inactive, constants.P_Upcoming}

	return contains(state, input)
}

func checkTransaction(body []byte) bool {

	var inputScaffolding map[string]interface{}
	err := json.Unmarshal(body, &inputScaffolding)
	if err != nil {
		return false
	}

	_, toProject := inputScaffolding[constants.P_ToProjectID].(float64)
	_, fromProject := inputScaffolding[constants.P_fromProjectID].(float64)
	_, scaffold := inputScaffolding[constants.P_scaffold]
	scaffoldingInput := checkScaffoldingBody(inputScaffolding[constants.P_scaffold])

	return toProject && fromProject && scaffold && scaffoldingInput

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

	for i := range scaffoldMap {
		_, scaffoldingOk := scaffoldMap[i][constants.P_QuantityField].(float64)
		if !scaffoldingOk {
			return false
		}

		_, typeOk := scaffoldMap[i][constants.P_typeField].(string)
		if !typeOk {
			return false
		}

		for _, m2 := range scaffoldMap {
			isEqual := contains(constants.ScaffoldingTypes, m2[constants.P_typeField].(string))
			if !isEqual {
				return false
			}
		}
	}

	return true
}

//https://freshman.tech/snippets/go/check-if-slice-contains-element/
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//checkScaffoldingBody will check the body, to ensure the required fields are filled
func checkAddressBody(address interface{}) bool {

	periodByte, err := json.Marshal(address)
	if err != nil {
		return false
	}
	var addressMap map[string]interface{}
	err = json.Unmarshal(periodByte, &addressMap)
	if err != nil {
		return false
	}

	_, streetOk := addressMap[constants.P_AddressStreet]
	if !streetOk {
		return false
	}

	_, zipcodeOk := addressMap[constants.P_AddressZipCode]
	if !zipcodeOk {
		return false
	}

	_, municipalityOk := addressMap[constants.P_AddressMunicipality]
	if !municipalityOk {
		return false
	}

	_, countyOk := addressMap[constants.P_AddressCounty]
	if !countyOk {
		return false
	}

	return true
}
