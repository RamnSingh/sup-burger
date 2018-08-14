package controllers

import(
  _ "../core"
  "net/http"
  "github.com/gorilla/mux"
)

type HomeController struct{}

func (controller *HomeController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/", HomePageHandler)
}


func HomePageHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Hello"))
}
