package server

import (
	"fmt"
	"net/http"
	"shiro/internal/database"
	// "os"
	// "strconv"
)

type Server struct {
  port int
  db database.Service
}

func NewServer() *http.Server {
 
  // port, _ := strconv.Atoi(os.Getenv("PORT"))
  port := 42069 

  NewServer := &Server {
    port: port,
    db: database.New(),
  }

  server := &http.Server{
    Addr: fmt.Sprintf(":%d", NewServer.port), 
    Handler: NewServer.RegisterRouts(),
  }

  return server

}
