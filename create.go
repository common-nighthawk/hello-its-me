package main

import(
  "fmt"
  "net/http"
  "unicode/utf8"
)

func create(w http.ResponseWriter, r *http.Request) {
  var err error

  err = r.ParseForm()
  username, password, confirmation := r.FormValue("username"), r.FormValue("password"), r.FormValue("confirmation")

  row, err := db.Exec("SELECT * FROM users WHERE username = $1", username)
  count, err := row.RowsAffected()

  userError, msg := false, ""
  if count > 0 {
    userError, msg = true, "Username is already taken"
  } else if utf8.RuneCountInString(username) < 3 && utf8.RuneCountInString(password) < 3 {
    userError, msg = true, "Username and password must be greater than 2 characters"
  } else if password != confirmation {
    userError, msg = true, "Password and confirmation do not match"
  }

  if userError {
    fmt.Fprint(w, signupTop)
    fmt.Fprintf(w, errorMsg, msg)
    fmt.Fprint(w, signupForm, pageBottom)
    return
  }

  result, err := db.Exec("INSERT INTO users (username, password) VALUES($1, $2)", username, password)
  _, err = result.RowsAffected()

  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  http.Redirect(w, r, "/", http.StatusFound)
}
