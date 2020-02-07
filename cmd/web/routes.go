package main

import "net/http"

func (app *application) routes() http.Handler {
	// use for all route declarations
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// pass servemux as param to middleware. Fcn returns http.Handler
	// so no need to do anything
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
