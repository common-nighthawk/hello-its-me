package main

import(
  "database/sql"
  "fmt"
  "net/http"
  "time"
)

func session(w http.ResponseWriter, r *http.Request) {
  var err error

  err = r.ParseForm()
  username, password := r.FormValue("username"), r.FormValue("password")

  row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

  user := new(User)
  err = row.Scan(&user.username, &user.password)

  userError, msg := false, ""
  if err == sql.ErrNoRows {
    userError, msg = true, "No account with that username"
  } else if password != user.password {
    userError, msg = true, "Incorrect password"
  }

  if userError {
    fmt.Fprint(w, loginTop)
    fmt.Fprintf(w, errorMsg, msg)
    fmt.Fprint(w, loginForm, pageBottom)
    return
  }

  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  // TODO: make secure
  expiration := time.Now().Add(365 * 24 * time.Hour)
  cookie := http.Cookie{Name: "username", Value: user.username, Expires: expiration}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/", http.StatusFound)
}
