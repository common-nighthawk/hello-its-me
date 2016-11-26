package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func messages(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, " + currentUser.Username)

  messages, err := currentUser.Messages(db)
  if err != nil {
    http.Error(w, "failed finding messages of current user", 500)
    return
  }
  fmt.Fprintf(w, "You currently have %d messages. <br><br>", len(messages))
  for _, message  := range messages {
    fmt.Fprint(w, "From: ", message.SenderUsername, "<br>")
    fmt.Fprint(w, templates.AudioPlayer(currentUser, message), "<br>")
  }

  fmt.Fprint(w, templates.HTMLBottom())
}
