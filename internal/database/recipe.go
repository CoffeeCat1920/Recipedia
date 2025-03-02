package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s* service) AddRecipe(recipe *modals.Recipe) (error)  {
  query := fmt.Sprintf("INSERT INTO recipes(uuid, name, ownerid, views) VALUES('%s', '%s', '%s', 0);", recipe.UUID, recipe.Name, recipe.OwnerId)

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}


func (s *service) DeleteRecipe(uuid string) error {
  q := "DELETE FROM recipes WHERE uuid = $1" 

  _, err := s.db.Query(q, uuid) 

  if err != nil {
    return  err
  }

  return nil
}

func (s *service) MostViewedRecipes() ([]modals.Recipe, error) {
  var recipes []modals.Recipe

  rows, err := s.db.Query("SELECT * FROM recipes ORDER BY views LIMIT 10;")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var recipe modals.Recipe
    err := rows.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, &recipe.Views)
    if err != nil {
      return nil, err
    }
    recipes = append(recipes, recipe) 
  }

  return recipes, nil
} 


func (s *service) SearchRecipe(name string) ([]modals.Recipe, error) {
  var recipes []modals.Recipe

  searchTerm := "%" + name + "%" 
  query := "SELECT * FROM recipes WHERE name ILIKE $1"

  // Log the query and parameter
  fmt.Println("Executing Query:", query, "with parameter:", searchTerm)

  rows, err := s.db.Query(query, searchTerm)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var recipe modals.Recipe
    err := rows.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, &recipe.Views)
    if err != nil {
      return nil, err
    }
    recipes = append(recipes, recipe)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }

  return recipes, nil
}

func (s *service) RecipeAddView(uuid string) error {
  q := "UPDATE recipes SET views = views + 1 WHERE uuid = $1;"

  _, err := s.db.Exec(q, uuid)
  if err != nil {
    return err
  }

  fmt.Printf("Added view for recipe with UUID %s\n", uuid)
  return nil
}

func (s *service) GetRecipesByUser(uuid string) ([]modals.Recipe, error) {
  var recipes []modals.Recipe

  q := "SELECT * FROM recipes WHERE ownerid = $1" 
  rows, err := s.db.Query(q, uuid)

  if err != nil {
    return nil, err
  }

  for rows.Next() {
    var recipe modals.Recipe 
    err := rows.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, &recipe.Views)

    if err != nil {
      return nil, err
    }

    recipes = append(recipes, recipe)
  }

  return recipes, nil
}

func (s *service) GetRecipe(uuid string) (*modals.Recipe, error) {
  var recipe modals.Recipe 
   
  q := "SELECT * FROM recipes WHERE uuid = $1" 
  row := s.db.QueryRow(q, uuid)

  err := row.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, &recipe.Views)
  if err != nil {
    return nil, err
  }

  return &recipe, nil
}

func (s *service) ChangeRecipeName(uuid string, newName string) (error) {
  q := "UPDATE recipes SET name = $1 WHERE uuid = $2"
  _, err := s.db.Exec(q, newName, uuid)
  if err != nil {
    return err
  }

  return nil
}

func (s *service) NumberOfRecipes() (int, error) {
  var count int   

  q := "SELECT COUNT(*) FROM recipes"
  err := s.db.QueryRow(q).Scan(&count)
  if err != nil {
    return 0, err 
  }
  
  return count, nil 
}

func (s *service) GetAllRecipe() ([]modals.Recipe, error) {
  var recipes []modals.Recipe

  rows, err := s.db.Query("SELECT * FROM recipes;")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var recipe modals.Recipe
    err := rows.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, &recipe.Views)
    if err != nil {
      return nil, err
    }
    recipes = append(recipes, recipe) 
  }

  return recipes, nil
}
