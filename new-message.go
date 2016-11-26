package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
)

func newMessage(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }
  receiverUsername := r.FormValue("receiver_username")
  receiverUser, found := models.FindUserFromUsername(db, receiverUsername)

  fmt.Fprint(w, templates.HTMLTop(templates.Style("centered")))
  templates.WriteBanner(w, "Hello, " + currentUser.Username)
  if !found && receiverUsername != "" {
    msg := "There is no user with the username " + receiverUsername
    fmt.Fprint(w, templates.HTMLError(msg))
  }

  if found {
    fmt.Fprintf(w, "<button id='start' value='%s'>Start</button>", receiverUser.Username)
    fmt.Fprintf(w, "<button id='stop' value='%s'>Stop</button>", receiverUser.Username)
    fmt.Fprint(w, templates.HTMLScript(templates.Script()))
  } else {
    fmt.Fprint(w, templates.FindUserForm)
  }

  fmt.Fprint(w, templates.HTMLBottom())
}
