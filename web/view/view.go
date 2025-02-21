package view

import (
	"fmt"
	"html/template"
	"net/http"
	"shiro/internal/api"
)

var views map[string]*template.Template

func init() {
  if views == nil {
    views = make(map[string]*template.Template)
  }

  views["search"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/search.html", "web/view/base.html"))
  views["signup"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/signup.html", "web/view/base.html"))
  views["login"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/login.html", "web/view/base.html"))
  views["test"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/recipes/friedRice/recipe.html", "web/view/base.html"))
  views["upload-recipe"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/upload-recipe.html", "web/view/base.html"))
  views["dashboard"] = template.Must(template.ParseFiles("web/templ/topbar.html", "web/view/dashboard.html", "web/view/base.html"))
}

func RenderView( w http.ResponseWriter, r *http.Request, name string, tmpl string, viewModel interface{} ) {
  temp, ok := views[name] 
  if !ok {
    fmt.Printf("Can't find Page named %s", name)
    http.Error(w, "Page of name " + name + " not found", 404)
    return
  }
  
  isLoggedIn := api.IsLoggedIn(r)
  data := struct {
    LoggedIn  bool
    ViewModel interface{}
  }{
    LoggedIn:  isLoggedIn,
    ViewModel: viewModel,
  }

  err := temp.ExecuteTemplate(w, tmpl, data) 
  if err != nil {
    fmt.Printf("Can't find Page named %s", name )
    http.Error(w, "Can't find page of name " + name, 404)
    return
  }
}
