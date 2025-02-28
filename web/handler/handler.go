package web

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
	"shiro/web/view"
)

func Render(name string, data interface{}) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    view.RenderView(w, r, name, "base", data)
  }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  recipes, err := database.New().MostViewedRecipes()

  data := struct {
    Recipes []modals.Recipe  
  }{
    Recipes: recipes,
  }

  if err != nil {
    http.Error(w, "Can't get the most viewed recipes", http.StatusInternalServerError)
    fmt.Printf("Can't get the most viewed recipes cause, %s", err.Error())
    return
  }
  
  view.RenderView(w, r, "index", "base", data)
}
