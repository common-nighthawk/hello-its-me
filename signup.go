package main

import (
  "./templates"
  "net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{StyleSheet: "left"}
  template := findTemplate("signup-form")

  templateHTMLTop.Execute(w, tArgs)
  template.Execute(w, tArgs)
  templateHTMLTop.Execute(w, tArgs)
}
