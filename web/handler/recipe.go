package web

import (
	"fmt"
	"net/http"
	"os"
	"shiro/internal/database"
	"shiro/internal/modals"
	"shiro/web/view"

	"github.com/gorilla/mux"
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

func EditRecipeHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid := vars["uuid"]
  
  recipe, err := database.New().GetRecipe(uuid)
  if err != nil {
    http.Error(w, "No recipe found to edit", http.StatusInternalServerError)
    fmt.Printf("Can't get the recipe to edit, %s", err.Error())
    return
  }
  
  path := "web/recipes/" + recipe.UUID 
  mdContent, err := os.ReadFile(path + "/recipe.md")  
  if err != nil {
    http.Error(w, "No recipe found to edit", http.StatusInternalServerError)
    fmt.Printf("Can't find the recipe.md to edit, %s", err.Error())
    return
  }

  data := struct {
    RecipeName string
    RecipeContent string
    UUID string
  }{
    RecipeName: recipe.Name,
    RecipeContent: string(mdContent),
    UUID: recipe.UUID,
  }

  fmt.Printf("\nThe Content of the recipe\n%s", string(mdContent))

  view.RenderView(w, r, "edit-recipe", "base", data) 
}


func AdminEditRecipeHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid := vars["uuid"]
  
  recipe, err := database.New().GetRecipe(uuid)
  if err != nil {
    http.Error(w, "No recipe found to edit", http.StatusInternalServerError)
    fmt.Printf("Can't get the recipe to edit, %s", err.Error())
    return
  }
  
  path := "web/recipes/" + recipe.UUID 
  mdContent, err := os.ReadFile(path + "/recipe.md")  
  if err != nil {
    http.Error(w, "No recipe found to edit", http.StatusInternalServerError)
    fmt.Printf("Can't find the recipe.md to edit, %s", err.Error())
    return
  }

  data := struct {
    RecipeName string
    RecipeContent string
    UUID string
  }{
    RecipeName: recipe.Name,
    RecipeContent: string(mdContent),
    UUID: recipe.UUID,
  }

  fmt.Printf("\nThe Content of the recipe\n%s", string(mdContent))

  view.RenderView(w, r, "admin-edit-recipe", "base", data) 
}
