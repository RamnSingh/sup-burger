package config

import(
  "log"
  "net/http"
  "../core"
  ctrl "../controllers"
  "github.com/gorilla/mux"
  "../db"
)

var router *mux.Router

func InitApp(){
  router = mux.NewRouter()
}

func RegisterStaticFiles() {
  router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
}

func RegisterControllers() {
  controller(new(ctrl.HomeController))
  controller(new(ctrl.AccountController))
  controller(new(ctrl.UserController))
  controller(new(ctrl.BurgerController))
  controller(new(ctrl.CityController))
  controller(new(ctrl.StuffController))
  controller(new(ctrl.BasketController))
}

func controller (controller core.IController) {
  controller.RegisterHandles(router);
}

func StartApp(at string){
  db.InitDatabase()
  log.Fatal(http.ListenAndServe(at, router))
}
