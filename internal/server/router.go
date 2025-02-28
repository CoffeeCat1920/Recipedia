package server

import (
	"net/http"
	"shiro/internal/api"
	web "shiro/web/handler"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRouts() http.Handler {
  r := mux.NewRouter()

  r.HandleFunc("/test", web.Render("test", nil))

  // Homepage
  r.HandleFunc("/", web.IndexHandler)
  r.HandleFunc("/view/home", web.IndexHandler)

  // Views
  r.HandleFunc("/view/signup", web.Render("signup", nil))
  r.HandleFunc("/view/login", web.Render("login", nil))
  r.HandleFunc("/view/mostViewed", web.MostViewHandler)

  // Recipes
  r.HandleFunc("/view/recipe/{id}", web.RecipeHandler)
  r.HandleFunc("/view/search", web.SearchRecipes)

  // Secure Views
  r.HandleFunc("/view/upload-recipe", api.Auth( web.Render("upload-recipe", nil) ))
  r.HandleFunc("/view/dashboard", api.Auth( web.Render("dashboard", nil) ))  

  // Api
  r.HandleFunc("/api/add-user", api.AddUser).Methods("POST")
  r.HandleFunc("/api/verify-user", api.VerifyUser).Methods("POST")

  // Secure Api
  r.HandleFunc("/api/add-recipe",  api.Auth( api.AddRecipe )).Methods("POST")
  r.HandleFunc("/api/log-out",  api.Auth( api.LogOut ))

  //Css
  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

  return r
}
