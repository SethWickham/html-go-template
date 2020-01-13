package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//HTMLPage for basic html template
type HTMLPage struct {
	Heading   string
	Paragraph string
}

func htmlPageHandler(w http.ResponseWriter, r *http.Request) {
	h := HTMLPage{Heading: "This is the Heading"}
	p := HTMLPage{Paragraph: "This is the paragraph"}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, h)
	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> A simple way to use html in go</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/htmltemplate", htmlPageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
