package web 

import (
	"net/http"
	"shiro/web/view"
)

func Render(name string) func(w http.ResponseWriter, r *http.Request) {

  return func(w http.ResponseWriter, r *http.Request) {
    view.RenderView(w, name, "base", nil)
  }

}
