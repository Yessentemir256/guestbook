package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html") // take a file, create a template and parse the content from this file
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook) // write in writer the executed template
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html") // take a file, create a template and parse the content from this file
	check(err)
	err = html.Execute(writer, nil) // write in writer the executed template
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
