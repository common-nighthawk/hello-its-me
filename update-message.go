package main

import(
  "./models"
  "net/http"
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

  if r.FormValue("expire") != "" {
    err = message.SetExpiresAt(db, currentUser)
  }

  if r.FormValue("archive") != "" {
    err = message.Archive(db, currentUser)
  }

  if err != nil {
    http.Error(w, "failed to update expires_at", 500)
  }

  http.Redirect(w, r, "/messages", http.StatusFound)
}

func findMessage(file string, messages []*models.Message) (*models.Message, bool) {
  for _, message := range messages {
    if message.File == file {
      return message, true
    }
  }
  return nil, false
}
