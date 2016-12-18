package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "html/template"
)

func newMessage(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }
  receiverUsername, explodeParam := r.FormValue("receiver_username"), r.FormValue("explode")
  receiverUser, found := models.FindUserFromUsername(db, receiverUsername)

  tArgs := templates.Args{StyleSheet: "centered"}

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteBanner(w, "Hello, " + currentUser.Username)
  if !found && receiverUsername != "" {
    tArgs.ErrorMsg = "There is no user with the username " + receiverUsername
    templateErrorMsg.Execute(w, tArgs)
  }

  if found {
    tArgs.ErrorMsg = "Uh oh. This page is interactive. Please either enable JavaScript or update your web browser."
    templateErrorMsg.Execute(w, tArgs)
    fmt.Fprintf(w, "<p id='message'>Record your message for %s:</p>", receiverUser.Username)
    fmt.Fprintf(w, "<button id='start' value='%s'>Start</button>", explodeParam)
    fmt.Fprint(w,  "<button id='stop'>Stop</button>")
    fmt.Fprint(w,  "<button id='dismiss'>Dismiss</button>")
    fmt.Fprintf(w, "<button id='send' value='%s'>Send to %s</button>", receiverUser.Username, receiverUser.Username)
    fmt.Fprintf(w, "<p id='rec'>recording</p>")
    fmt.Fprintf(w, "<p id='audio-holder'></p>")
    fmt.Fprint(w, templates.HTMLScript(templates.MsgScript()))
  } else {
    template, _ := template.ParseFiles("templates/find-user-form.html")
    template.Execute(w, tArgs)
  }

  templateHTMLBottom.Execute(w, tArgs)
}
