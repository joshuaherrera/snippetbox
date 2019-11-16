package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to init a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// the servemux stores mapping between url patterns and correspongind
	// handlers. usually only use one per app for all routes
	// NOTE: could also just use http.HandleFunc which uses a
	// default servmux, but the default is a global var and using it
	// is a sec risk.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal((err))
}
