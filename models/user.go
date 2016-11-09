package models

import(
  "database/sql"
  "net/http"
)

type User struct {
  Username string
  Password string
}

func FindUser(cookies []*http.Cookie, db *sql.DB) (user *User, found bool) {
  var username string
  for _, cookie := range cookies {
    if cookie.Name == "username" {
      username = cookie.Value
    }
  }

  row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
  user = new(User)
  err := row.Scan(&user.Username, &user.Password)

  if err != nil {
    return nil, false
  }
  return user, true
}
