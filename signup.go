package main

import (
  "./templates"
  "fmt"
  "net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, templates.HTMLTop(templates.Style("error")))
  fmt.Fprint(w, templates.SignupForm)
  fmt.Fprint(w, templates.HTMLBottom())
}
