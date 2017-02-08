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

  tArgs := templates.Args{StyleSheet: "landing"}

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteTextBanner(w, "Hello, It's Me")
  fmt.Fprint(w, "<a class='log-in' href='/login'>Log In</a>")
  fmt.Fprint(w, "<a class='sign-up' href='/signup'>Sign Up</a>")
  templateHTMLBottom.Execute(w, tArgs)
}
