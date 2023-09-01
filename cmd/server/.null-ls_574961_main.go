package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
  fmt.Println("Listening on port 8080")


  h1 := func(w http.ResponseWriter, _ *http.Request) {
    tmpl := template.Must(template.ParseFiles("../../public/index.html"))
    tmpl.Execute(w, nil)
  }

  http.HandleFunc("/", h1) 

  // Creates the web server
  http.ListenAndServe(":8000", nil) 
}
