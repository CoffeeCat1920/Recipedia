package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s* service) AddSession(session *modals.Session) (error) {
  query := fmt.Sprintf(`
    INSERT INTO sessions(sessionid, ownerid, exp) 
    VALUES('%s', '%s', '%s')
  `, session.SessionId, session.OwnerId, session.Exp)

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}

func (s *service)GetSession(sessionId string) (*modals.Session, error) {
  var session modals.Session
  query := fmt.Sprintf("SELECT * FROM sessions WHERE sessionid = '%s';", sessionId)
  row := s.db.QueryRow(query)
  err := row.Scan(&session.SessionId, &session.OwnerId, &session.Exp)

  if err != nil {
    return nil, err 
  }

  return &session, nil
}

func (s *service) DeleteSession(sessionId string) error {
  _, err := s.db.Exec("DELETE FROM sessions WHERE sessionid = $1;", sessionId)
  if err != nil {
    fmt.Printf("Error deleting session: %v\n", err)
  }
  return err
}
