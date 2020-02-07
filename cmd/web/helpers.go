package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// serverError writes err msg and stack trace to errorLog,
// sends generic 500 error res to user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace) // report filename and line # one step back in stack trace; frame depth of 2

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends specific err code and description to user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper around clientError to send a 404
// not found res to user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()

	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// grab template set from cache

	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	//init buffer
	buf := new(bytes.Buffer)

	// write template to buffer. if error, throw user to
	// an error page and return. pass default data to template
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	// since no error, send buffered bytes to the response
	// writer.
	buf.WriteTo(w)
}
