package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s* service) AddSession(session *modals.Session) (error) {
  query := fmt.Sprintf(`
    INSERT INTO sessions(sessionid, ownerid, exp) 
    VALUES('%s', '%s', '%s')
  `, session.SessionId, session.OwnerId, session.Exp.String())

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}

func (s *service)GetSession(sessionId string) (*modals.Session, error) {
  var session modals.Session
  row := s.db.QueryRow(`SELECT * FROM sessions WHERE sessionid = '%s';`, sessionId)
  err := row.Scan(&session.SessionId, &session.OwnerId, &session.Exp)
  
  fmt.Printf("\nTaken Session from db - %s\n", session.SessionId)

  if err != nil {
    return nil, err 
  }
  return &session, nil
}


func (s *service) DeleteSession(sessionId string) (error) {
  _, err := s.db.Exec(`DELETE FROM sessions WHERE sessionid = '%s';`, sessionId)

  if err != nil {
    return  err 
  } else {
    return nil
  }
}
