package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	templates := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Some dashboard"))
	})

	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		user := User{"John Cena", "john.cena@email.com", 55}

		templates.ExecuteTemplate(w, "index.html", user)
	})

	fmt.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
