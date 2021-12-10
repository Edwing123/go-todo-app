package main

import (
	"html/template"
	"log"
)

// The application struct hold the shared dependencies
// that will be used across the handlers and middlewares.
type application struct {
	templates  map[string]*template.Template
	infoLogger *log.Logger
	errLogger  *log.Logger
}

// the viewData struct holds per-view data, in order words,
// data that will be used in the app views.
type viewData struct {
	Year            int
	IsAuthenticated bool
	CSRFToken       string
}

// Represents the configuration command-line flags
type configFlags struct {
	addr        string // Server network address.
	dsn         string // Database data source name.
	tlsCertPath string // Path of the TLS certificate file.
	privKeyPath string // Path of the private key associated with the TLS certificate.
}
