package main

import (
  "./templates"
  "html/template"
  "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{StyleSheet: "left"}
  template, _ := template.ParseFiles("templates/login-form.html")

  templateHTMLTop.Execute(w, tArgs)
  template.Execute(w, tArgs)
  templateHTMLBottom.Execute(w, tArgs)
}
