package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *viewData) {
	// Get template set from templates map.
	// The name is suffixed with the extension "go.html",
	// this way I don't have to pass the whole name with the extesion.
	ts, ok := app.templates[name+".go.html"]

	// Check if the template set exists.
	if !ok {
		app.serverError(w, fmt.Errorf("the template %q does not exist", name))
		return
	}

	// Executing a template could produce an error, so instead of writing
	// to the response body directly, first write to a buffer.
	buffer := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buffer, name+".go.html", app.addDefaults(data, r))

	// Check for any error while executing the template.
	if err != nil {
		app.serverError(w, err)
		return
	}

	// If everything is fine write the rendered template
	// to the response body.
	buffer.WriteTo(w)
}

// Includes default data values to viewData struct fields.
func (app *application) addDefaults(data *viewData, r *http.Request) *viewData {
	if data == nil {
		data = &viewData{}
	}

	data.Year = time.Now().Year()
	data.CSRFToken = nosurf.Token(r)
	return data
}

// The serverError helper writes an error message and stack trace to the errLogger output,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLogger.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Sends generic client-related error messages.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
