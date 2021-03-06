package main

import (
  "./models"
  "./templates"
  "net/http"
)

func profile(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  tArgs := templates.Args{StyleSheet: "users", Email: currentUser.Email, Timezone: currentUser.Timezone}
  template := findTemplate("profile-form")

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteTextBanner(w, "Hello, " + currentUser.Username)
  template.Execute(w, tArgs)
  templateHTMLTop.Execute(w, tArgs)
}
