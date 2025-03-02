package view

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"shiro/internal/api"
	"shiro/internal/database"

	"github.com/gorilla/mux"
)

var views map[string]*template.Template

func init() {
  if views == nil {
    views = make(map[string]*template.Template)
  }

  views["index"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/templ/cards.html", "web/view/index.html", "web/view/base.html"))
  views["search"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/search.html", "web/view/base.html"))
  views["signup"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/signup.html", "web/view/base.html"))
  views["login"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/login.html", "web/view/base.html"))
  views["test"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/recipes/friedRice/recipe.html", "web/view/base.html"))
  views["upload-recipe"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/upload-recipe.html", "web/view/base.html"))
  views["dashboard"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/templ/userCard.html", "web/view/dashboard.html", "web/view/base.html"))
  views["edit-recipe"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/edit-recipe.html", "web/view/base.html"))
  views["mostViewed"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/templ/cards.html", "web/view/mostViewed.html", "web/view/base.html"))
}

func getData(r *http.Request, viewModel interface{}) (interface{}, error) {
  isLoggedIn := api.IsLoggedIn(r)
  logInfo, err := api.LogInfo(r)

  if err != nil {
    logInfo = nil
  }

  data := struct {
    LoggedIn  bool
    LogInfo interface{}
    ViewModel interface{}
  }{
    LoggedIn:  isLoggedIn,
    LogInfo: logInfo,
    ViewModel: viewModel,
  }

  return data, nil;
}


func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func RenderView( w http.ResponseWriter, r *http.Request, name string, tmpl string, viewModel interface{} ) {
  temp, ok := views[name] 
  if !ok {
    fmt.Printf("Can't find Page named %s, cause ", name)
    http.Error(w, "Page of name " + name + " not found", 404)
    return
  }
  
  data, err := getData(r, viewModel) 
  if err != nil {
    fmt.Printf("Can't render Page named %s cause %s", name, err.Error())
    http.Error(w, "Can't render page of name " + name, 404)
    return
  }

  err = temp.ExecuteTemplate(w, tmpl, data) 
  if err != nil {
    fmt.Printf("Can't render Page named %s cause %s", name, err.Error())
    http.Error(w, "Can't render page of name " + name, 404)
    return
  }
}

func RecipeHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  uuid := vars["id"]

  path := "web/recipes/" + uuid + "/recipe.html"    

  if !fileExists(path) {
    http.Error(w, "Can't find this recipe", http.StatusInternalServerError)
    fmt.Printf("Can't get the most viewed recipes cause, %s", uuid)
    return
  } 

  tmpl:= template.Must(template.ParseFiles("web/templ/topbar.html", path, "web/view/base.html"))

  data, err := getData(r, nil) 
  
  err = tmpl.ExecuteTemplate(w, "base", data)
  if err != nil {
    http.Error(w, "Can't find this recipe", http.StatusInternalServerError)
    fmt.Printf("Can't get the most viewed recipes cause, %s", uuid)
    return
  }

  db := database.New()
  err = db.RecipeAddView(uuid)
  if err != nil {
    http.Error(w, "Can't add view to the recipe", http.StatusInternalServerError)
    fmt.Printf("Can't add view to recipe cause, %s", err.Error())
    return
  }
  
  return
}
