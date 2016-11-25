package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func actions(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindCurrentUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, " + user.Username)
  fmt.Fprint(w, "<a href=\"/messages\">View My Message</a>")
  fmt.Fprint(w, "<a href=\"/message_new\">Send A Message</a>")
  fmt.Fprint(w, templates.HTMLBottom())
}
