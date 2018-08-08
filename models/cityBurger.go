package models

type CityBurger struct{
  City City `json:"city"`
  Burger Burger `json:"burger"`
}
