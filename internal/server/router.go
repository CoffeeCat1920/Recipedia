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
  r.HandleFunc("/view/admin-login", web.Render("admin-login", nil))

  // Recipes
  r.HandleFunc("/view/recipe/{id}", web.RecipeHandler)
  r.HandleFunc("/view/search", web.SearchRecipes)
  r.HandleFunc("/view/edit/{uuid}", web.EditRecipeHandler)

  // Secure Views
  r.HandleFunc("/view/upload-recipe", api.Auth( web.Render("upload-recipe", nil) ))
  r.HandleFunc("/view/dashboard", api.Auth(web.DashBoardsHandler))  

  // Api
  r.HandleFunc("/api/add-user", api.AddUser).Methods("POST")
  r.HandleFunc("/api/verify-user", api.VerifyUser).Methods("POST")

  // Secure Api
  r.HandleFunc("/api/add-recipe",  api.Auth( api.AddRecipe )).Methods("POST")
  r.HandleFunc("/api/log-out",  api.Auth( api.LogOut ))
  r.HandleFunc("/api/edit-recipe/{uuid}",  api.Auth( api.EditRecipe ))
  r.HandleFunc("/api/delete/{uuid}",  api.Auth( api.DeleteRecipe )) 

  // Admin Api
  r.HandleFunc("/api/verify-admin",  api.VerifyAdmin)
  r.HandleFunc("/api/admin-edit-recipe/{uuid}", api.AdminAuth( api.EditRecipeAdmin ) )
  r.HandleFunc("/api/admin-delete-recipe/{uuid}", api.AdminAuth( api.AdminDeleteRecipe) )
  r.HandleFunc("/api/admin-delete-user/{uuid}", api.AdminAuth( api.AdminDeleteUser) )
  r.HandleFunc("/api/admin-logout", api.AdminAuth( api.AdminLogout) )

  // Admin views
  r.HandleFunc("/view/admin-dashboard", api.AdminAuth( web.AdminDashboardHandler ) )
  r.HandleFunc("/view/admin-manage-recipe", api.AdminAuth( web.AdminManageRecipes ) )
  r.HandleFunc("/view/admin-manage-user", api.AdminAuth( web.AdminManageUsers ) )
  r.HandleFunc("/view/admin-edit/{uuid}", api.AdminAuth( web.AdminEditRecipeHandler) )

  //Css
  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

  return r
}
