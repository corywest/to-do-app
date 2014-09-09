package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	//	http.Handle("/static/", fileServer)

	http.HandleFunc("/", HandleIndex)

	http.ListenAndServe(":8080", nil)
}

func HandleIndex(res http.ResponseWriter, req *http.Request) {
	var templates *template.Template

	type Page struct {
		Title string
	}

	p := Page{
		Title: "home",
	}

	templates = template.Must(template.ParseFiles("templates/home.html", "templates/layout.html"))
	err := templates.ExecuteTemplate(res, "base", p)

	if err != nil {
		fmt.Println(err)
	}
}
