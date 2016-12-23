package models

import(
  "database/sql"
  "time"
)

type Message struct {
  SenderUsername string
  ReceiverUUID string
  File string
  ExpiresAt time.Time
  ExplodeAfter int
  CreatedAt time.Time
  Archived bool
}


func (message Message) SetExpiresAt(db *sql.DB, user *User) (err error) {
  if message.ExpiresAt.Before(time.Now().UTC()) {
    expiresAt := time.Now().UTC().Add(time.Duration(message.ExplodeAfter) * time.Second)
    dbStatement, _ := db.Prepare("UPDATE messages SET expires_at=$1 WHERE receiver_uuid=$2 AND file=$3")
    _, err = dbStatement.Exec(expiresAt, user.UUID, message.File)
  }
  return err
}

func (message Message) Archive(db *sql.DB, user *User) (err error) {
  var dbStatement *sql.Stmt
  if message.Archived {
    dbStatement, _ = db.Prepare("UPDATE messages SET archived=false WHERE receiver_uuid=$1 AND file=$2")
  } else {
    dbStatement, _ = db.Prepare("UPDATE messages SET archived=true WHERE receiver_uuid=$1 AND file=$2")
  }
  _, err = dbStatement.Exec(user.UUID, message.File)
  return err
}
