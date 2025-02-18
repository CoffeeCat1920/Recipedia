package modals

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

type Session struct {
  UserId uuid.UUID 
  SessionId string
  Exp time.Time
}

func generateSessionToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}

func NewSession(UserID uuid.UUID) *Session {
  return &Session{
    UserId: UserID,
    SessionId: generateSessionToken(),
    Exp: time.Now().AddDate(0, 2, 0),
  }
}
