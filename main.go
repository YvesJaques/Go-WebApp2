package main

import (
	"log"
	"net/http"
	"text/template"
)

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
	log.Println("Server is running on port localhost:8080")
}
