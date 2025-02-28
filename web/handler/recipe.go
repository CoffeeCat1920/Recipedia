package web

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
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

func SearchRecipes(w http.ResponseWriter, r *http.Request) {
  searchterm := r.URL.Query().Get("search")

  fmt.Printf("\n%s\n", searchterm)

  recipes, err := database.New().SearchRecipe(searchterm)

  searchdata := struct {
    SearchTerm string
    Recipes []modals.Recipe  
  }{
    SearchTerm: searchterm,
    Recipes: recipes,
  }

  if err != nil {
    http.Error(w, "No recipes found", http.StatusInternalServerError)
    fmt.Printf("Can't get the searched recipes, %s", err.Error())
    return
  } 
  
  view.RenderView(w, r, "mostViewed", "base", searchdata)
} 


func RecipeHandler(w http.ResponseWriter, r *http.Request) {
  view.RecipeHandler(w, r)
}
