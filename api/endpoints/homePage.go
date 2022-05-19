package endpoints

import (
	"encoding/json"
	"net/http"
	tool "stillasTracker/api/apiTools"
)

/**
Class homepage.go created but never used,
@version 1.0
last edit 19.05.2022
*/

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	welcomeMessage := "Welcome to the scaffolding tracker API!\n " +
		"The api uses enpoints in order to create, remove and move scaffolding parts in between construction sites\n" +
		"The api was created by Tormod Mork Müller, ALeksander Aaboen and Martin Iversen\n" +
		"The api was created along during a bachelor-thesis at NTNU in 2022, " +
		"you can find the official api guide in the readme file in the gitlab repo of the project: https://git.gvk.idi.ntnu.no/aleksaab/stillastracker/-/tree/main/api" +
		"You have access to the following endpoints: \n"

	scaffoldingEndpoints := "To insert scaffolding parts into the database, use the following endpoints:" +
		"get all scaffolding parts: http://10.212.138.205:8080/stillastracking/v1/api/unit\n" +
		"get a specific scaffolding part: "

	err := json.NewEncoder(w).Encode(welcomeMessage)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}
	err = json.NewEncoder(w).Encode(scaffoldingEndpoints)
	if err != nil {
		tool.HandleError(tool.ENCODINGERROR, w)
		return
	}
}
