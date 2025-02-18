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
