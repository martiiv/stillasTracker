package tests

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
	"testing"
)

var Ctx context.Context
var Client *firestore.Client

func dataBaseTestConnection() {

	file, err := filepath.Abs("stillastestdatabase-firebase-adminsdk-tvp5e-1eb0fe0a3b.json")
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
Test_getScaffolding
Function for testing the get scaffolding endpoint
Creates a server and router before sending the request and getting the desired statuscode and output
*/
func Test_getScaffoldingAPITEST(t *testing.T) {
	dataBaseTestConnection()
	t.Run("IndividualScaffold", func(t *testing.T) {

	})
}
