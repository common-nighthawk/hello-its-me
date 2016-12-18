package main

import (
  "./templates"
  "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{StyleSheet: "left"}
  template := findTemplate("login-form")

  templateHTMLTop.Execute(w, tArgs)
  template.Execute(w, tArgs)
  templateHTMLBottom.Execute(w, tArgs)
}
