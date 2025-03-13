package web

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	"shiro/internal/modals"
	"shiro/web/view"
)

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
  db := database.New()

  nor, err := db.NumberOfRecipes()
  if err != nil {
    http.Error(w, "Can't get the number of recipes", http.StatusInternalServerError)
    fmt.Printf("\nCan't get the number of recipes cause, %s\n", err.Error())
    return
  }

  nou, err := db.NumberOfUsers()
  if err != nil {
    http.Error(w, "Can't get the number of users", http.StatusInternalServerError)
    fmt.Printf("\nCan't get the number of users cause, %s\n", err.Error())
    return
  }

  data := struct {
    NOU int
    NOR int 
  }{
    NOU: nou,
    NOR: nor,
  }

  view.RenderView(w, r, "admin-dashboard", "base", data)
}

func AdminManageRecipes(w http.ResponseWriter, r *http.Request) {
  recipes, err := database.New().GetAllRecipe() 
  if err != nil {
    http.Error(w, "Can't get all the recipes", http.StatusInternalServerError)
    fmt.Printf("\nCan't get all the recipes, %s\n", err.Error())
    return
  }

  data := struct {
    Recipes []modals.Recipe
  }{
    Recipes: recipes,
  }

  view.RenderView(w, r, "admin-manage-recipes", "base", data) 
}

func AdminManageUsers(w http.ResponseWriter, r *http.Request) {
  users, err := database.New().GetAllUsers() 
  if err != nil {
    http.Error(w, "Can't get all the recipes", http.StatusInternalServerError)
    fmt.Printf("\nCan't get all the recipes, %s\n", err.Error())
    return
  }

  data := struct {
    Users []modals.User
  }{
    Users: users,
  }

  view.RenderView(w, r, "admin-manage-users", "base", data) 
}
