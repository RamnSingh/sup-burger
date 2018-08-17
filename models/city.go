package models

import(
  "../db"
)


type City struct {
  ID int64 `json:"id"`
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
    var id int64
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
