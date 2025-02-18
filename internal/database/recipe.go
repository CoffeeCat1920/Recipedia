package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s* service) AddRecipe(recipe *modals.Recipe) (error)  {
  query := fmt.Sprintf("INSERT INTO recipes(uuid, name, ownerid) VALUES('%s', '%s', '%s');", recipe.UUID, recipe.Name, recipe.OwnerId)

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}

func (s *service) GetAllRecipes() (*[]modals.Recipe, error) {
  var Recipes []modals.Recipe

  rows, err := s.db.Query("SELECT uuid, name, ownerId FROM recipes")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    var recipe modals.Recipe
    err := rows.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId)
    if err != nil {
      return nil, err
    }
    Recipes = append(Recipes, recipe)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }

  return &Recipes, nil
}
