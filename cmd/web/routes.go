package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	router := mux.NewRouter()
	router.Handle("/", dynamicMiddleware.ThenFunc(app.home))
	router.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm)).Methods(http.MethodGet)
	router.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet)).Methods(http.MethodPost)
	router.Handle("/snippet/{id:[0-9]+}", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods(http.MethodGet)

	// static assets, see zbrains QDT for Noah's implementation
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	// pass servemux as param to middleware. Fcn returns http.Handler
	// so no need to do anything. Alice package essentially looks like
	// this commented out code.
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(router)
}
