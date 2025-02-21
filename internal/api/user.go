package api

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  user := modals.NewUser(r.PostFormValue("name"), r.PostFormValue("password"))
  err := database.New().AddUser(user)

  if err != nil {
    http.Error(w, "Can't add user", 403)
    return
  }

  fmt.Print("\nAdded user")

  http.Redirect(w, r, "/view/search", 302)
}
