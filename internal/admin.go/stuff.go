package admin

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (admin *Admin) CheckPassword(password string) bool {
	hashedPassword := []byte(admin.Password)
	passwordBytes := []byte(password)
	
	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)

  fmt.Printf("\nOp: %s", password)
  p, _ := hashPassword(password)
  fmt.Printf("\nOp: %s", p) 
  fmt.Printf("\nHp: %s", hashedPassword)

	return err == nil
}
