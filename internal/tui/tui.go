package tui

import (
	"fmt"
	"log"
	"shiro/internal/database"
	"shiro/internal/modals"
)


type TUI struct {
  db database.Service
} 

var (
  tuiInstance *TUI = nil
)

func New() *TUI {
  if (tuiInstance != nil) {
    return tuiInstance
  }
  tuiInstance = &TUI {
    db: database.New(),
  }
  return tuiInstance
}

func (tui *TUI) AddUser() {

  var name string
  var password string
  log.Printf("Enter Name: ")
  fmt.Scan(&name)
  log.Printf("Enter password: ")
  fmt.Scan(&password)

  user := modals.NewUser(name, password)
  tui.db.AddUser(user)

}

func (tui *TUI) Index() {
  
  var choice int = -1

  log.Printf("Please chose right option: \n 0) for exit \n 1) for sign-up \n 2) to list users \n")
  fmt.Scan(&choice)
  
  if choice == 1 {
    tui.AddUser()   
  } else if choice == 0 {
    tui.db.Close()
  }

}
