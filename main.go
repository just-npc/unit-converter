package main

import (
	"fmt"
	"html/template"
	"net/http"
	"unit-converter/handlers"
)

var templateWeb = template.Must(template.ParseFiles("templates/length-converter.html"))


func main() {
	fmt.Print("Hello, Golang")
	http.HandleFunc("/length-convert", handlers.LengthHandler(templateWeb) ) 
	// http.HandleFunc("/length-converter", handlers.LengthHandler(templateWeb))
	http.ListenAndServe(":8080", nil)
}
