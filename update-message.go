package main

import(
  "./models"
  "net/http"
  "time"
)

func updateMessage(w http.ResponseWriter, r *http.Request) {
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

  if message.ExpiresAt.Before(time.Now().UTC()) {
    expiresAt := time.Now().UTC().Add(time.Duration(message.ExplodeAfter) * time.Second)
    dbStatement, _ := db.Prepare("UPDATE messages SET expires_at=$1 WHERE receiver_uuid=$2 AND file=$3")
    _, err = dbStatement.Exec(expiresAt, currentUser.UUID, message.File)
  }

  if err != nil {
    http.Error(w, "failed to update expires_at", 500)
  }
}

func findMessage(file string, messages []*models.Message) (*models.Message, bool) {
  for _, message := range messages {
    if message.File == file {
      return message, true
    }
  }
  return nil, false
}
