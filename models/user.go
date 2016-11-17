package models

import(
  "database/sql"
  "net/http"
)

type User struct {
  Username string
  Password string
  UUID string
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
  err := row.Scan(&user.Username, &user.Password, &user.UUID)

  if err != nil {
    return nil, false
  }
  return user, true
}

func (user User) Messages(db *sql.DB) (messages []*Message, err error) {
  rows, err := db.Query("SELECT * FROM messages WHERE receiver_uuid = $1", user.Username)
  for rows.Next() {
    message := new(Message)
    err = rows.Scan(&message.SenderUUID, &message.ReceiverUUID, &message.Path)
    messages = append(messages, message)
  }
  return messages, err
}
