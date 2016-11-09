package main

import (
  "fmt"
  "net/http"
  "./templates"
)

func signup(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, templates.HTMLTop(templates.Style("error")))
  fmt.Fprint(w, templates.SignupForm)
  fmt.Fprint(w, templates.HTMLBottom())
}
