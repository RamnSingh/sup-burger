package controllers

import(
  core "../core"
  "net/http"
  "github.com/gorilla/mux"
)

type HomeController struct{}

func (controller *HomeController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/", HomePageHandler)
}


func HomePageHandler(res http.ResponseWriter, req *http.Request){
  core.View(res, "account/index.html", nil)
}
