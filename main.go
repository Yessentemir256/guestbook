package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html") // take a file, create a template and parse the content from this file
	check(err)
	err = html.Execute(writer, nil) // write in writer the executed template
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
