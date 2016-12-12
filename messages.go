package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "time"
)

const(
  msgTimeFmt = "January 2, 2006 @ 03:04 pm"
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
    fmt.Fprint(w, "<div class='message'>")
    fmt.Fprint(w, "From: ", message.SenderUsername, "<br>")
    fmt.Fprint(w, templates.AudioPlayer(currentUser, message), "<br>")
    fmt.Fprint(w, "<span>")
    fmt.Fprint(w, "Sent: ", message.CreatedAt.Format(msgTimeFmt), " | ")
    fmt.Fprint(w, "Explodes: ", explodesAt(message), "<br>")
    fmt.Fprint(w, "</span></div>")
  }

  fmt.Fprint(w, templates.HTMLScript(templates.ExpireScript()))
  fmt.Fprint(w, templates.HTMLBottom())
}

func explodesAt(message *models.Message) string {
  if message.ExpiresAt.After(time.Now().UTC()) {
    return message.ExpiresAt.Format(msgTimeFmt)
  }
  duration := time.Duration(message.ExplodeAfter) * time.Second
  return fmt.Sprintf("%s after listening", duration)
}
