package tests

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
	"stillasTracker/api/database"
)

/*
Function for establishing connection to the test database
Takes in a private key to the firestore database
*/
func dataBaseTestConnection() {

	file, err := filepath.Abs("stillastestdatabase-firebase-adminsdk-tvp5e-1eb0fe0a3b.json")
	if err != nil {
		log.Fatal(err)
	}

	// Creates instance of firebase
	database.Ctx = context.Background()
	sa := option.WithCredentialsFile(file) //Initializes database
	app, err := firebase.NewApp(database.Ctx, nil, sa)
	if err != nil {
		log.Println("error occured when initializing database" + err.Error())
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	database.Client, err = app.Firestore(database.Ctx) //Connects to the database
	if err != nil {
		log.Fatalln(err)
	}
}
