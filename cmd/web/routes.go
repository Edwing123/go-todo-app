package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

// Returns the router for the server.
func (app *application) getRouter() http.Handler {
	router := pat.New()

	// Register app handlers
	router.Get("/", http.HandlerFunc(app.home))

	router.Get("/register", http.HandlerFunc(app.registerForm))
	router.Post("/auth/register", http.HandlerFunc(app.registerUser))

	router.Get("/login", http.HandlerFunc(app.loginForm))
	router.Get("/auth/login", http.HandlerFunc(app.loginUser))

	router.Get("/todos/list", http.HandlerFunc(app.todosManager))
	router.Get("/todos/create", http.HandlerFunc(app.createTodoForm))
	router.Post("/todos/create", http.HandlerFunc(app.createTodo))

	router.Post("/auth/logout", http.HandlerFunc(app.logoutUser))

	return router
}

// Display the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display home page"))
}

// Display the sign up form page.
func (app *application) registerForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display sign up form page"))
}

// Process user registration request.
func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register user"))
}

// Display the login in form page.
func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display login form page"))
}

// Process user login request.
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate user"))
}

// Display todos list page.
func (app *application) todosManager(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display todos manager page"))
}

// Display create todo page.
func (app *application) createTodoForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display create todo form page"))
}

// Process create new todo request.
func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new todo"))
}

// Process user logout request.
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout user"))
}
