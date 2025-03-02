package api

import (
	"net/http"
	"shiro/internal/modals"
)

func createCookie(w http.ResponseWriter, ownerId string) (*modals.Session) {
  session := modals.NewSession(ownerId) 
  exp, err := session.GetExpTime()
  if err != nil {
    panic(err)
  }

  http.SetCookie(w, &http.Cookie{
    Name: "session-token",
    Value: session.SessionId,
    Expires: exp,
    Path:     "/",           // Add this
    Domain:   "",
    HttpOnly: true,          // Add this
    SameSite: http.SameSiteLaxMode,  // Add this
  })
  return session
}

func createAdminCookie(w http.ResponseWriter) (*modals.AdminSession) {
  session := modals.NewAdminSession()
  exp, err := session.GetExpTime()
  if err != nil {
    panic(err)
  }

  http.SetCookie(w, &http.Cookie{
    Name: "admin-session-token",
    Value: session.SessionId,
    Expires: exp,
    Path:     "/",           // Add this
    Domain:   "",
    HttpOnly: true,          // Add this
    SameSite: http.SameSiteLaxMode,  // Add this
  })
  return session
}
