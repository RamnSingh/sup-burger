package config

import(
  "net/http"
  "../core"
  ctrl "../controllers"
  "github.com/gorilla/mux"
)

var router *mux.Router

func InitApp(){
  router = mux.NewRouter()
}

func RegisterControllers() {
  controller(new(ctrl.HomeController))
  controller(new(ctrl.UserController))
}

func controller (controller core.IController) {
  controller.RegisterHandles(router);
}

func StartApp(at string){
  http.ListenAndServe(at, router)
}
