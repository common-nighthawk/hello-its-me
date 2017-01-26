package main

import (
  "./models"
  "net/http"
)

func updateProfile(w http.ResponseWriter, r *http.Request) {
  currentUser, found := models.FindCurrentUser(r.Cookies(), db)
  if !found {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  err := r.ParseForm()
  email, timezone := r.FormValue("email"), r.FormValue("timezone")

  result, err := db.Exec("UPDATE users SET email=$1, timezone=$2 WHERE uuid=$3", email, timezone, currentUser.UUID)
  _, err = result.RowsAffected()

  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  http.Redirect(w, r, "/actions", http.StatusFound)
}
