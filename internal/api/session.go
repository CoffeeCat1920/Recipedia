package api

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
)

func VerifyUser(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  userName := r.PostFormValue("name")
  userPassword := r.PostFormValue("password")
  
  user, err := database.New().GetUser(userName)
  if err != nil {
    fmt.Print("\nCan't find User in db")
    return
  }
  fmt.Print("\nVarifying User " + user.Name)

  check := user.CheckPassword(userPassword)
  if !check {
    fmt.Print("\nWrong password")
    w.WriteHeader(http.StatusUnauthorized)
    return
  } else {
    fmt.Print("\nRight password")
    w.WriteHeader(http.StatusOK)
    return
  }
}
