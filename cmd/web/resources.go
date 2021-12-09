package main

import (
	"net/http"
	"strings"
)

// This is a special middleware that will be passed to the http.StripPrefix
// function, the idea is to prevent access to paths like "/css/", that's because
// I don't want people to access the directory-like view provided by the file server
// handler provided by the http package.
func (app *application) serveStaticFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the URL path ends with a "/", if that's the case,
		// send a not found response.
		if strings.HasSuffix(r.URL.Path, "/") {
			app.notFound(w, r)
			return
		}

		// The next handler in the chain is a http.FileServer.
		next.ServeHTTP(w, r)
	})
}
