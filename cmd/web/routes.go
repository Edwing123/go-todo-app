package main

import (
	"errors"
	"net/http"

	"github.com/Edwing123/todo-app/pkg/forms"
	"github.com/Edwing123/todo-app/pkg/models"
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
	// - Session manager handler to load and save session data
	// - CSRF token provider/handler
	// - Sets auth status of the user to request context.
	dynamicMiddlewares := alice.New(app.sessionManager.LoadAndSave, noSurf, app.verifyUserAuth)

	// This middleware extends the dynamic one, plus it includes a middlware
	// to only allow access to authenticated users.
	dynamicwithRequireAuth := dynamicMiddlewares.Append(app.requireAuthentication)

	router := pat.New()

	// Register app handlers
	router.Get("/", dynamicMiddlewares.ThenFunc(app.home))

	router.Get("/register", dynamicMiddlewares.ThenFunc(app.registerForm))
	router.Post("/auth/register", dynamicMiddlewares.ThenFunc(app.registerUser))

	router.Get("/login", dynamicMiddlewares.ThenFunc(app.loginForm))
	router.Post("/auth/login", dynamicMiddlewares.ThenFunc(app.loginUser))

	router.Get("/todos/list", dynamicwithRequireAuth.ThenFunc(app.todosManager))
	router.Get("/todos/create", dynamicwithRequireAuth.ThenFunc(app.createTodoForm))
	router.Post("/todos/create", dynamicwithRequireAuth.ThenFunc(app.createTodo))

	router.Post("/auth/logout", dynamicwithRequireAuth.ThenFunc(app.logoutUser))

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
	app.render(w, r, "register", &viewData{
		Form: forms.New(nil),
	})
}

// Process user registration request.
func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("username", "password")
	form.NoWhiteSpace("username")
	form.MinLength("password", 10)

	// Re-render form if there's an error.
	if !form.IsValid() {
		app.render(w, r, "register", &viewData{
			Form: form,
		})

		return
	}

	// Insert user, and check for any errors
	// regarding duplicated usernames.
	err = app.userModel.Insert(form.Get("username"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicatedUsername) {
			form.Errors.Add("username", "This username is already taken")
			app.render(w, r, "register", &viewData{
				Form: form,
			})

			return
		}

		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Display the login in form page.
func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login", &viewData{Form: forms.New(nil)})
}

// Process user login request.
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("username", "password")
	form.NoWhiteSpace("username")

	// Form validation errors.
	if !form.IsValid() {
		app.render(w, r, "login", &viewData{
			Form: form,
		})

		return
	}

	// Credentials validation.
	id, err := app.userModel.Authenticate(form.Get("username"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Your username or password is incorrect")
			app.render(w, r, "login", &viewData{
				Form: form,
			})

			return
		}

		app.serverError(w, err)
		return
	}

	// Sucessfully login.
	// Add the loged-in user id to the session data.
	app.sessionManager.Put(r.Context(), "userID", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Display todos list page.
func (app *application) todosManager(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "todos", nil)
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
