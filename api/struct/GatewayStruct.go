package _struct

type Gateway struct {
	ProjectID   int     `json:"ProjectID"`
	Projectname string  `json:"ProjectName"`
	Status      bool    `json:"Status"`
	GatewayID   string  `json:"gatewayID"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
