package api

import (
	"crypto/rand"
	"encoding/base64"

	"fmt"
	"net/http"
	"shiro/internal/database"
)

func generateSessionToken() string {
  b := make([]byte, 32)
  rand.Read(b)
  return base64.URLEncoding.EncodeToString(b)
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  userName := r.PostFormValue("name")
  userPassword := r.PostFormValue("password")
  
  user, err := database.New().GetUser(userName)
  if err != nil {
    fmt.Print("\nCan't find User in db")
    return
  }
  fmt.Print("\nVarifying User" + user.Name)

  check := user.CheckPassword(userPassword)
  if !check {
    fmt.Print("\nWrong password")
    return
  } else {
    fmt.Print("\nRight password")
    return
  }
  
}
