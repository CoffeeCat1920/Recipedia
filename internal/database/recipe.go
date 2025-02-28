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

func (s *service) GetRecipe(uuid string) (*modals.Recipe, error) {
  var recipe modals.Recipe 
  
  row, err := s.db.Query("SELECT * FROM recipe WHERE uuid = %s", uuid) 
  if err != nil {
    return nil, err
  }

  err = row.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId, 0)
  if err != nil {
    return nil, err
  }

  return &recipe, nil
}

func (s *service) DeleteRecipe(uuid string) error {
  _, err := s.db.Query("DELETE * FROM recipe WHERE uuid = %s", uuid) 

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

func (s *service) GetRecipesByUser(username string) ([]modals.Recipe, error) {
  var recipes []modals.Recipe

  user, err := s.GetUserByName(username) 
  if err != nil {
    return nil, err
  }
  
  q := "SELECT * FROM recipes WHERE ownerid = $1" 
  rows, err := s.db.Query(q, user.UUID)

  for rows.Next() {
    var recipe modals.Recipe 
    err := rows.Scan(&recipe.UUID, recipe.Name, recipe.OwnerId, recipe.Views)

    if err != nil {
      return nil, err
    }

    recipes = append(recipes, recipe)
  }

  return recipes, nil
}
