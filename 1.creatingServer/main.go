package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error while starting server :%v", err)
	}
	fmt.Println("Server started")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("within hello handler")
	fmt.Fprintf(w, "heelllo ho gya ji paaji!! , path is %v ", r.URL.Path)
}
