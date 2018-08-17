package controllers

import(
  core "../core"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "../models"
  validation "../validators/form"
  // "errors"
)

type AccountController struct{}

func (controller *AccountController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/account/register", AccountRegisterHandler).Methods("GET", "POST")
  router.HandleFunc("/account/login", AccountLoginHandler).Methods("GET", "POST")
  router.HandleFunc("/account/logout", AccountLogoutHandler).Methods("POST")
}


func AccountRegisterHandler(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
    req.ParseForm()
    var user models.User
    user.PopulateFromForm(*req)
    fmt.Println(user)
    if len(req.Form["cf-password"]) > 0 {
      err := validation.Register(user, req.Form["cf-password"][0])

      if err != nil {
        fmt.Println(err.Error())
      }else{
        err = user.Register(1)
        if err != nil {
          fmt.Println(err.Error())
        }else{
          fmt.Println("Created")
        }
      }
    }else{
      fmt.Println("confirm password not found")
    }
  }else {
    data, err := models.GetAllCities()
    if err != nil {
      core.View(res, "account/index.html", nil)
    } else {
      core.View(res, "account/register.html", data)
    }

  }
}

func AccountLoginHandler(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Login"))
}

func AccountLogoutHandler(res http.ResponseWriter, req *http.Request){
  res.Write([]byte("Logout"))
}
