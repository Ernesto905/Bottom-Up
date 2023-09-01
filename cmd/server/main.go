package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)


func homePage(w http.ResponseWriter, _ *http.Request) {
  tmpl := template.Must(template.ParseFiles("../../public/index.html"))
  tmpl.Execute(w, nil)
}

func wringer(w http.ResponseWriter, r *http.Request) {
  branch := r.URL.Query().Get("branch") 
  io.WriteString(w, fmt.Sprintf("You chose %s", branch))
}

func main() {
  fmt.Println("Listening on port 8080")

  http.HandleFunc("/", homePage) 

  http.HandleFunc("/wringer", wringer) 

  // Creates the web server
  http.ListenAndServe(":8000", nil) 
}
