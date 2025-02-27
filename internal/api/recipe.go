package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"shiro/internal/database"
	"shiro/internal/modals"

	"github.com/gomarkdown/markdown"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func AddRecipe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("name")
	content := r.PostFormValue("content")

	// Get session token
	c, err := r.Cookie("session-token")
	if err != nil {
		fmt.Println("Can't find Cookie")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get user session
	session, err := database.New().GetSession(c.Value)
	if err != nil {
		fmt.Printf("Can't find Session %s\n", c.Value)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get user from DB
	user, err := database.New().GetUserByUUid(session.OwnerId)
	if err != nil {
		fmt.Println("Can't find User")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create recipe object
	recipe := modals.NewRecipe(name, user.UUID)
	err = database.New().AddRecipe(recipe)
	if err != nil {
		fmt.Printf("Can't add recipe to the DB: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create directory for recipe
	directoryPath := "web/recipes/" + recipe.UUID
	if fileExists(directoryPath) {
		fmt.Println("Recipe Directory Already Exists")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = os.Mkdir(directoryPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create the Markdown file
	mdPath := directoryPath + "/recipe.md"
	if fileExists(mdPath) {
		fmt.Println("Recipe Markdown file already exists")
		http.Error(w, "Recipe Markdown file already exists", http.StatusInternalServerError)
		return
	}

	mdFile, err := os.Create(mdPath)
	if err != nil {
		fmt.Printf("Error creating markdown file: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer mdFile.Close()

	// Write content to file
	_, err = mdFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to markdown file: %s\n", err.Error())
		http.Error(w, "Can't write to markdown file", http.StatusInternalServerError)
		return
	}

  // Creating an html file


  htmlContent := markdown.ToHTML([]byte(content), nil, nil)
  
  templateContent := fmt.Sprintf(`{{ define "head" }}
  <title> %s </title>
  {{ end }}
  {{ define "body" }}
  ` + string(htmlContent) + `
  {{ end }}`, recipe.Name)
  

  htmlFile := directoryPath + "/recipe.html"  
  if !fileExists(htmlFile) {
    err = os.WriteFile(htmlFile, []byte(templateContent), fs.ModePerm)
    if err != nil {
      panic(err)
    }
  }

	// Success message
	fmt.Printf("Added recipe: %s, UUID: %s, by user: %s\n", recipe.Name, recipe.UUID, user.Name)
	fmt.Printf("Recipe content:\n%s\n", content)

	// Redirect user
  url := "/view/recipe/" + recipe.UUID

	http.Redirect(w, r, url, http.StatusFound)
}
