package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// define command line flag for port # with name addr and
	// short explanation text. flag.String returns a pointer
	addr := flag.String("addr", ":4000", "HTTP network address")

	//use flag.Parse to parse the cli flag. need to do this b4
	//attempting to use addr or else it will use default variable
	flag.Parse()

	// create new loggers for writing info msgs. takes 3 params:
	// destination to write logs to, string prefix to msg, and flags
	// to indicate additional info to include. flags joined with
	// bitwise OR operator.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// make error log to write to stderr and log.Lshortfile flag
	// to include relevant file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	// create file server to serve static directory
	//path is relative to proj root
	//sanitizes all input with filepath.Clean() automatically
	//to avoid directory traversal attacks.
	//Can serve a single file with http.ServeFile(w,r, {file})
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//register fs as handler for all url paths that start with
	// /static using mux.Handle. strip /static prefix b4 req
	//reaches fs
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	// must dereference addr pointer
	infoLog.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
	//STOPPED ON 67
}
