package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "github.com/common-nighthawk/go-figure"
)

func newMessage(w http.ResponseWriter, r *http.Request) {
  user, found := models.FindCurrentUser(r.Cookies(), db)
  //TODO: handle not found user gracefully
  if found == false { panic("user not found") }

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  fmt.Fprint(w, "<pre class=\"figlet\">")
  figure.Write(w, figure.NewFigure(fmt.Sprintf("Hello, %s", user.Username), "puffy"))
  fmt.Fprint(w, "</pre>")

  receiverUsername := r.FormValue("username")
  toUser, found := models.FindUserFromUsername(db, receiverUsername)

  if !found && len(receiverUsername) > 0 {
    fmt.Fprintf(w, "Sorry, there is no user with the username %s", receiverUsername)
  }

  if found {
    fmt.Fprintf(w, "<button id='start' value='%s'>Start</button>", toUser.Username)
    fmt.Fprintf(w, "<button id='stop' value='%s'>Stop</button>", toUser.Username)
    fmt.Fprint(w, templates.HTMLScript(templates.Script()))
  } else {
    form := `<form action="/message_new" method="GET">
      <label for="username">Username:</label>
      <input type="text" name="username"><br/ >
      <input type="submit" value="Find User">
    </form>`
    fmt.Fprint(w, form)
  }

  fmt.Fprint(w, templates.HTMLBottom())
}
