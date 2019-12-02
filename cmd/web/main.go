package main

import (
	"database/sql"
	"flag"
	"github.com/joshuaherrera/snippetbox/pkg/models/mysql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Must use cli flag and replace pass with password to connect to db
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	//use flag.Parse to parse the cli flag. need to do this b4
	//attempting to use addr or else it will use default variable
	flag.Parse()

	// create new loggers for writing info msgs. takes 3 params:
	// destination to write logs to, string prefix to msg, and flags
	// to indicate additional info to include. flags joined with
	// bitwise OR operator.
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// init new instance of application containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	// init new http.Server struct to use custom error logger.
	// set to use same network address and routes as b4
	// with http.ListenAndServe(*addr, mux)
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// due to lazy connection establishment, must ping db ti
	// start a connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// stopped at handlers for db, page 102
