package api

import (
	"fmt"
	"net/http"
	"shiro/internal/admin.go"
	"shiro/internal/database"
	"shiro/internal/modals"
	"time"
)

func VerifyAdmin(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()  
  password := r.PostFormValue("password")

  ad := admin.New()
  if !(ad.CheckPassword(password)) {
    http.Error(w, "Wrong Password", http.StatusUnauthorized)
    return
  }

  db := database.New()

  session := createAdminCookie(w) 
  err := db.AddAdminSession(session) 
  if err != nil {
    panic(err)
  }

  fmt.Printf("\nSession created and cookie set. Session ID: %s\n", session.SessionId)
  http.Redirect(w, r, "/view/admin-dashboard", 302)
}


func authAdmin(r *http.Request) (*modals.AdminSession, error) {
  cookie, err := r.Cookie("admin-session-token")
  if err != nil {
    fmt.Printf("\nCan't find cookie\n")
    return nil, err 
  }

  sessionid := cookie.Value
  session, err := database.New().GetAdminSession(sessionid)

  if err != nil {
    fmt.Printf("\nCan't find session %s in db case, %s\n", sessionid, err.Error())
    return nil, err 
  }

  fmt.Printf("\nFond session %s in db\n", sessionid)
  return session, nil
} 

func AdminAuth(next http.HandlerFunc) (http.HandlerFunc) {
  return func(w http.ResponseWriter, r *http.Request) {
    _, err := authAdmin(r)
    if err != nil {
      http.Redirect(w, r, "/view/admin-login", 302)
      fmt.Printf("Can't log the admin in cause, %s", err.Error())
      return
    }

    next(w, r)
  }
}

func AdminLogout(w http.ResponseWriter, r *http.Request) {
  session, err := authAdmin(r) 

  if err != nil {
    fmt.Printf("\nCan't Autherize session with id %s in db", session.SessionId)
    http.Error(w, "Can't Logout", http.StatusInternalServerError)
    return 
  }

  err = database.New().DeleteAdminSession(session.SessionId)
  if err != nil {
    fmt.Printf("\nCan't delete Admin Session %s in db cause,\n%s", session.SessionId, err.Error())
    http.Error(w, "Can't Logout", http.StatusInternalServerError)
    return 
  }

  http.SetCookie(w, &http.Cookie{
    Name:     "admin-session-token",
    Value:    "",
    Expires:  time.Unix(0, 0), // Expire immediately
    Path:     "/",
    HttpOnly: true,
    SameSite: http.SameSiteLaxMode,
  })

  http.Redirect(w, r, "/view/admin-login", 302)
  return 
} 
