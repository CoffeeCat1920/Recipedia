package server

import (
	"net/http"
	"os"
	"shiro/internal/api"
	web "shiro/web/handler"
  "io/fs"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
)


func fileExists(path string) bool {
  _, err := os.Stat(path)
  if err != nil {
    return !os.IsNotExist(err) // Correctly handles cases where `err` is not nil
  }
  return true
}

func (s *Server) RenderMarkdown() {

  file := "web/recipes/friedRice/recipe.md"
  
  content, err := os.ReadFile(file)
  if err != nil {
    panic(err)
  } 

  htmlContent := markdown.ToHTML(content, nil, nil)

  wrappedContent := `{{ define "head" }}
  <h1> Fried Rice </h1>
  {{ end }}

  {{ define "body" }}

  ` + string(htmlContent) + `

  {{ end }}`
  
  outputPath := "web/recipes/friedRice/recipe.html"

  if !fileExists(outputPath) {
    err = os.WriteFile(outputPath, []byte(wrappedContent), fs.ModePerm)
    if err != nil {
      panic(err)
    }
  }

}

func (s *Server) RegisterRouts() http.Handler {

  s.RenderMarkdown()

  r := mux.NewRouter()

  r.HandleFunc("/test", web.Render("test", nil))

  // Views
  r.HandleFunc("/view/search", web.Render("search", nil))
  r.HandleFunc("/view/signup", web.Render("signup", nil))
  r.HandleFunc("/view/upload-recipe", web.Render("upload-recipe", nil))
  r.HandleFunc("/view/dashboard", api.Auth( web.Render("dashboard", nil) ))  
  r.HandleFunc("/view/login", web.Render("login", nil))
  // r.HandleFunc("/view/index", web.IndexHandler)

  // Api
  r.HandleFunc("/api/add-user", api.AddUser).Methods("POST")
  r.HandleFunc("/api/verify-user", api.VerifyUser).Methods("POST")
  r.HandleFunc("/api/add-recipe", api.AddRecipe).Methods("POST")
  
  return r

}
