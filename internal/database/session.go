package database

import (
	"fmt"
	"shiro/internal/modals"
)

func (s* service) AddSession(session *modals.Session) (error) {
  query := fmt.Sprintf(`
    INSERT INTO sessions(userid, sessionid, exp) 
    VALUES('%s', '%s', '%s')
  `, session.UserId, session.SessionId, session.Exp)

  _, err := s.db.Exec(query) 

  if err != nil {
    return err
  }

  return nil
}

func (s *service)GetSession(sessionId string) (*modals.Session, error) {
  var session *modals.Session
  row := s.db.QueryRow(`SELECT * FROM sessions WHERE ownerid = '%s';`, sessionId)
  err := row.Scan(&session.UserId, &session.SessionId, &session.Exp)

  if err != nil {
    return nil, err 
  } else {
    return session, nil
  }
}


func (s *service) DeleteSession(sessionId string) (error) {
  _, err := s.db.Exec(`DELETE FROM sessions WHERE ownerid = '%s';`, sessionId)

  if err != nil {
    return  err 
  } else {
    return nil
  }
}
