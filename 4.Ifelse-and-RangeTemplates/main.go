package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template
var err error

func main() {

	tmpl, err = template.ParseGlob("./static/*.html")
	if err != nil {
		// Handle the error properly if templates can't be loaded
		fmt.Println("Error parsing templates:", err)
		return
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/todo", todoHandler)

	fmt.Println("starting Server")
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("In index handler")

	err := tmpl.ExecuteTemplate(w, "index.html", "Sachin!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("In List Handler")

	list := []string{"lauki", "karela", "baigan", "turai", "beans", "kaddu", "muli"}
	tmpl.ExecuteTemplate(w, "list.html", list)
}

type Todo struct {
	Task      string
	Completed bool
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/todo" {
		http.NotFound(w, r)
		return
	}

	fmt.Println("In todo Handler")

	todo := []Todo{
		{"wake up at 5", true}, {"drink boil water", true}, {"exercise", true}, {"bath", true}, {"study", false},
		{"more study", false}, {"watch movie ", true}, {"study?????", false},
	}
	tmpl.ExecuteTemplate(w, "todo.html", todo)
}
