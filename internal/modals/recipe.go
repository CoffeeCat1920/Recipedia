package modals

import (
	"github.com/google/uuid"
)

type Recipe struct {
  UUID string `json:"uuid"` 
  Name string `json:"name"`
  OwnerId string `json:"ownerId"`
}

func NewRecipe(name string, ownerId string) *Recipe{
  
  recipe := &Recipe{
    UUID: uuid.NewString(),
    Name: name,
    OwnerId: ownerId,
  }

  return recipe

}
