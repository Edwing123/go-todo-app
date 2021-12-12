package main

import (
	"html/template"
	"log"

	"github.com/Edwing123/todo-app/pkg/forms"
	"github.com/Edwing123/todo-app/pkg/models"
	"github.com/Edwing123/todo-app/pkg/models/mysql"
	"github.com/alexedwards/scs/v2"
)

// The application struct hold the shared dependencies
// that will be used across the handlers and middlewares.
type application struct {
	templates      map[string]*template.Template
	infoLogger     *log.Logger
	errLogger      *log.Logger
	sessionManager *scs.SessionManager
	userModel      *mysql.UserModel
}

// the viewData struct holds per-view data, in order words,
// data that will be used in the app views.
type viewData struct {
	Year            int
	IsAuthenticated bool
	CSRFToken       string
	Form            *forms.Form
	User            *models.User
}

// Represents the configuration command-line flags
type configFlags struct {
	addr        string // Server network address.
	dsn         string // Database data source name.
	tlsCertPath string // Path of the TLS certificate file.
	privKeyPath string // Path of the private key associated with the TLS certificate.
}
