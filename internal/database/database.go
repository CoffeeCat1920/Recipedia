package database

import (
	"database/sql"
	"fmt"
	"log"
	"shiro/internal/modals"

	_"github.com/lib/pq"
)

type Service interface {
  Close() error
  // User
  AddUser(user *modals.User) (error) 
  GetUserByUUid(uuid string) (*modals.User, error) 
  GetUserByName(name string) (*modals.User, error) 
  GetUserUUid(name string) (string, error) 
  
  // Recipes
  AddRecipe(recipe *modals.Recipe) (error)  
  GetRecipe(uuid string) (*modals.Recipe, error) 
  DeleteRecipe(uuid string) error 
  MostViewedRecipes() ([]modals.Recipe, error) 
  SearchRecipe(name string) ([]modals.Recipe, error) 
  RecipeAddView(uuid string) (error) 

  // Session
  AddSession(session *modals.Session) (error)
  GetSession(sessionId string) (*modals.Session, error) 
  DeleteSession(sessionId string) (error) 
}

type service struct {
  db *sql.DB   
}

var (
  database = "Production"
  username = "postgres"
  password = "123"
  port = "5432"
  host = "localhost"
	dbInstance *service
)

func New() Service {
  if dbInstance != nil {
    return dbInstance
  }
  
  connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
  
  db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
