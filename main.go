package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Make it so this url path only renders for '/',
	// otherwise, return a 404 response to the client
	// Must return or else rest of fcn executes.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// check to see if our req is not a POST req
	// if not, return a 405 error and return, else
	// continue with creation logic.
	// NOTE: can only call w.WriteHeader once per res.
	//       Also, if we don't use the method to send a status code
	//		 w.Write will automatically send a 200 OK status code.
	if r.Method != "POST" {
		w.WriteHeader((405))
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet... "))
}

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
