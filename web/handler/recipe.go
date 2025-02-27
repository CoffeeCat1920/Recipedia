package web 

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/web/view"
)

func MostViewHandler(w http.ResponseWriter, r *http.Request) {
  recipes, err := database.New().MostViewedRecipes()
  if err != nil {
    http.Error(w, "Can't get the most viewed recipes", http.StatusInternalServerError)
    fmt.Printf("Can't get the most viewed recipes cause, %s", err.Error())
    return
  }
  view.RenderView(w, r, "mostViewed", "base", recipes)
}


func RecipeHandler(w http.ResponseWriter, r *http.Request) {
  view.RecipeHandler(w, r)
}
