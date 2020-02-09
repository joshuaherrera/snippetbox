package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// use for all route declarations
	router := mux.NewRouter()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/snippet/create", app.createSnippetForm).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", app.createSnippet).Methods(http.MethodPost)
	router.HandleFunc("/snippet/{id:[0-9]+}", app.showSnippet).Methods(http.MethodGet)

	// static assets, see zbrains QDT for Noah's implementation
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	// pass servemux as param to middleware. Fcn returns http.Handler
	// so no need to do anything. Alice package essentially looks like
	// this commented out code.
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(router)
}
