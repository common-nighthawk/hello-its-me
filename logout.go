package main

import (
  "net/http"
)

func logout(w http.ResponseWriter, r *http.Request) {
  cookie := http.Cookie{Name: "user", Value: ""}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/", http.StatusFound)
}
