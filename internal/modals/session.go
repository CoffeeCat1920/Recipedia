package modals

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Session struct {
  UserId string 
  SessionId string
  Exp time.Time
}

func generateSessionToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}

func NewSession(UserID string) *Session {
  return &Session{
    UserId: UserID,
    SessionId: generateSessionToken(),
    Exp: time.Now().AddDate(0, 2, 0),
  }
}

func (s *Session)IsExpired() bool {
  return s.Exp.Before(time.Now())
}
