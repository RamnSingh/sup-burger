package viewsmodels

import(
  //"../models"
)

type Dashboard struct {
  Orders []OrderData `json:"orders"`
  NumberOfAdmins int `json:"numberOfAdmins"`
  NumberOfClients int `json:"numberOfClients"`
  MoneyPerMonth []MoneyData `json:"moneyPerMonth"`
  UsersCities []UserCityData `json:"usersCities"`
  // RecentBurgers []models.burger
}

type OrderData struct {
  Month int `json:"month"`
  Orders int `json:"orders"`
}

type MoneyData struct {
  Month int `json:"month"`
  Amount float64 `json:"amount"`
}

type UserCityData struct {
  NumberOfUser int `json:"numberOfUser"`
  City string `json:"city"`
}
