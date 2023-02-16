// / greeting    greeting
// Welcome to my page this is my main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"
)

// create a new type
type application struct {
}

func main() {
	// Create a flag for specifiying the port number \
	// when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	flag.Parse()
	// create an instance of the application type
	app := &application{}

	// create customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Start server on port %s", *addr)

	err := srv.ListenAndServe()
	log.Fatal(err) //should not reach here
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //always to this
	defer cancel()                                                          // then this to clean up
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
