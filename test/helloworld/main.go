package main

//https://freshman.tech/web-development-with-go/

import (
	//	"fmt"
	"log"
	"net/http"
	"os"
	//	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.kslawtawar.page", 301)
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
