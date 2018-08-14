package core

import(
  "github.com/gorilla/mux"
)

type IController interface {
  RegisterHandles(router *mux.Router)
}
