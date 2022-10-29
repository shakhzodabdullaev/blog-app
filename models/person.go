package models

import "time"

type Person struct {
	ID        int        `json:"id"`
	FirstName string     `json:"fist_name"`
	LastName  string     `json:"last_name"`
	CreatedAt *time.Time `json:"craeted_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type PersonCreatedModel struct {
	FirstName string `json:"firs_name"`
	LastName  string `json:"last_name"`
}
type PersonUpdateModel struct {
	ID        string `json:"-"`
	FirstName string `json:"firs_name"`
	LastName  string `json:"last_name"`
}
