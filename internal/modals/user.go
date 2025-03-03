package modals

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct { 
  UUID string `json:"uuid"` 
  Name string `json:"name"`
  Password string `json:"password"`
}


// TODO - Change it to return an error type
func (user *User) CheckPassword(password string) (bool) {
	hashedPassword := []byte(user.Password)
	passwordBytes := []byte(password)
	
	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)

  fmt.Printf("\nOp: %s", password)
  p, _ := hashPassword(password)
  fmt.Printf("\nOp: %s", p) 
  fmt.Printf("\nHp: %s", hashedPassword)

	return err == nil
}

func NewUser(name string, password string) *User {
  hashedPassword, err := hashPassword(password)
  if err != nil {
    panic(err)
  }

  user := &User{
    UUID: uuid.NewString(),
    Name: name,
    Password: hashedPassword, 
  }

  return user
}
