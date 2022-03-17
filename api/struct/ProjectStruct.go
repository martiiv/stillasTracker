package _struct

type Period struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Address struct {
	Street       string `json:"street"`
	Zipcode      string `json:"zipcode"`
	Municipality string `json:"municipality"`
	County       string `json:"county"`
}

type Customer struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type Scaffolding struct {
	Units []struct {
		Type     string `json:"type"`
		Quantity struct {
			Expected   int `json:"expected"`
			Registered int `json:"registered"`
		} `json:"quantity"`
	} `json:"units"`
}

type Geofence struct {
	WPosition struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"w-position"`
	XPosition struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"x-position"`
	YPosition struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"y-position"`
	ZPosition struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"z-position"`
}

type Project struct {
	ProjectID   int     `json:"projectID"`
	ProjectName string  `json:"projectName"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Size        int     `json:"size"`
	State       string  `json:"state"`
	Period
	Address
	Customer
	Scaffolding
	Geofence
}

type IDStruct []struct {
	ID int `json:"id"`
}

type StateStruct struct {
	State string `json:"state"`
}

type MovingStruct []struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}
