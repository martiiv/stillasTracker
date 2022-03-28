package _struct

// Period Start date and end date of the project.
type Period struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// Address of the project sight.
type Address struct {
	Street       string `json:"street"`
	Zipcode      string `json:"zipcode"`
	Municipality string `json:"municipality"`
	County       string `json:"county"`
}

// Customer information for the project.
type Customer struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Email  string `json:"email"`
}

// Scaffolding information at the project for expected and registered scaffolding units.
type Scaffolding struct {
	Units []struct {
		Type     string `json:"type"`
		Quantity struct {
			Expected   int `json:"expected"`
			Registered int `json:"registered"`
		} `json:"quantity"`
	} `json:"units"`
}

// Geofence at the project sight
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

// Project a collection of information.
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

// IDStruct to insert id of each project.
type IDStruct []struct {
	ID int `json:"id"`
}

// StateStruct to change the state of a project.
type StateStruct struct {
	ID    int    `json:"id"`
	State string `json:"state"`
}

// MovingStruct for moving a scaffolding piece.
type MovingStruct []struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type NewProject struct {
	ProjectID   int    `json:"projectID"`
	ProjectName string `json:"projectName"`
	Period      struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	} `json:"period"`
	Size      int     `json:"size"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Adresse   string  `json:"adresse"`
	Customer  struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
		Email  string `json:"email"`
	} `json:"customer"`
	Geofence struct {
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
	} `json:"geofence"`
}
