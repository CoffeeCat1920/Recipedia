package api

import (
	// "context"
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"

)

func setCookie(w http.ResponseWriter, user *modals.User) {
  s := modals.NewSession(user.UUID)

  err := database.New().AddSession(s)
  if err != nil {
    panic(err)
  }

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
  
  user, err := database.New().GetUserByName(userName)
  if err != nil {
    fmt.Print("\nCan't find User in db\n")
    return
  }
  fmt.Print("\nVarifying User " + user.Name)

  check := user.CheckPassword(userPassword)
  if !check {
    fmt.Print("\nWrong password\n")
    http.Redirect(w, r, "/view/login", 302)
  } else {
    fmt.Print("\nCorrect password\n")
    setCookie(w, user) 
    http.Redirect(w, r, "/view/dashboard", 303)
  }
}

// NOTE: Test fucntiont to test secure paths
func authorize(r *http.Request) (*modals.User, error) {
  c, err := r.Cookie("sessionCookie")
  if err != nil {
    return nil, err
  }

  sessionToken := c.Value
  fmt.Printf("\nThis is the session toke - %s\n", sessionToken)
  session, err := database.New().GetSession(sessionToken)
  if err != nil {
    return nil, err 
  }

  fmt.Printf("\nThis is the session toke in the database - %s\n", sessionToken)
  fmt.Printf("\nThis is the ownerid in the database - %s\n", sessionToken)
  
  // if session.IsExpired() {
  //   err = database.New().DeleteSession(sessionToken)
  //   return nil, err 
  // } 
  
  // if err != nil {
  //   return nil, err 
  // }
  user, err  := database.New().GetUserByUUid(session.OwnerId)
  return user, nil
}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    user, err := authorize(r)
    if err == nil {
      fmt.Printf("Can't verify user")
      http.Redirect(w, r, "/view/login", http.StatusFound)
    }
    _ = user
    // fmt.Printf("Verified user %s", user.Name)
    // ctx := context.WithValue(r.Context(), "user", user)
    // next(w, r.WithContext(ctx))
    // fmt.Printf("\nAuthorized user %s\n", user.Name)
    next(w, r)
  }
}
