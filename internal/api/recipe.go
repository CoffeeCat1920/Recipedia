package api

import (
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
)

func AddRecipe(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  
  // Addind the form to the database
  name := r.PostFormValue("name")
  ownerId, err := database.New().GetUserUUid(r.PostFormValue("ownername"))
  if err != nil {
    http.Error(w, "Can't find user of name" + r.PostFormValue("ownername"), 404)
    return
  }

  recipe := modals.NewRecipe(name, ownerId)

  err = database.New().AddRecipe(recipe)
  if err != nil {
    http.Error(w, "Can't find user of name" + r.PostFormValue("ownername"), 404)
    return
  }

  http.Redirect(w, r, "/view/search", 302)
}
