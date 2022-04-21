package endpoints

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	tool "stillasTracker/api/apiTools"
	"stillasTracker/api/constants"
	"stillasTracker/api/database"
	_struct "stillasTracker/api/struct"
)

func GatewayRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Defines data type
	w.Header().Set("Access-Control-Allow-Origin", "*") //Allows mobile and web application to access the api
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	lastElement := tool.GetLastUrlElement(r)
	baseCollection = database.Client.Doc(constants.G_GatewayCollection)

	if lastElement != constants.G_Gateway {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	requestType := r.Method //Defines the method of the request
	switch requestType {    //Forwards the request to the appropriate function
	case http.MethodGet:
		getGateway(w, r)
	case http.MethodPost:
		createGateway(w, r)
	case http.MethodPut:
		updateGateway(w, r)
	case http.MethodDelete:
		deleteGateway(w, r)
	}
}

func getGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := mux.Vars(r) //Gets queries

	switch true { //Forwards the request to the appropriate function based on the passed in query
	case query[constants.G_idURL] != "":
		getGatewayByID(w, r)
	case query[constants.G_ProjectID] != "" || query[constants.G_ProjectName] != "":
		getGatewayByProject(w, r)
	default:
		getAllGateways(w)
	}
}

func updateGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tool.HandleError(tool.READALLERROR, w)
		return
	}

	batch := database.Client.Batch()

	var gatewayStruct map[string]interface{}
	err = json.Unmarshal(data, &gatewayStruct)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	gateway := gatewayStruct[constants.G_GatewayID].(string)

	documentReference, err := iterateGateways(gateway)
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	var updates []firestore.Update

	for s, i := range gatewayStruct {
		update := firestore.Update{
			Path:  s,
			Value: i,
		}

		updates = append(updates, update)
	}

	for _, ref := range documentReference {
		batch.Update(ref, updates)
	}
	_, err = batch.Commit(database.Ctx) //Commits the database changes if all changes pass
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	}
}

func createGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	requestBody, err := ioutil.ReadAll(r.Body)
	var gateway _struct.Gateway

	err = json.Unmarshal(requestBody, &gateway)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	_, err = iterateGateways(gateway.GatewayID)
	if err != nil {
		tool.HandleError(tool.CouldNotAddSameID, w)
		return
	}

	documentPath := baseCollection
	var firebaseInput map[string]interface{}

	err = json.Unmarshal(requestBody, &firebaseInput)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
	}

	err = database.AddDocument(documentPath, firebaseInput)
	if err != nil {
		tool.HandleError(tool.COULDNOTADDDOCUMENT, w)
		return
	} else {
		tool.HandleError(tool.ADDED, w)
	}
}

func deleteGateway(w http.ResponseWriter, r *http.Request) {

}

func getGatewayByProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queries := mux.Vars(r)
	query := ""
	if queries[constants.G_ProjectID] != "" {
		query = queries[constants.G_ProjectID]
	} else if queries[constants.G_ProjectName] != "" {
		query = queries[constants.G_ProjectName]
	} else {
		tool.HandleError(tool.INVALIDREQUEST, w)
		return
	}

	documentPath := baseCollection.Collection(query).Documents(database.Ctx)
	var gateways []_struct.Gateway

	for {
		documentRef, err := documentPath.Next()
		if err == iterator.Done {
			break
		}

		var gateway _struct.Gateway
		doc, _ := database.GetDocumentData(documentRef.Ref)
		gatewayByte, err := json.Marshal(doc)
		err = json.Unmarshal(gatewayByte, &gateway)
		if err != nil {
			tool.HandleError(tool.UNMARSHALLERROR, w)
			return
		}

		gateways = append(gateways, gateway)
	}
	if gateways == nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	err := json.NewEncoder(w).Encode(gateways)
	if err != nil {
		tool.HandleError(tool.NEWENCODERERROR, w)
		return
	}
}

func getGatewayByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queries := mux.Vars(r)

	documentReference, err := iterateGateways(queries[constants.G_idURL])
	if err != nil {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
		return
	}

	data, _ := database.GetDocumentData(documentReference[0])
	marshalled, err := json.Marshal(data)
	if err != nil {
		tool.HandleError(tool.MARSHALLERROR, w)
		return
	}

	var gateway _struct.Gateway
	err = json.Unmarshal(marshalled, &gateway)
	if err != nil {
		tool.HandleError(tool.UNMARSHALLERROR, w)
		return
	}

	if gateway.GatewayID == "" {
		tool.HandleError(tool.COULDNOTFINDDATA, w)
	}

	err = json.NewEncoder(w).Encode(gateway)
	if err != nil {
		return
	}
}

func iterateGateways(id string) ([]*firestore.DocumentRef, error) {
	var documentReferences []*firestore.DocumentRef

	collection := baseCollection.Collections(database.Ctx)
	for {
		collref, err := collection.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}

		var document *firestore.DocumentIterator
		document = baseCollection.Collection(collref.ID).Where(constants.G_idURL, "==", id).Documents(database.Ctx)
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

/*
getAllGateways function gets all the gateways
*/
func getAllGateways(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var gateways []_struct.Gateway

	collection := baseCollection.Collections(database.Ctx)
	for {
		collRef, err := collection.Next()
		if err == iterator.Done || err != nil {
			break
		}

		document := baseCollection.Collection(collRef.ID).Documents(database.Ctx)
		for {
			documentRef, err := document.Next()
			if err == iterator.Done {
				break
			}

			var gateway _struct.Gateway
			doc, _ := database.GetDocumentData(documentRef.Ref)
			gatewayByte, err := json.Marshal(doc)
			err = json.Unmarshal(gatewayByte, &gateway)
			if err != nil {
				tool.HandleError(tool.UNMARSHALLERROR, w)
				return
			}
			gateways = append(gateways, gateway)
		}
	}
	err := json.NewEncoder(w).Encode(gateways)
	if err != nil {
		return
	}
}
