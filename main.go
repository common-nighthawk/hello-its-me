package main

import (
  "figure"
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", landing)
  http.ListenAndServe(":8080", nil)
}


func landing(w http.ResponseWriter, r *http.Request) {
  welcome := figure.NewFigure("Hello, It's Me", "puffy")
  for _, row := range welcome.Rowify() {
    fmt.Fprintf(w, "%v\n", row)
  }
}
