package main

import (
  "context"
  "bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
  openai "github.com/sashabaranov/go-openai"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
  tmp, err := template.ParseFiles("../../public/wip.html")
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
  tmp, err := template.ParseFiles("../../public/wip.html")
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
  response, err := initializeGPT("ethics")
  if err != nil {
    log.Printf("Error initializing GPT: %v", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  // Create a data structure to hold the template data
  data := struct {
    Branch string
    GPT string
  }{
    Branch: "Ethics",
    GPT: response, 
  }

  // Parse the template file
  tmp, err := template.ParseFiles("../../public/wringer.html")
  if err != nil {
    log.Printf("Error parsing template: %v", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  // Create a buffer to execute the template into
  buffer := new(bytes.Buffer)
  err = tmp.Execute(buffer, data)
  if err != nil {
    log.Printf("Error executing template: %v", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  // Write the buffer content to the ResponseWriter
  buffer.WriteTo(w)
}

func initializeGPT(branch string) (string, error) {
  openAIKey := os.Getenv("OPENAI_API_KEY")
  prompt := "Hello, you are an " + branch + " Philosopher" 
  client := openai.NewClient(openAIKey)
  resp, err := client.CreateChatCompletion(
    context.Background(),
    openai.ChatCompletionRequest{
      Model: openai.GPT3Dot5Turbo,
      Messages: []openai.ChatCompletionMessage{
        {
          Role:    openai.ChatMessageRoleUser,
          Content: prompt,
        },
      },
    },
  )

  if err != nil {
    fmt.Printf("ChatCompletion error: %v\n", err)
  }

  // fmt.Println(resp.Choices[0].Message.Content)
  return resp.Choices[0].Message.Content, nil
}

func main() {

  fmt.Println("Listening on port 8080")

  err := godotenv.Load("../../.env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }


  r := mux.NewRouter()

  r.HandleFunc("/", homePage)
  r.HandleFunc("/metaphysics", metaphysics)
  r.HandleFunc("/epistemology", epistemology)
  r.HandleFunc("/ethics", ethics)

  http.ListenAndServe(":8000", r)
}


