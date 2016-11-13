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

	config := flag.String(
		"config",
		"",
		"Configuration file",
	)

	flag.Parse()
	cfg := ParseConfig(*config)

	// Logging //
	{
		handle, err := os.OpenFile(
			cfg.LogFilename,
			os.O_RDWR|os.O_CREATE|os.O_APPEND,
			0600)
		defer handle.Close()
		if nil != err {
			log.Fatalf("Failed to open log file: %s", cfg.LogFilename)
		} else {
			log.SetOutput(io.MultiWriter(os.Stdout, handle))
		}
	}

	// Database //
	context := ModelServerContext{}
	{
		db, err := bolt.Open(cfg.DbFilename, 0600, &bolt.Options{Timeout: 5 * time.Second})
		defer db.Close()
		if nil != err {
			log.Fatalf("Failed to initialize database %s", cfg.DbFilename)
		} else {
			context.DB = db
		}
	}

	// Server //
	log.Println("Starting the Model Server!")
	http.HandleFunc("/", EndpointRouter(context))
	http.ListenAndServe(":"+cfg.Port, nil)
}
