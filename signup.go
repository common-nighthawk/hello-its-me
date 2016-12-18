package main

import (
  "./templates"
  "html/template"
  "net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{StyleSheet: "left"}
  template, _ := template.ParseFiles("templates/signup-form.html")

  templateHTMLTop.Execute(w, tArgs)
  template.Execute(w, tArgs)
  templateHTMLTop.Execute(w, tArgs)
}
