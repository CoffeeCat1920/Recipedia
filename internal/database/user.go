package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s *service) AddUser(user *modals.User) (error) {
  query := fmt.Sprintf("INSERT INTO users(uuid, name, password) VALUES('%s', '%s', '%s');", user.UUID, user.Name, user.Password )

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}

func (s *service) GetUserUUid(name string) (string, error) {
  var user modals.User

  query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)
  err := s.db.QueryRow(query).Scan(&user.UUID, &user.Name, &user.Password)
  if err != nil {
    return "", err
  }
  
  return user.UUID, nil
}

func (s *service) GetUserByName(name string) (*modals.User, error) {
  var user modals.User

  query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s';", name)
  row := s.db.QueryRow(query)

  err := row.Scan(&user.UUID, &user.Name, &user.Password)
  if err != nil {
    return nil, err 
  }

  return &user, nil
}

func (s *service) GetUserByUUid(uuid string) (*modals.User, error) {
  var user modals.User

  query := fmt.Sprintf("SELECT * FROM users WHERE uuid = '%s';", uuid)
  row := s.db.QueryRow(query)

  err := row.Scan(&user.UUID, &user.Name, &user.Password)
  if err != nil {
    return nil, err 
  }

  return &user, nil
}

func (s *service) NumberOfUsers() (int, error) {
  var count int   

  q := "SELECT COUNT(*) FROM users"
  err := s.db.QueryRow(q).Scan(&count)
  if err != nil {
    return 0, err 
  }
  
  return count, nil 
}

func (s *service) GetAllUsers() ([]modals.User, error) {
  var users []modals.User
  query := "SELECT * FROM users;"
  
  rows, err := s.db.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  
  for rows.Next() {
    var user modals.User
    err := rows.Scan(&user.UUID, &user.Name, &user.Password)
    if err != nil {
      return nil, err
    }
    users = append(users, user)
  }
  
  if err = rows.Err(); err != nil {
    return nil, err
  }
  
  return users, nil
}
