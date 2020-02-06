package main

import (
	"html/template"
	"path/filepath"

	"github.com/joshuaherrera/snippetbox/pkg/models"
)

// struct will hold dynamic data we want to render with templates.
// do this since we can only dynamically render one thing at a
// time.
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// grab all files with .page.tmpl extension
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract whole file name and assign
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// use Glob to add any layout files to the
		// template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// do same with partials
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))

		// add template set to cache
		cache[name] = ts
	}
	return cache, nil
}
