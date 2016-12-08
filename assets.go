package main

import(
  "./models"
  "net/http"
  "strings"
)

func assets(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }
  var uuid string
  if urlSplit := strings.Split(r.URL.Path, "/"); len(urlSplit) > 3 {
    uuid = urlSplit[3]
  }
  if currentUser.UUID != uuid {
    http.Error(w, "unauthorized", 500)
    return
  }
  file := messagesDir + strings.SplitAfter(r.URL.Path, "/messages")[1]
  http.ServeFile(w, r, file)
}
