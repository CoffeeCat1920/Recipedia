package modals

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

type Session struct {
  SessionId string `json:"sessionid"`
  OwnerId string `json:"ownerid"` 
  Exp string `json:"exp"`
}

func generateSessionToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}

func getTime(input string) (time.Time, error) {
  layout := "2006-01-02 15:04:05"  
  t, err := time.Parse(layout, input)
  if err != nil {
    fmt.Println("Error parsing time:", err)
    return time.Now(), err
  }
  return t, nil
}

func NewSession(OwnerId string) *Session {
  return &Session{
    SessionId: generateSessionToken(),
    OwnerId: OwnerId,
    Exp: time.Now().AddDate(0, 2, 0).Format("2006-01-02 15:04:05"), 
  }
}

func (s *Session)IsExpired() bool {
  t,err := getTime(s.Exp)
  if err != nil {
    return false
  }
  return t.Before(time.Now())
}

func (s *Session)GetExpTime() (time.Time, error) {
  t,err := getTime(s.Exp)
  if err != nil {
    return time.Now(), err 
  }
  return t, nil 
}
