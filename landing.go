package main

import (
  "./models"
  "./templates"
  "fmt"
  "html/template"
  "net/http"
)

func landing(w http.ResponseWriter, r *http.Request) {
  if _, found := models.FindCurrentUser(r.Cookies(), db); found {
    http.Redirect(w, r, "/actions", http.StatusFound)
    return
  }

  tArgs := templates.Args{"centered"}
  htmlTop, _ := template.ParseFiles("templates/html-top.html")
  htmlBottom, _ := template.ParseFiles("templates/html-bottom.html")

  htmlTop.Execute(w, tArgs)
  templates.WriteBanner(w, "Hello, It's Me")
  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")
  htmlBottom.Execute(w, nil)
}
