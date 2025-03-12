package api

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"

	"github.com/gorilla/mux"
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

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  userUUID := vars["uuid"] 

  if !(IsAdmin(r)) {
		fmt.Println("Doesn't have the permission to edit user")
		http.Error(w, "Doesn't have the permission to edit user", http.StatusInternalServerError)
		return
  }

  err := database.New().DeleteUser(userUUID)
	if err != nil {
		fmt.Printf("Can't find Recipe To Edit cause, %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  http.Redirect(w, r, "/view/admin-dashboard", http.StatusFound)
}
