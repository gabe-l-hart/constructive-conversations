package main

import (
	"flag"
	"github.com/boltdb/bolt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Helpers /////////////////////////////////////////////////////////////////////

type ModelServerContext struct {
	DB *bolt.DB
}

// Main ////////////////////////////////////////////////////////////////////////
func main() {

	// Command line args //
	port := flag.String(
		"port",
		"54321",
		"port for incoming traffic",
	)

	dbFilename := flag.String(
		"db-file",
		"model.db",
		"Filename for backend database storage",
	)

	logFilename := flag.String(
		"log-file",
		"model.log",
		"Filename for backend log output",
	)

	flag.Parse()

	// Logging //
	{
		handle, err := os.OpenFile(
			*logFilename,
			os.O_RDWR|os.O_CREATE|os.O_APPEND,
			0600)
		defer handle.Close()
		if nil != err {
			log.Fatalf("Failed to open log file: %s", *logFilename)
		} else {
			log.SetOutput(io.MultiWriter(os.Stdout, handle))
		}
	}

	// Database //
	context := ModelServerContext{}
	{
		db, err := bolt.Open(*dbFilename, 0600, &bolt.Options{Timeout: 5 * time.Second})
		defer db.Close()
		if nil != err {
			log.Fatalf("Failed to initialize database %s", *dbFilename)
		} else {
			context.DB = db
		}
	}

	// Server //
	log.Println("Starting the Model Server!")
	http.HandleFunc("/", EndpointRouter(context))
	http.ListenAndServe(":"+*port, nil)
}
