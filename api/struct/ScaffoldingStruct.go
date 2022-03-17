package _struct

/**
All structs taken from api modeling in postman
Version 0.5
Last edit: Martin Iversen 17.03.2022
*/

//ScaffoldingType will be used when getting info from scaffolding units
type ScaffoldingType struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Address   string  `json:"address"`
	} `json:"location"`
	BatteryLevel int `json:"batteryLevel"`
}

//ScaffoldHistory will be used when getting the history of scaffolding parts
type ScaffoldHistory struct {
	ID       int `json:"id"`
	Location []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Date      string  `json:"date"`
		Time      string  `json:"time"`
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
type AddScaffolding []struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	BatteryLevel int    `json:"batteryLevel"`
	Location     struct {
		Longitude interface{} `json:"longitude"`
		Latitude  interface{} `json:"latitude"`
	} `json:"location"`
}

//DeleteScaffolding used for deleteing scaffolding units
type DeleteScaffolding struct {
	Id int `json:"id"`
}
