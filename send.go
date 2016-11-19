package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "github.com/common-nighthawk/go-figure"
)

func send(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  fmt.Fprint(w, "<pre class=\"figlet\">")
  figure.Write(w, figure.NewFigure(fmt.Sprintf("Hello, %s", user.Username), "puffy"))
  fmt.Fprint(w, "</pre>")

  fmt.Fprint(w, "<button id='start'>Start</button>")
  fmt.Fprint(w, "<button id='stop'>Stop</button>")

  fmt.Fprint(w, templates.HTMLScript(templates.Script()))
  fmt.Fprint(w, templates.HTMLBottom())
}
