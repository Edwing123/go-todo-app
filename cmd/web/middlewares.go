package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Edwing123/todo-app/pkg/models"
	"github.com/justinas/nosurf"
)

// Log request important details.
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLogger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

// Inject secure headers to every response.
func addSecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add security headers to response headers
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		// call next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// Recover from any panic error produced within the goroutine.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Create Nosurf CSRF token handler.
func noSurf(next http.Handler) http.Handler {
	CSRFHandler := nosurf.New(next)
	CSRFHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	return CSRFHandler
}

// TODO: think about a good description :|
func (app *application) verifyUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the existance of the userID key.
		if !app.sessionManager.Exists(r.Context(), "userID") {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the user with that id exists.
		_, err := app.userModel.GetById(app.sessionManager.GetInt(r.Context(), "userID"))
		if err != nil {
			if errors.Is(err, models.ErrNoRecords) {
				app.sessionManager.Remove(r.Context(), "userID")
				next.ServeHTTP(w, r)
				return
			}

			app.serverError(w, err)
			return
		}

		// Add key to context values.
		ctx := context.WithValue(r.Context(), contextKeyIsUserAuthenticated, true)

		// Call next with a new request that contains the next context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
