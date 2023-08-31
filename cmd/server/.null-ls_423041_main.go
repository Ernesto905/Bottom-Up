package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
  fmt.Println("Listening on port 8080")


  h1 := func(w http.ResponseWriter, r *http.Request) {

    io.WriteString(w, "I am a string\n")
  }

  http.HandleFunc("/", h1) 

  // Creates the web server
  http.ListenAndServe(":8000", nil) 
}
