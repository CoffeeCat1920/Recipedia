package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"shiro/internal/database"
	"shiro/internal/modals"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func mdFileGenreator(content string, uuid string) (error) {
  
	directoryPath := "web/recipes/" + uuid 
	if !fileExists(directoryPath) {
		fmt.Println("Recipe Directory Doesn't Exists")
    return fmt.Errorf("Recipe Directory Doesn't Exists, %s", directoryPath) 
	}

  mdPath := directoryPath + "/recipe.md"
	mdFile, err := os.Create(mdPath)
	if err != nil {
		fmt.Printf("Error creating markdown file: %s\n", err.Error())
		return err
	}
	defer mdFile.Close()

	_, err = mdFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to markdown file: %s\n", err.Error())
		return err
	}

  return nil
} 

func htmlFileGenerator(recipeName string, content string, uuid string) (error) {
	directoryPath := "web/recipes/" + uuid 
	if !fileExists(directoryPath) {
		fmt.Println("Recipe Directory Doesn't Exists")
    return fmt.Errorf("Recipe Directory Doesn't Exists, %s", directoryPath) 
	}

  htmlFile := directoryPath + "/recipe.html"

  htmlContent := markdown.ToHTML([]byte(content), nil, nil)
  
  templateContent := fmt.Sprintf(`{{ define "head" }}
  <title> %s </title>
  {{ end }}
  {{ define "body" }}
  ` + string(htmlContent) + `
  {{ end }}`, recipeName)


  if !fileExists(htmlFile) {
    err := os.WriteFile(htmlFile, []byte(templateContent), fs.ModePerm)
    if err != nil {
      panic(err)
    }
  }
  
  return nil
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

  err = mdFileGenreator(content, recipe.UUID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  // Creating an html file
  err = htmlFileGenerator(recipe.Name, content, recipe.UUID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Success message
	fmt.Printf("Added recipe: %s, UUID: %s, by user: %s\n", recipe.Name, recipe.UUID, user.Name)
	fmt.Printf("Recipe content:\n%s\n", content)

	// Redirect user
  url := "/view/recipe/" + recipe.UUID

	http.Redirect(w, r, url, http.StatusFound)
}

func EditRecipe(w http.ResponseWriter, r *http.Request) {
  // Getting the form Inputs
	r.ParseForm()
	name := r.PostFormValue("name")
	content := r.PostFormValue("content")
   
  // Getting the recipe
  vars := mux.Vars(r)
  recipeUUID := vars["uuid"] 
  recipe, err := database.New().GetRecipe(recipeUUID)
	if err != nil {
		fmt.Println("Can't find Recipe To Edit")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  if !(authSameUser(r)) {
		fmt.Println("Doesn't have the permission to edit recipe")
		http.Error(w, "Doesn't have the permission to edit recipe", http.StatusInternalServerError)
		return
  }

  // Editing the old files
	directoryPath := "web/recipes/" + recipe.UUID
	if !fileExists(directoryPath) {
		fmt.Println("Recipe Directory Doesn't Exists")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
  
  // Editing md file
	mdPath := directoryPath + "/recipe.md"
	if !fileExists(mdPath) {
		fmt.Println("Recipe Markdown file already exists")
		http.Error(w, "Recipe Markdown file already exists", http.StatusInternalServerError)
		return
	}
  os.Remove(mdPath)

  err = mdFileGenreator(content, recipe.UUID)
  if err != nil {
    http.Error(w, "Can't generate new md", http.StatusInternalServerError)
  }


  htmlFile := directoryPath + "/recipe.html"  
  if !fileExists(htmlFile) {
		fmt.Println("Recipe Directory Doesn't Exists")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
  }
  os.Remove(htmlFile)

  err = htmlFileGenerator(name, content, recipe.UUID)
  if err != nil {
    http.Error(w, "Can't generate new md", http.StatusInternalServerError)
  }
  err = database.New().ChangeRecipeName(recipe.UUID, name)

	// Redirect user
  url := "/view/recipe/" + recipe.UUID

	http.Redirect(w, r, url, http.StatusFound)
} 

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  recipeUUID := vars["uuid"] 

	directoryPath := "web/recipes/" + recipeUUID

  if !(authSameUser(r)) {
		fmt.Println("Doesn't have the permission to edit recipe")
		http.Error(w, "Doesn't have the permission to edit recipe", http.StatusInternalServerError)
		return
  }

  err := database.New().DeleteRecipe(recipeUUID) 
	if err != nil {
		fmt.Printf("Can't find Recipe To Edit cause, %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  err = os.RemoveAll(directoryPath)

  http.Redirect(w, r, "/view/dashboard", http.StatusFound)
}

func EditRecipeAdmin(w http.ResponseWriter, r *http.Request) {
  // Getting the form Inputs
	r.ParseForm()
	name := r.PostFormValue("name")
	content := r.PostFormValue("content")
   
  // Getting the recipe
  vars := mux.Vars(r)
  recipeUUID := vars["uuid"] 
  recipe, err := database.New().GetRecipe(recipeUUID)
	if err != nil {
		fmt.Println("Can't find Recipe To Edit")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  if !(IsAdmin(r)) {
		fmt.Println("Doesn't have the permission to edit recipe")
		http.Error(w, "Doesn't have the permission to edit recipe", http.StatusInternalServerError)
		return
  }

  // Editing the old files
	directoryPath := "web/recipes/" + recipe.UUID
	if !fileExists(directoryPath) {
		fmt.Println("Recipe Directory Doesn't Exists")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
  
  // Editing md file
	mdPath := directoryPath + "/recipe.md"
	if !fileExists(mdPath) {
		fmt.Println("Recipe Markdown file already exists")
		http.Error(w, "Recipe Markdown file already exists", http.StatusInternalServerError)
		return
	}
  os.Remove(mdPath)

  err = mdFileGenreator(content, recipe.UUID)
  if err != nil {
    http.Error(w, "Can't generate new md", http.StatusInternalServerError)
  }


  htmlFile := directoryPath + "/recipe.html"  
  if !fileExists(htmlFile) {
		fmt.Println("Recipe Directory Doesn't Exists")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
  }
  os.Remove(htmlFile)

  err = htmlFileGenerator(name, content, recipe.UUID)
  if err != nil {
    http.Error(w, "Can't generate new md", http.StatusInternalServerError)
  }
  err = database.New().ChangeRecipeName(recipe.UUID, name)

	// Redirect user
  url := "/view/recipe/" + recipe.UUID

	http.Redirect(w, r, url, http.StatusFound)
}


func AdminDeleteRecipe(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  recipeUUID := vars["uuid"] 

	directoryPath := "web/recipes/" + recipeUUID

  if !(IsAdmin(r)) {
		fmt.Println("Doesn't have the permission to edit recipe")
		http.Error(w, "Doesn't have the permission to edit recipe", http.StatusInternalServerError)
		return
  }

  err := database.New().DeleteRecipe(recipeUUID) 
	if err != nil {
		fmt.Printf("Can't find Recipe To Edit cause, %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

  err = os.RemoveAll(directoryPath)

  http.Redirect(w, r, "/view/admin-manage-recipe", http.StatusFound)
}
