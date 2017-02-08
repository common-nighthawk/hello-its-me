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
  fmt.Fprint(w, "<a class='view' href='/messages'>View My Messages</a>")
  fmt.Fprint(w, "<a class='send' href='/message_new'>Send A Message</a>")
  templateHTMLBottom.Execute(w, tArgs)
}
