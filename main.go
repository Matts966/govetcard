package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Matts966/govetcard/api"
)

func resultHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/result.html"))
	result := api.GetZIP(r.FormValue("name"), r.FormValue("version"))
	if err := t.ExecuteTemplate(w, "result.html", result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":8080", nil)
}
