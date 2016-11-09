package main

import (
  "fmt"
  "net/http"
  "./templates"
)

func login(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, templates.HTMLTop(templates.Style("error")))
  fmt.Fprint(w, templates.LoginForm)
  fmt.Fprint(w, templates.HTMLBottom())
}
