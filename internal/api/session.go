package api

import (
	"context"
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
)

func setCookie(w http.ResponseWriter, user *modals.User) {
  s := modals.NewSession(user.UUID)
  database.New().AddSession(s)
  http.SetCookie(w, &http.Cookie{
    Name: "sessionCookie",
    Value: s.SessionId,
    Expires: s.Exp,
    HttpOnly: true,                
    Secure:   true,                 
    SameSite: http.SameSiteStrictMode, 
    Path:     "/",   
  })
}

// TODO: Better errors for login stuff
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
    http.Redirect(w, r, "/view/login", 302)
    return
  } else {
    fmt.Print("\nCorrect password")
    setCookie(w, user) 
    http.Redirect(w, r, "/view/dashboard", 303)
    return
  }
}

// NOTE: Test fucntiont to test secure paths
func authorize(w http.ResponseWriter, r *http.Request) (*modals.User) {
  c, err := r.Cookie("sessionCookie")
  if err != nil {
    http.Redirect(w, r, "/view/login", 400)
    return nil
  }

  sessionToken := c.Value
  session, err := database.New().GetSession(sessionToken)
  if err != nil {
    http.Redirect(w, r, "/view/login", 400)
    return nil 
  }
  
  if session.IsExpired() {
    err = database.New().DeleteSession(sessionToken)
    return nil 
  } 

  user, err := database.New().GetUser(session.UserId) 
  if err != nil {
    http.Redirect(w, r, "/view/login", 400)
    return nil 
  }

  return user
}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if user := authorize(w, r); user == nil {
      http.Redirect(w, r, "/view/login", 302)
    } else {
      ctx := context.WithValue(r.Context(), "user", user)
      r = r.WithContext(ctx)
      next(w, r)
    }
  }
} 
