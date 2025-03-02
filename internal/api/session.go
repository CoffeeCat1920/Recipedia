package api

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
	"time"
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

func VerifyUser(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()  
  username := r.PostFormValue("name")
  password := r.PostFormValue("password")

  db := database.New()
   
  user, err := db.GetUserByName(username)
  if err != nil {
    fmt.Printf("\nCan't find user\n")
    http.Redirect(w, r, "/view/login", 302)
    return
  }
  if !user.CheckPassword(password) {
    fmt.Printf("\nPassword Incorrect\n")
    http.Redirect(w, r, "/view/login", 302)
    return
  } 
  
  session := createCookie(w, user.UUID)
  err = db.AddSession(session)
  if err != nil {
    panic(err)
  }

  fmt.Printf("\nSession created and cookie set. Session ID: %s, User UUID: %s\n", session.SessionId, user.UUID)
  http.Redirect(w, r, "/view/dashboard", 302)
}

func auth(r *http.Request) (*modals.Session, error) {
  cookie, err := r.Cookie("session-token")
  if err != nil {
    fmt.Printf("\nCan't find cookie\n")
    return nil, err 
  }

  sessionid := cookie.Value
  session, err := database.New().GetSession(sessionid)

  if err != nil {
    fmt.Printf("\nCan't find session %s in db case, %s\n", sessionid, err.Error())
    return nil, err 
  }

  fmt.Printf("\nFond session %s in db\n", sessionid)
  return session, nil
} 

func IsLoggedIn(r *http.Request) (bool) {
  _, err := auth(r)
  return err == nil
}

func LogInfo(r *http.Request) (interface{}, error) {
  session, err := auth(r) 
  if err != nil {
    return nil, err
  }

  user, err := database.New().GetUserByUUid(session.OwnerId)
  if err != nil {
    return nil, err
  }

  data := struct {
    Username string
  }{
    Username: user.Name,
  }
  
  return data, nil
} 

func LogOut(w http.ResponseWriter, r *http.Request) {
  session, err := auth(r) 

  if err != nil {
    fmt.Printf("\nCan't Autherize session with id %s in db", session.SessionId)
    http.Error(w, "Can't Logout", http.StatusInternalServerError)
    return 
  }

  err = database.New().DeleteSession(session.SessionId)
  if err != nil {
    fmt.Printf("\nCan't delete Session %s in db cause,\n%s", session.SessionId, err.Error())
    http.Error(w, "Can't Logout", http.StatusInternalServerError)
    return 
  }

  http.SetCookie(w, &http.Cookie{
    Name:     "session-token",
    Value:    "",
    Expires:  time.Unix(0, 0), // Expire immediately
    Path:     "/",
    HttpOnly: true,
    SameSite: http.SameSiteLaxMode,
  })

  http.Redirect(w, r, "/view/login", 302)
  return 
}

func LogedUser(r *http.Request) (*modals.User, error) {
  session, err := auth(r) 
  if err != nil {
    return nil, err
  }

  user, err := database.New().GetUserByUUid(session.OwnerId)
  fmt.Printf("\nThe sessions' uuid is the following, %s\n", session.OwnerId)
  if err != nil {
    return nil, err
  }

  return user, nil
}

func Auth( next http.HandlerFunc ) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    session, err := auth(r)
    if err != nil {
      http.Redirect(w, r, "/view/login", 302)
      return
    }

    user, err := database.New().GetUserByUUid(session.OwnerId)
    if err != nil {
      fmt.Printf("Error adding session: %v\n", err)
      http.Error(w, "Internal Server Error", http.StatusInternalServerError)
      return
    }
    _ = user
    next(w, r)
  }
}
