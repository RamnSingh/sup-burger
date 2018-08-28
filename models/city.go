package models

import(
  "../db"
)


type City struct {
  ID int `json:"id"`
  Name string `json:"name" validate:"required"`
}

func GetAllCities() ([]City, error){
  rows, err := db.Select("SELECT * FROM city")
  cities := make([]City,0)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err !=  nil {
      return nil, err
    }
    dbCity := City{id, name}
    cities = append(cities, dbCity)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return cities, nil
}

func (city *City) Add() (error){
  query := "INSERT INTO city (name) VALUES (?)"

  args := []interface{}{city.Name}

  res, err := db.Insert(query, args...)

  if err != nil {
    return err
  }
  int64Id, err := res.LastInsertId()
  if err == nil {
    city.ID = int(int64Id)
  }
  return nil
}

func (city *City) Update() error {
  query := "UPDATE city SET name = ? where id = ?"
  args := []interface{}{city.Name, city.ID}
  _, err := db.Update(query, args...)

  if err != nil {
    return err
  }
  return nil
}

func (city *City) Delete() ( error){
  var err error
  query := "set foreign_key_checks = 0"
  if err = db.Exec(query); err == nil {
    query = "DELETE FROM city WHERE id = ?"
    args := []interface{}{city.ID}
    _, err = db.Delete(query, args...)
  }
  return err
}
