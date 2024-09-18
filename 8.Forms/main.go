package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template
var err error

func main() {
	tmpl, err = template.ParseGlob("./static/*html")
	if err != nil {
		log.Fatalf("error in parsing html file :%v", err)
	}

	http.HandleFunc("/getform", getformHandler)               //loads the form
	http.HandleFunc("/processGetform", processGetformHandler) // once submitted, this is called through form action

	http.HandleFunc("/postform", postformHandler)
	http.HandleFunc("/processPostform", processPostformHandler)

	fmt.Println("stating server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getformHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Getform.html", nil)
}

type Users struct {
	Name       string
	SecretData string
}

func processGetformHandler(w http.ResponseWriter, r *http.Request) {
	var user Users
	user.Name = r.FormValue("usernameName")
	user.SecretData = r.FormValue("dataName")

	fmt.Println("user name :", user.Name)
	fmt.Println("Secret data :", user.SecretData)

	tmpl.ExecuteTemplate(w, "thanks.html", user.Name)
}

func postformHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "postform.html", nil)
}

func processPostformHandler(w http.ResponseWriter, r *http.Request) {
	var user Users
	user.Name = r.FormValue("username")
	user.SecretData = r.FormValue("data")

	fmt.Println("user name :", user.Name)
	fmt.Println("Secret data :", user.SecretData)

	tmpl.ExecuteTemplate(w, "thanks.html", user.Name)
}
