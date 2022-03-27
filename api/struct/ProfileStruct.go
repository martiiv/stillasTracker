package _struct

type Employee struct {
	EmployeeID int `json:"employeeID"`
	Name       struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
	Role        string `json:"role"`
	Phone       any    `json:"phone"`
	Email       string `json:"email"`
	Admin       bool   `json:"admin"`
	Projects    []struct {
		ProjectID int `json:"projectID"`
	} `json:"projects"`
}

type ProfileDelete []struct {
	Id int `json:"id"`
}
