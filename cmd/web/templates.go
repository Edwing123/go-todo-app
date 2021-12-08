package main

import (
	"html/template"
	"path/filepath"
)

// Create templates map from views found in "[dir]/views".
func createTemplatesMap(dir string) (map[string]*template.Template, error) {
	// Get views file paths.
	views, err := filepath.Glob(filepath.Join(dir, "views/*.go.html"))
	if err != nil {
		return nil, err
	}

	// Create map with a length of len(views).
	templates := make(map[string]*template.Template)

	// Iterable on the views files and construct a template set that inclues:
	// - the view
	// - components
	// - layouts
	for _, view := range views {
		name := filepath.Base(view)

		// Parse view file
		ts, err := template.New(name).ParseFiles(view)
		if err != nil {
			return nil, err

		}

		// Parse layouts files.
		ts, err = ts.ParseGlob(filepath.Join(dir, "layout/*.go.html"))
		if err != nil {
			return nil, err

		}

		// Parse components files.
		ts, err = ts.ParseGlob(filepath.Join(dir, "components/*.go.html"))
		if err != nil {
			return nil, err

		}

		// Add template set to the map.
		templates[name] = ts
	}

	return templates, nil
}
