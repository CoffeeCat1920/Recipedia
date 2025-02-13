package session

import (
	"shiro/internal/modals"
)

type Session struct {
  User *modals.User
}

func New(user *modals.User) *Session {
  return &Session{
    User: user,    
  }
}  
