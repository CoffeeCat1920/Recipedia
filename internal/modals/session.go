package modals

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
  UserId uuid.UUID 
  UserName string
  SessionToken string
  expire time.Time
}
