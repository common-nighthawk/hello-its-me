package main

import(
  "fmt"
  "figure"
  "./models"
  "net/http"
  "./templates"
)

func actions(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  fmt.Fprint(w, "<pre class=\"figlet\">")
  figure.Write(w, figure.NewFigure(fmt.Sprintf("Hello, %s", user.Username), "puffy"))
  fmt.Fprint(w, "</pre>")
  fmt.Fprint(w, "<a href=\"/messages\">View My Message</a>")
  fmt.Fprint(w, "<a href=\"/send\">Send A Message</a>")
  fmt.Fprint(w, templates.HTMLBottom())
}
