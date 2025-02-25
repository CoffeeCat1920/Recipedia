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

func (s *service) GetRecipe(uuid string) (*modals.Recipe, error) {
  var recipe modals.Recipe 
  
  row, err := s.db.Query("SELECT * FROM recipe WHERE uuid = %s", uuid) 
  if err != nil {
    return nil, err
  }

  err = row.Scan(&recipe.UUID, &recipe.Name, &recipe.OwnerId)
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

// DONE: Make it work
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

