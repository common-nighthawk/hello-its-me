package main

import(
  "./models"
  "./templates"
  "database/sql"
  "fmt"
  "net/http"
  "time"
)

func createLogin(w http.ResponseWriter, r *http.Request) {
  var err error

  err = r.ParseForm()
  username, password := r.FormValue("username"), r.FormValue("password")

  row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

  user := new(models.User)
  err = row.Scan(&user.Username, &user.Password, &user.UUID, &user.Timezone)

  userError, msg := false, ""
  if err == sql.ErrNoRows {
    userError, msg = true, "No account with that username"
  } else if password != user.Password {
    userError, msg = true, "Incorrect password"
  }

  if userError {
    fmt.Fprint(w, templates.HTMLTop(templates.Style("error")))
    fmt.Fprint(w, templates.HTMLError(msg))
    fmt.Fprint(w, templates.LoginForm)
    fmt.Fprint(w, templates.HTMLBottom())
    return
  }

  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  // TODO: make secure
  expiration := time.Now().Add(365 * 24 * time.Hour)
  cookie := http.Cookie{Name: "user", Value: user.UUID, Expires: expiration}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/actions", http.StatusFound)
}
