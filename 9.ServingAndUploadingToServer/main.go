package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tmpl *template.Template
var err error

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "uploadfile.html", nil)
		return
	}
	//whole req body is parsed, up to a total of maxMemory bytes of its file parts are stored in memory
	r.ParseMultipartForm(50)
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer file.Close()
	fmt.Println("file header filename :", fileHeader.Filename)
	fmt.Println("fileheader size :", fileHeader.Size)
	fmt.Println("file header :", fileHeader.Header)

	contentType := fileHeader.Header["Content-Type"][0]
	fmt.Println("Content Type:", contentType)
	var osFile *os.File

	if contentType == "image/jpeg" {
		osFile, err = os.CreateTemp("./public/images", "*.jpg")
	} else if contentType == "application/pdf" {
		osFile, err = os.CreateTemp("./public/PDFs", "*.pdf")
	} else if contentType == "text/javascript" {
		osFile, err = os.CreateTemp("./public/js", "*.js")
	}
	fmt.Println("error:", err)
	defer osFile.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// func (f *File) Write(b []byte) (n int, err error)

	osFile.Write(fileBytes)

	fmt.Fprintf(w, "Your File was Successfully Uploaded!\n")
}

func main() {

	tmpl, err = template.ParseGlob("./static/*.html")
	if err != nil {
		fmt.Println("error in parsing html :", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/uploadfile", uploadFileHandler)

	//serve files to server
	myDir := http.Dir("./public")
	myHandler := http.FileServer(myDir)
	http.Handle("/serveFile/", http.StripPrefix("/serveFile", myHandler))

	fmt.Println("starting server ....")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", "sachin!!")
}
