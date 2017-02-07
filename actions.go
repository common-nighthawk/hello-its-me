package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func actions(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  tArgs := templates.Args{StyleSheet: "actions"}

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteTextBanner(w, "Hello, " + currentUser.Username)
  fmt.Fprint(w, "<div class='view'><a href='/messages'>View My Messages</a></div>")
  fmt.Fprint(w, "<div class='send'><a href='/message_new'>Send A Message</a></div>")
  templateHTMLBottom.Execute(w, tArgs)
}
