package models

import(
  "../db"
)

type CityBurger struct{
  CityId int
  BurgerId int
  City City `json:"city"`
  Burger Burger `json:"burger"`
}


type CityBurgerArr [] CityBurger

func (cityBurgerArr CityBurgerArr)AddN () error {

  query := "INSERT INTO city_burger (burger_id, city_id) VALUES "

  args := make([]interface{}, 0)

  for _, bs := range cityBurgerArr {
    query += "(?, ?),"
    args = append(args, bs.Burger.ID, bs.City.ID)
  }

  query = query[0:len(query)-1]

  _, err := db.Insert(query, args...)

  if err != nil {
    return err
  }
  return nil
}

func GetCitiesBurger() ([]CityBurger, error){
  rows, err := db.Select("SELECT * FROM city_burger")
  citiesBurgers := make([]CityBurger,0)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next() {
    var burger_id, city_id int
    if err := rows.Scan(&city_id, &burger_id); err !=  nil {
      return nil, err
    }
    dbCityBurger := CityBurger {
      CityId : city_id,
      BurgerId : burger_id,
    }
    citiesBurgers = append(citiesBurgers, dbCityBurger)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return citiesBurgers, nil
}

func (cityBurger *CityBurger) Add() (error) {
  query := "set foreign_key_checks = 0;INSERT INTO city_burger (city_id, burger_id) VALUES (?, ?);set foreign_key_checks = 1;"
  args := []interface{}{cityBurger.CityId, cityBurger.BurgerId}

  _, err := db.Insert(query, args...)

  return err
}

func (cityBurger *CityBurger) Delete() error {
  query := "DELETE FROM city_burger where city_id = ? AND burger_id = ?"
  args := []interface{}{cityBurger.CityId, cityBurger.BurgerId}

  _, err := db.Delete(query, args...)

  return err
}
