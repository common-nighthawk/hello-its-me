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

  tArgs := templates.Args{StyleSheet: "centered"}
  templateHTMLTop.Execute(w, tArgs)
  templates.WriteBanner(w, "Hello, It's Me")
  fmt.Fprint(w, "<a href=\"/login\">Log In</a>")
  fmt.Fprint(w, "<a href=\"/signup\">Sign Up</a>")
  templateHTMLBottom.Execute(w, tArgs)
}
