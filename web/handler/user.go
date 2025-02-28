package web

import (
	"net/http"
	"shiro/internal/api"
	"shiro/internal/database"
	"shiro/internal/modals"
	"shiro/web/view"
)

func DashBoardsHandler(w http.ResponseWriter, r *http.Request) {
  user, err := api.LoggedUserName(r)
  if err != nil {
    http.Error(w, "Can't get the requested user", http.StatusNotFound)
    return
  }

  userUUID, err := database.New().GetUserUUid(user)
  if err != nil {
    http.Error(w, "Can't get the requested user", http.StatusNotFound)
    return
  }

  recipes, err := database.New().GetRecipesByUser(userUUID)
  if err != nil {
    http.Error(w, "Can't get the requested user", http.StatusNotFound)
    return
  }


  data := struct {
    Recipes []modals.Recipe  
  }{
    Recipes: recipes,
  }

  view.RenderView(w, r, "dashboard", "base", data)
}
