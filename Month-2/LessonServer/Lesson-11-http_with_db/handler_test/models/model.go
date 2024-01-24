package models

type User struct {
	ID        string `json:"id"`
	IdNum     int    `json:"idnum"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
