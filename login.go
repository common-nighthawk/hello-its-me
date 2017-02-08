package main

import (
  "./templates"
  "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{StyleSheet: "users"}
  template := findTemplate("login-form")

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteTextBanner(w, "Hello, It's Me")
  template.Execute(w, tArgs)
  templateHTMLBottom.Execute(w, tArgs)
}
