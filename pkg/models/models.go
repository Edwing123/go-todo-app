package models

import "time"

// Represents a record from the user table.
type User struct {
	Id              int
	Username        string
	hashed_password string
}

// Represents a record from the todo table.
type Todo struct {
	Id        int
	Title     string
	Done      bool
	CreatedAt time.Time
	Expires   time.Time
}
