package config

import(
  "./models"
  "./controllers"
  "github.com/gorilla/mux"
)

var router *mux.Router

func InitApp(){
  router = mux.NewRouter()
}

func RegisterController() {
  controller(&HomeController)
}

func controller (controller *controller) {
  controller.init(&router);
}

func StartApp(at string){
  http.ListenAndServe(at, router)
}
