package main

import (
  "./templates"
  "fmt"
  "net/http"
  "github.com/common-nighthawk/go-figure"
)

func landing(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  fmt.Fprint(w, "<pre class=\"figlet\">")
  figure.Write(w, figure.NewFigure("Hello, It's Me", "puffy"))
  fmt.Fprint(w, "</pre>")
  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")
  fmt.Fprint(w, templates.HTMLBottom())
}
