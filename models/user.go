package models

import(
  "crypto/rand"
  "encoding/base64"
  "database/sql"
  "net/http"
)

type User struct {
  Username string
  Password string
  UUID string
}

func FindUser(cookies []*http.Cookie, db *sql.DB) (user *User, found bool) {
  var uuid string
  for _, cookie := range cookies {
    if cookie.Name == "user" {
      uuid = cookie.Value
    }
  }

  row := db.QueryRow("SELECT * FROM users WHERE uuid = $1", uuid)
  user = new(User)
  err := row.Scan(&user.Username, &user.Password, &user.UUID)

  if err != nil {
    return nil, false
  }
  return user, true
}

func (user User) Messages(db *sql.DB) (messages []*Message, err error) {
  rows, err := db.Query("SELECT * FROM messages WHERE receiver_uuid = $1", user.UUID)
  for rows.Next() {
    message := new(Message)
    err = rows.Scan(&message.SenderUsername, &message.ReceiverUUID, &message.Path)
    messages = append(messages, message)
  }
  return messages, err
}

func GenerateUUID() (string, error) {
  b := make([]byte, 32)
  _, err := rand.Read(b)
  if err != nil {
    return "", err
  }
  return base64.URLEncoding.EncodeToString(b), nil
}
