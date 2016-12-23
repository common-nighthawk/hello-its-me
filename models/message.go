package models

import("time")

type Message struct {
  SenderUsername string
  ReceiverUUID string
  File string
  ExpiresAt time.Time
  ExplodeAfter int
  CreatedAt time.Time
  Archived bool
}
