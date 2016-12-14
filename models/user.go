package models

import(
  "crypto/rand"
  "encoding/base64"
  "database/sql"
  "net/http"
  "time"
)

type User struct {
  Username string
  Password string
  UUID string
  Timezone string
}

func FindCurrentUser(cookies []*http.Cookie, db *sql.DB) (*User, bool) {
  var uuid string
  for _, cookie := range cookies {
    if cookie.Name == "user" {
      uuid = cookie.Value
    }
  }
  return FindUserFromUUID(db, uuid)
}

func FindUserFromUsername(db *sql.DB, username string) (user *User, found bool) {
  row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
  return findUserFromRow(row)
}

func FindUserFromUUID(db *sql.DB, uuid string) (*User, bool) {
  row := db.QueryRow("SELECT * FROM users WHERE uuid = $1", uuid)
  return findUserFromRow(row)
}

func findUserFromRow(row *sql.Row) (*User, bool) {
  user := new(User)
  err := row.Scan(&user.Username, &user.Password, &user.UUID, &user.Timezone)
  if err != nil {
    return nil, false
  }
  return user, true
}

func (user User) Messages(db *sql.DB) (messages []*Message, err error) {
  rows, err := db.Query("SELECT * FROM messages WHERE receiver_uuid = $1 ORDER BY created_at DESC", user.UUID)
  for rows.Next() {
    message := new(Message)
    err = rows.Scan(&message.SenderUsername, &message.ReceiverUUID, &message.File, &message.ExpiresAt, &message.ExplodeAfter, &message.CreatedAt)
    messages = append(messages, message)
  }
  return messages, err
}

func (user User) TZLocation() *time.Location {
  location, err := time.LoadLocation(user.Timezone)
  if err != nil {
    location, _ = time.LoadLocation("UTC")
  }
  return location
}

func GenerateUUID() (string, error) {
  b := make([]byte, 32)
  _, err := rand.Read(b)
  if err != nil {
    return "", err
  }
  return base64.URLEncoding.EncodeToString(b), nil
}
