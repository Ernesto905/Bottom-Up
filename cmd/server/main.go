package main

import (
	"fmt"
  "github.com/gorilla/mux"
  "net/http"
	"text/template"
)


func homePage(w http.ResponseWriter, _ *http.Request) {
  tmpl := template.Must(template.ParseFiles("../../public/index.html"))
  tmpl.Execute(w, nil)
}

func metaphysics(w http.ResponseWriter, r *http.Request) {
  // Create a data structure to hold the template data
  data := struct {
    Branch string
  }{
    Branch: "Metaphysics",
  }

  // Parse the template file
  tmp, err := template.ParseFiles("../../public/wringer.html")
  if err != nil {
    http.Error(w, "could not load template", http.StatusInternalServerError)
    return
  }

  // Execute the template and write to ResponseWriter
  err = tmp.Execute(w, data)
  if err != nil {
    http.Error(w, "could not execute template", http.StatusInternalServerError)
    return
  }
}

func epistemology(w http.ResponseWriter, r *http.Request) {
  // Create a data structure to hold the template data
  data := struct {
    Branch string
  }{
    Branch: "Epistemology",
  }

  // Parse the template file
  tmp, err := template.ParseFiles("../../public/wringer.html")
  if err != nil {
    http.Error(w, "could not load template", http.StatusInternalServerError)
    return
  }

  // Execute the template and write to ResponseWriter
  err = tmp.Execute(w, data)
  if err != nil {
    http.Error(w, "could not execute template", http.StatusInternalServerError)
    return
  }
}

func ethics(w http.ResponseWriter, r *http.Request) {
  // Create a data structure to hold the template data
  data := struct {
    Branch string
  }{
    Branch: "Ethics",
  }

  // Parse the template file
  tmp, err := template.ParseFiles("../../public/wringer.html")
  if err != nil {
    http.Error(w, "could not load template", http.StatusInternalServerError)
    return
  }

  // Execute the template and write to ResponseWriter
  err = tmp.Execute(w, data)
  if err != nil {
    http.Error(w, "could not execute template", http.StatusInternalServerError)
    return
  }
}

func wringer(w http.ResponseWriter, r *http.Request) {
}

func main() {
  fmt.Println("Listening on port 8080")

  r := mux.NewRouter()

  r.HandleFunc("/", homePage)
  r.HandleFunc("/metaphysics", metaphysics)
  r.HandleFunc("/epistemology", epistemology)
  r.HandleFunc("/ethics", ethics)

  http.ListenAndServe(":8000", r)
}
