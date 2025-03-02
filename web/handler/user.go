package web

import (
	"fmt"
	"net/http"
	"shiro/internal/api"
	"shiro/internal/database"
	"shiro/internal/modals"
	"shiro/web/view"
)

func DashBoardsHandler(w http.ResponseWriter, r *http.Request) {
  user, err := api.LogedUser(r)
  if err != nil {
    fmt.Printf("\nCan't get the requested user cause, %s\n", err.Error())
    http.Error(w, "Can't get the requested user", http.StatusNotFound)
    return
  }

  recipes, err := database.New().GetRecipesByUser(user.UUID)
  if err != nil {
    fmt.Printf("\nCan't get the requested user's recipes cause, %s\n", err.Error())
  }

  data := struct {
    Recipes []modals.Recipe  
  }{
    Recipes: recipes,
  }

  fmt.Print(recipes)

  view.RenderView(w, r, "dashboard", "base", data)
}
