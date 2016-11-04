package main

import (
  "figure"
  "fmt"
  // "html/template"
  "net/http"
)

const (
    pageTop =
    `<!DOCTYPE HTML>
      <html>
        <head>
          <style>
            body { text-align: center; }
            a { margin: 14px; }
          </style>
        </head>
        <title>Hello, It's Me</title>
        <body>`
    pageBottom =
        `</body>
      </html>`
)

func main() {
  http.HandleFunc("/", landing)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/login", login)
  http.ListenAndServe(":8080", nil)
}


func landing(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, pageTop)

  welcome := figure.NewFigure("Hello, It's Me", "puffy")
  fmt.Fprint(w, "<pre class=\"figlet\">")
  for _, row := range welcome.Rowify() {
    fmt.Fprintf(w, "%v\n", row)
  }
  fmt.Fprint(w, "</pre>")

  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")

  fmt.Fprint(w, pageBottom)
}

func signup(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "foo")
}

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "foo")
}
