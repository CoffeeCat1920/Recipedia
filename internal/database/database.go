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
  AddUser(user *modals.User) (error) 
  GetUser(uuid string) (*modals.User, error) 
  GetUserUUid(name string) (string, error) 
  AddRecipe(recipe *modals.Recipe) (error)  
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
