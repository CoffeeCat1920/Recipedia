package api

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"

	"github.com/gorilla/mux"
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

func authSameUser(r *http.Request) (bool) {
  // Getting the recipe
  vars := mux.Vars(r)
  recipeUUID := vars["uuid"] 
  recipe, err := database.New().GetRecipe(recipeUUID)
	if err != nil {                                                                                                                                                                                       
		fmt.Println("Can't find Recipe To Edit")
		return false
	}
 
	// Get session token
	c, err := r.Cookie("session-token")
	if err != nil {
		fmt.Println("Can't find Cookie")
		return false
	}

	// Get user session
	session, err := database.New().GetSession(c.Value)
	if err != nil {
		fmt.Printf("Can't find Session %s\n", c.Value)
		return false
	}
  
  // Verify the permission 
  if !(recipe.OwnerId == session.OwnerId) {
		fmt.Printf("You don't have ther permission to edit the recipe\n")
		return false
  }

  return true
}

func IsAdmin(r *http.Request) (bool) {
	// Get session token
	_, err := r.Cookie("admin-session-token")
	if err != nil {
		fmt.Println("Can't find Cookie")
		return false
	}

  return true
}
