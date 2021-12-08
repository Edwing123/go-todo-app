package main

import (
	"database/sql"
	"html/template"
	"log"
)

// The application struct hold the shared dependencies
// that will be used across the handlers and middlewares.
type application struct {
	db         *sql.DB
	templates  map[string]*template.Template
	infoLogger *log.Logger
	errLogger  *log.Logger
}

// the viewData struct holds per-view data, in order words,
// data that will be used in the app views.
type viewData struct {
	IsAuthenticated bool
	CSRFToken       string
}

// Represents the configuration command-line flags
type configFlags struct {
	addr string // Server network address
	dsn  string // Database data source name
}
