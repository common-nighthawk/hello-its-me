package main

import(
  "./models"
  "./templates"
  "fmt"
  "net/http"
  "html/template"
  "time"
  "unicode/utf8"
)

func createSignup(w http.ResponseWriter, r *http.Request) {
  var err error

  err = r.ParseForm()
  username, password, confirmation, timezone := r.FormValue("username"), r.FormValue("password"), r.FormValue("confirmation"), r.FormValue("timezone")

  row, err := db.Exec("SELECT * FROM users WHERE username = $1", username)
  count, err := row.RowsAffected()

  userError, msg := false, ""
  if count > 0 {
    userError, msg = true, "Username is already taken"
  } else if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(password) < 3 {
    userError, msg = true, "Username and password must be greater than 2 characters"
  } else if password != confirmation {
    userError, msg = true, "Password and confirmation do not match"
  }

  if userError {
    tArgs := templates.Args{"error"}
    htmlTop, _ := template.ParseFiles("templates/html-top.html")
    htmlBottom, _ := template.ParseFiles("templates/html-bottom.html")
    template, _ := template.ParseFiles("templates/signup-form.html")

    htmlTop.Execute(w, tArgs)
    fmt.Fprint(w, templates.HTMLError(msg))
    template.Execute(w, nil)
    htmlBottom.Execute(w, nil)
    return
  }

  uuid, err := models.GenerateUUID()
  result, err := db.Exec("INSERT INTO users (username, password, uuid, timezone) VALUES($1, $2, $3, $4)", username, password, uuid, timezone)
  _, err = result.RowsAffected()

  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  expiration := time.Now().Add(365 * 24 * time.Hour)
  cookie := http.Cookie{Name: "user", Value: uuid, Expires: expiration}
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/actions", http.StatusFound)
}
