package main

import (
  "net/http"
  "fmt"
  "figure"
)

const (
    landingTop =
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
)

func landing(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, landingTop)

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
