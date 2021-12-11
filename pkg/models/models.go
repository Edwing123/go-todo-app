package models

// Represents a record from the user table.
type User struct {
	Id              int
	Username        string
	hashed_password string
}
