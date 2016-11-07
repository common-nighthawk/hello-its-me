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

  fmt.Fprint(w, "<pre class=\"figlet\">")
  welcome := figure.NewFigure("Hello, It's Me", "puffy")
  figure.Write(w, welcome)
  fmt.Fprint(w, "</pre>")

  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")

  fmt.Fprint(w, pageBottom)
}
