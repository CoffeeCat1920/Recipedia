package web

import (
	"net/http"
	"shiro/web/view"
)

func Render(name string) func(w http.ResponseWriter, r *http.Request) {

  return func(w http.ResponseWriter, r *http.Request) {
    view.RenderView(w, name, "base", nil)
  }

}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
//   recipes, err := database.New().GetAllRecipes() 
//   if err != nil {
//     http.Error(w, "No recipes Found", 404)
//     return
//   }
//   view.RenderView(w, "index", "base", recipes)
// } 
