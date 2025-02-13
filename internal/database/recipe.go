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
