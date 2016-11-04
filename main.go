package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", splash)
  http.ListenAndServe(":8080", nil)
}


func splash(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, it's me")
}
