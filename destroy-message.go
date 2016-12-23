package main

import(
  "./models"
  "net/http"
)

func destroyMessage(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Error(w, "no user logged in", 500)
    return
  }

  messages, err := currentUser.Messages(db, "all")
  if err != nil {
    http.Error(w, "failed to find messages for user", 500)
    return
  }

  message, found := findMessage(r.FormValue("file"), messages)
  if !found {
    http.Error(w, "no message with given filename", 500)
    return
  }

  dbStatement, _ := db.Prepare("DELETE FROM messages WHERE receiver_uuid=$1 AND file=$2")
  _, err = dbStatement.Exec(currentUser.UUID, message.File)

  if err != nil {
    http.Error(w, "failed to delete message", 500)
    return
  }

  http.Redirect(w, r, "/messages", http.StatusFound)
}
