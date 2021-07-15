package main
//https://freshman.tech/web-development-with-go/
import (
	"html/template"
	"net/http"
	"os"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}



