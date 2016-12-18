package main

import (
  "./templates"
  "html/template"
  "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
  tArgs := templates.Args{"error"}
  htmlTop, _ := template.ParseFiles("templates/html-top.html")
  htmlBottom, _ := template.ParseFiles("templates/html-bottom.html")
  template, _ := template.ParseFiles("templates/login-form.html")

  htmlTop.Execute(w, tArgs)
  template.Execute(w, nil)
  htmlBottom.Execute(w, nil)
}
