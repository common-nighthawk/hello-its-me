package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "github.com/common-nighthawk/go-figure"
)

func messages(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindCurrentUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  fmt.Fprint(w, "<pre class=\"figlet\">")
  figure.Write(w, figure.NewFigure(fmt.Sprintf("Hello, %s", user.Username), "puffy"))
  fmt.Fprint(w, "</pre>")

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
