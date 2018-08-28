package controllers

import(
  "../core"
  "net/http"
  "github.com/gorilla/mux"
  "../models"
  "encoding/json"
  mw "../middleware"
)

type HomeController struct{}

func (controller *HomeController) RegisterHandles(router *mux.Router){
  router.HandleFunc("/dashboard", mw.Admin(DashboardHandler))
  router.HandleFunc("/dashboard/data", mw.Admin(DashboardDataHandler))
}


func DashboardHandler(res http.ResponseWriter, req *http.Request){
  core.View(res, req,"admin/dashboard.html", nil)
}

func DashboardDataHandler(res http.ResponseWriter, req *http.Request){
  orders := models.GetOrdersPerMonth()
  numberOfAdmins, numberOfClients := models.GetNumberOfAdminsAndClients()
  moneyPerMonth := models.GetMoneyPerMOnth()
  usersCities := models.GetUsersPerCity()

  data := map[string]interface{}{
    "orders" : orders,
    "numberOfAdmins" : numberOfAdmins,
    "numberOfClients" : numberOfClients,
    "moneyPerMonth" : moneyPerMonth,
    "usersCities" : usersCities,
  }


  dataBytes, err := json.Marshal(data)

  if err != nil {
    res.WriteHeader(http.StatusNoContent)
    res.Write([]byte("Nothing to show"))
  }else{
    res.WriteHeader(http.StatusOK)
    res.Write(dataBytes)
  }
}
