package admin 

import (
  "golang.org/x/crypto/bcrypt"
)

type Admin struct { 
  Password string `json:"password"`
}

var (
  admin *Admin
  admin_password = "123"
) 

func hashPassword(password string) (string, error) {
  passwordBytes := []byte(password) 
  hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

  if err != nil {
    return "", err
  }

  return string(hashedPassword), nil 
}

func New() (*Admin) {
  if admin != nil {
    return admin 
  }

  hashedPassword,_ := hashPassword(admin_password)

  admin = &Admin{
    Password: hashedPassword,
  }
  return admin
}
