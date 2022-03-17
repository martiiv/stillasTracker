package Database

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Ctx Initializing the context to be used with firebase
var Ctx context.Context

// Client Initializing the firebase client
var Client *firestore.Client

func databaseConnection() {
	// Creates instance of firebase
	Ctx = context.Background()
	sa := option.WithCredentialsFile("stillas-16563-firebase-adminsdk-wd82v-a9fe8919b7.json") //Initializes database
	app, err := firebase.NewApp(Ctx, nil, sa)
	if err != nil {
		log.Println("error occured when initializing database" + err.Error())
		_ = fmt.Errorf("error initializing app: %v", err)
	}

	Client, err = app.Firestore(Ctx) //Connects to the database
	if err != nil {
		log.Fatalln(err)
	}
	defer Client.Close()
}
