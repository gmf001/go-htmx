package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Movie struct {
	Title    string
	Director string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		movies := map[string][]Movie{
			"Movies": {
				{Title: "The Matrix", Director: "Wachowski Brothers"},
				{Title: "Up", Director: "Pete Docter"},
				{Title: "Alien", Director: "Ridley Scott"},
			},
		}

		tmpl.Execute(w, movies)
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Movie{Title: title, Director: director})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
