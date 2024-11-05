package models

type Employee struct {
	EmployeeID int    `json:"id"`
	Type       string `json:"type"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
