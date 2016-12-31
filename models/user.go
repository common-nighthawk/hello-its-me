package models

import(
  "../secrets"
  "crypto/rand"
  "encoding/base64"
  "database/sql"
  "fmt"
  "net/http"
  "net/smtp"
  "time"
)

type User struct {
  Username string
  Password string
  UUID string
  Timezone string
  Email string
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
  err := row.Scan(&user.Username, &user.Password, &user.UUID, &user.Timezone, &user.Email)
  if err != nil {
    return nil, false
  }
  return user, true
}

func (user User) Messages(db *sql.DB, status string) (messages []*Message, err error) {
  rows, err := db.Query(messageQuery(status), user.UUID)
  for rows.Next() {
    message := new(Message)
    err = rows.Scan(&message.SenderUsername, &message.ReceiverUUID, &message.File, &message.ExpiresAt,
                    &message.ExplodeAfter, &message.CreatedAt, &message.Archived)
    messages = append(messages, message)
  }
  return messages, err
}

func messageQuery(status string) string {
  switch status {
  case "all":
    return "SELECT * FROM messages WHERE receiver_uuid = $1 ORDER BY created_at DESC"
  case "active":
    return "SELECT * FROM messages WHERE receiver_uuid = $1 AND archived = false ORDER BY created_at DESC"
  case "archived":
    return "SELECT * FROM messages WHERE receiver_uuid = $1 AND archived = true ORDER BY created_at DESC"
  }
  panic("undefined archive criteria for messages")
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

func (user User) Notify() error {
  auth := smtp.PlainAuth("", secrets.Email, secrets.SMTPPassword, "smtp.gmail.com")
  subject := "New Message on 'Hello, Its Me'"
  text := "You have a new message! Check it out at helloitsme.site"
  message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", secrets.Email, user.Email, subject, text)
  return smtp.SendMail("smtp.gmail.com:587", auth, "helloitsmewebsite@gmail.com", []string{user.Email}, []byte(message))
}
