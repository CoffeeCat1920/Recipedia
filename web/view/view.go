package view 

import (
	"html/template"
	"net/http"
)

var views map[string]*template.Template

func init() {
  if views == nil {
    views = make(map[string]*template.Template)
  }
  views["search"] = template.Must(template.ParseFiles("web/view/search.html", "web/view/base.html"))
  views["signup"] = template.Must(template.ParseFiles("web/view/signup.html", "web/view/base.html"))
  views["login"] = template.Must(template.ParseFiles("web/view/login.html", "web/view/base.html"))
  views["test"] = template.Must(template.ParseFiles("web/recipes/friedRice/recipe.html", "web/view/base.html"))
  views["upload-recipe"] = template.Must(template.ParseFiles("web/view/upload-recipe.html", "web/view/base.html"))
}

func RenderView( w http.ResponseWriter, name string, tmpl string, viewModel interface{} ) {
  temp, ok := views[name] 
  if !ok {
    http.Error(w, "Page of name " + name + " not found", 404)
  }
  
  err := temp.ExecuteTemplate(w, tmpl, viewModel) 
  if err != nil {
    http.Error(w, "Can't find page of name " + name, 404)
  }
}
