package _struct

/**
All structs taken from api modeling in postman
Version 0.5
Last edit: Martin Iversen 17.03.2022
*/

//ScaffoldingType will be used when getting info from scaffolding units
type ScaffoldingType struct {
	Type         string  `json:"type"`
	Project      string  `json:"project"`
	BatteryLevel float32 `json:"batteryLevel"`
	TagID        string  `json:"tagID"`
}

//ScaffoldHistory will be used when getting the history of scaffolding parts
type ScaffoldHistory struct {
	Id       string `json:"id"`
	Location []struct {
		Project string `json:"project"`
		Date    string `json:"date"`
		Time    string `json:"time"`
	} `json:"location"`
}

//MoveScaffolding will be used when moving scaffolding to new projects
type MoveScaffolding struct {
	Move []struct {
		Type     string `json:"type"`
		Quantity int    `json:"quantity"`
	} `json:"move"`
	ToProject   int `json:"toProject"`
	FromProject int `json:"fromProject"`
}

//AddScaffolding will be used when adding scaffolding units to the database

type DeleteScaffolding []struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type AddScaffolding []struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	BatteryLevel int    `json:"batteryLevel"`
	Project      string `json:"project"`
	TagID        string `json:"tagID"`
}
