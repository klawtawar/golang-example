package main
//https://freshman.tech/web-development-with-go/
import (
	"html/template"
	"net/http"
	"os"
	"fmt"
        "log"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
        log.Print("starting server...")
	//http.HandleFunc("/", handler)


	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}
	
        // Start HTTP server.
        log.Printf("listening on port %s", port)

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}



