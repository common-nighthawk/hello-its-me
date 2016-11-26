package main

import (
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func landing(w http.ResponseWriter, r *http.Request) {
  if _, found := models.FindCurrentUser(r.Cookies(), db); found {
    http.Redirect(w, r, "/actions", http.StatusFound)
    return
  }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, It's Me")
  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")
  fmt.Fprint(w, templates.HTMLBottom())
}
