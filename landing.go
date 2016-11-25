package main

import (
  "./templates"
  "fmt"
  "net/http"
)

func landing(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, It's Me")
  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")
  fmt.Fprint(w, templates.HTMLBottom())
}
