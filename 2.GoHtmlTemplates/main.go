package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templ *template.Template

func main() {
	fmt.Println("********Parsing Html Templates***********")

	//templ, _ = template.ParseFiles("./static/index.html")  // for parsing single file
	templ, _ = template.ParseGlob("./static/*.html") // for parsing mutilple file
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("sarting server")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In index handler")
	//templ.Execute(w, nil)  // when loading single file
	templ.ExecuteTemplate(w, "index.html", nil) // when loading multiple files
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In about handler")
	templ.ExecuteTemplate(w, "about.html", nil)
}
