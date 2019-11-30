package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
