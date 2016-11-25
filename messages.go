package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func messages(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindCurrentUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, " + user.Username)

  messages, err := user.Messages(db)
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  fmt.Fprintf(w, "You currently have %d messages. <br><br>", len(messages))
  for _, message  := range messages {
    fmt.Fprint(w, "From: ", message.SenderUsername, "<br>")
    fmt.Fprint(w, templates.AudioPlayer(user, message), "<br>")
  }

  fmt.Fprint(w, templates.HTMLBottom())
}
