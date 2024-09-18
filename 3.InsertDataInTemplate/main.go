package main

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
	prod1 = product{
		ProdID: 15,
		Name:   "Wicked Cool Phone",
		Cost:   899,
		Specs: prodSpec{
			Size:   "150 x 70 x 7 mm",
			Weight: 65,
			Descr:  "Over priced shiny thing designed to shatter impact",
		},
	}
*/

type prodSpec struct {
	Size   float32
	Weight int
	Descr  string
}

type ProductDetails struct {
	ProdID int
	Name   string
	Cost   float32
	Specs  prodSpec
}

var tmpl *template.Template
var name = "sachin"

func main() {

	tmpl, _ = template.ParseGlob("./static/*html")

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/about", abouthandler)
	http.HandleFunc("/productInfo1", prodInfoHandler)
	http.HandleFunc("/productStructInfo", prodInfoStructHandler)
	fmt.Println("starting server at port 8080")
	http.ListenAndServe(":8080", nil)

}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in Index-handler")
	tmpl.ExecuteTemplate(w, "index.html", name)
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in About-handler")
	tmpl.ExecuteTemplate(w, "about.html", name)
}

func prodInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Product info Non structured")
	myProduct := ProductDetails{
		ProdID: 69,
		Name:   "Jhonny's toy",
		Cost:   69.69,
		Specs: prodSpec{
			Size:   18.56,
			Weight: 10,
			Descr:  "Toy to be handled very carefully, is delicate and hard at the same time ;)",
		},
	}
	tmpl.ExecuteTemplate(w, "productinfo.html", myProduct)
}

func prodInfoStructHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Product info - structured")
	myProduct := ProductDetails{
		ProdID: 69,
		Name:   "Jhonny's toy",
		Cost:   69.69,
		Specs: prodSpec{
			Size:   18.56,
			Weight: 10,
			Descr:  "Toy to be handled very carefully, is delicate and hard at the same time ;)",
		},
	}
	tmpl.ExecuteTemplate(w, "productStructInfo.html", myProduct)

}
