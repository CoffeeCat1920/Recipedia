package main

import (
	"fmt"
	"log"
	"shiro/internal/server"
	_ "github.com/lib/pq"
)

func Shutdown() {
  log.Println("Shutdown")
}

func main() {

  s := server.NewServer()

  err := s.ListenAndServe()
  if err != nil {
    panic(fmt.Sprintf("http server error: %s", err))
  }

  Shutdown()

}
