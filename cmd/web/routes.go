package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// Returns the router for the server.
func (app *application) getRouter() http.Handler {
	// The base middlewares chain includes:
	// - Panic recovery
	// - Request logging
	// - Addition of secure headers to response
	baseMiddlewares := alice.New(app.recoverPanic, app.logRequest, addSecureHeaders)

	// The dynamic middlewares chain includes:
	// - CSRF token provider/handler
	dynamicMiddlewares := alice.New(noSurf)

	router := pat.New()

	// Register app handlers
	router.Get("/", dynamicMiddlewares.ThenFunc(app.home))

	router.Get("/register", dynamicMiddlewares.ThenFunc(app.registerForm))
	router.Post("/auth/register", dynamicMiddlewares.ThenFunc(app.registerUser))

	router.Get("/login", dynamicMiddlewares.ThenFunc(app.loginForm))
	router.Get("/auth/login", dynamicMiddlewares.ThenFunc(app.loginUser))

	router.Get("/todos/list", dynamicMiddlewares.ThenFunc(app.todosManager))
	router.Get("/todos/create", dynamicMiddlewares.ThenFunc(app.createTodoForm))
	router.Post("/todos/create", dynamicMiddlewares.ThenFunc(app.createTodo))

	router.Post("/auth/logout", dynamicMiddlewares.ThenFunc(app.logoutUser))

	// Static assets file server
	fileServer := http.StripPrefix("/static", app.serveStaticFiles(http.FileServer(http.Dir("./ui/assets"))))
	router.Get("/static/", fileServer)

	return baseMiddlewares.Then(router)
}

// Display the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", nil)
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
