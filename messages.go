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

  tArgs := templates.Args{StyleSheet: "centered", Script: "message-expire", UUID: currentUser.UUID}

  templateHTMLTop.Execute(w, tArgs)
  templates.WriteBanner(w, "Hello, " + currentUser.Username)

  activeMessages, err := currentUser.Messages(db, "active")
  archivedMessages, err := currentUser.Messages(db, "archived")
  if err != nil {
    http.Error(w, "failed finding messages of current user", 500)
    return
  }
  fmt.Fprintf(w, "You currently have %d messages. <br><br>", len(activeMessages))

  for _, message  := range activeMessages {
    fmt.Fprint(w, "<div class='message-container'>")
    writeMessage(w, message, currentUser)
    fmt.Fprint(w, "<div class='message-opts'>")
    fmt.Fprint(w, "<ul>")
    fmt.Fprintf(w, "<li><a href='message_new?receiver_username=%s'>Reply</a></li>", message.SenderUsername)
    fmt.Fprintf(w, "<li><a href='message_update?archive=true&file=%s'>Archive</a></li>", message.File)
    fmt.Fprint(w, "<li><a href='destroy-message'>Delete</a></li>")
    fmt.Fprint(w, "</div></div>")
  }
  for i, message  := range archivedMessages {
    if i == 0 { fmt.Fprint(w, "<hr>") }
    fmt.Fprint(w, "<div class='message-container'>")
    writeMessage(w, message, currentUser)
    fmt.Fprint(w, "<div class='message-opts'>")
    fmt.Fprint(w, "<ul>")
    fmt.Fprintf(w, "<li><a href='message_update?archive=true&file=%s'>Unarchive</a></li>", message.File)
    fmt.Fprint(w, "<li><a href='destroy-message'>Delete</a></li>")
    fmt.Fprint(w, "</div></div>")
  }

  templateScript.Execute(w, tArgs)
  templateHTMLBottom.Execute(w, tArgs)
}

func writeMessage(w http.ResponseWriter, message *models.Message, currentUser *models.User) {
  fmt.Fprint(w, "<div class='message'>")
  fmt.Fprint(w, "From: ", message.SenderUsername, "<br>")

  tArgs := templates.Args{UUID: currentUser.UUID, File: message.File}
  template := findTemplate("audio-player")
  template.Execute(w, tArgs)

  fmt.Fprint(w, "<span>")
  fmt.Fprint(w, "Sent: ", message.CreatedAt.In(currentUser.TZLocation()).Format(msgTimeFmt), " | ")
  fmt.Fprint(w, "Explodes: ", explodesAt(message, currentUser.TZLocation()), "<br>")
  fmt.Fprint(w, "</span></div>")
}

func explodesAt(message *models.Message, location *time.Location) string {
  if message.ExpiresAt.After(time.Now().UTC()) {
    return message.ExpiresAt.In(location).Format(msgTimeFmt)
  }
  return fmt.Sprintf("%s after listening", displaySeconds(message.ExplodeAfter))
}

func displaySeconds(seconds int) (phrase string) {
  var times []map[string]interface{}
  times = append(times, map[string]interface{}{"unit": "year", "conv": 31536000})
  times = append(times, map[string]interface{}{"unit": "month", "conv": 2628000})
  times = append(times, map[string]interface{}{"unit": "week", "conv": 604800})
  times = append(times, map[string]interface{}{"unit": "day", "conv": 86400})
  times = append(times, map[string]interface{}{"unit": "hour", "conv": 3600})
  times = append(times, map[string]interface{}{"unit": "minute", "conv": 60})

  for _, time := range times {
    time["quan"] = seconds / time["conv"].(int)
    seconds -= time["quan"].(int) * time["conv"].(int)
    if time["quan"].(int) > 0 {
      phrase += fmt.Sprintf(" %d %s", time["quan"], time["unit"])
      if time["quan"].(int) > 1 { phrase += "s" }
    }
  }
  return phrase
}
