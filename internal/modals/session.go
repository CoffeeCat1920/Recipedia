package modals

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Session struct {
  SessionId string `json:"sessionid"`
  OwnerId string `json:"ownerid"` 
  Exp time.Time `json:"exp"`
}

func generateSessionToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}

func NewSession(OwnerId string) *Session {
  return &Session{
    SessionId: generateSessionToken(),
    OwnerId: OwnerId,
    Exp: time.Now().AddDate(0, 2, 0), 
  }
}

func (s *Session)IsExpired() bool {
  return s.Exp.Before(time.Now())
}
