package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
)

/*
Class databaseSetup.go created for communicating with database
Last update 19.05.2022
@version 1.0
*/
// Ctx Initializing the context to be used with firebase
var Ctx context.Context

// Client Initializing the firebase client
var Client *firestore.Client

//Code taken from https://firebase.google.com/docs/firestore/quickstart#go
func DatabaseConnection() {
	file, err := filepath.Abs("database/stillas-16563-firebase-adminsdk-wd82v-a9fe8919b7.json")
	if err != nil {
		log.Fatal(err)
	}
	// Creates instance of firebase
	Ctx = context.Background()
	sa := option.WithCredentialsFile(file) //Initializes database
	app, err := firebase.NewApp(Ctx, nil, sa)
	if err != nil {
		log.Println("error occured when initializing database" + err.Error())
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	Client, err = app.Firestore(Ctx) //Connects to the database
	if err != nil {
		log.Fatalln(err)
	}
}

/**
GetCollectionData Function inspired https://firebase.google.com/docs/firestore/quickstart#go

Will return all the documents from a specific collection.
iteratorRequest is the path to the collection of choice.
*/

func GetCollectionData(iteratorRequest *firestore.DocumentIterator) []map[string]interface{} {
	var documents []map[string]interface{}

	iter := iteratorRequest
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		documents = append(documents, doc.Data())
	}

	return documents
}

/**
GetDocumentData inspired from https://firebase.google.com/docs/firestore/query-data/get-data

The function will return a selected document from a selected collection.
document is the path to the selected document.
*/
func GetDocumentData(document *firestore.DocumentRef) (map[string]interface{}, error) {
	dsnap, err := document.Get(Ctx)
	if err != nil {
		return nil, err
	}
	m := dsnap.Data()
	return m, nil
}

/**
AddDocument inspired from https://firebase.google.com/docs/firestore/manage-data/add-data

Function will add a document to the database.
Document is the path to the new document.
structure is the data structure of how the data should be added into the database.
*/
func AddDocument(document *firestore.DocumentRef, structure map[string]interface{}) error {
	_, err := document.Set(Ctx, structure)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

/**
UpdateDocument inspired from https://firebase.google.com/docs/firestore/manage-data/add-data

Function will update fields of a selected document.

document is the path to a document.
update is the input fields we want to update.
*/
func UpdateDocument(document *firestore.DocumentRef, update []firestore.Update) error {
	_, err := document.Update(Ctx, update)
	if err != nil {
		return err
	}
	return nil
}

/**
DeleteDocument inspired from https://firebase.google.com/docs/firestore/manage-data/delete-data

Function will delete a selected document from the database.
document is the path to a selected document.
*/
func DeleteDocument(document *firestore.DocumentRef) error {
	_, err := document.Delete(Ctx)
	if err != nil {
		return err
	}
	return nil
}
