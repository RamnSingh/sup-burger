package models

import(
  "../db"
)

type Stuff struct{
  ID int `json:"id"`
  Name string `json:"name" validate:"required"`
}

func GetAllStuff() ([]Stuff, error){
  rows, err := db.Select("SELECT id, name FROM stuff")
  stuffs := make([]Stuff,0)
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
    dbStuff := Stuff{
      ID : id,
      Name : name,
    }
    stuffs = append(stuffs, dbStuff)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return stuffs, nil
}


func (stuff *Stuff) Add() (error){
  query := "INSERT INTO stuff (name) VALUES (?)"

  args := []interface{}{stuff.Name}

  res, err := db.Insert(query, args...)

  if err != nil {
    return err
  }
  int64Id, err := res.LastInsertId()
  if err == nil {
    stuff.ID = int(int64Id)
  }
  return nil
}

func (stuff *Stuff) Update() error {
  query := "UPDATE stuff SET name = ? where id = ?"
  args := []interface{}{stuff.Name, stuff.ID}
  _, err := db.Update(query, args...)

  if err != nil {
    return err
  }
  return nil
}

func (stuff *Stuff) Delete() ( error){
  var err error
  query := "set foreign_key_checks = 0"
  if err = db.Exec(query); err == nil {
    query = "DELETE FROM stuff WHERE id = ?"
    args := []interface{}{stuff.ID}
    _, err = db.Delete(query, args...)
  }
  return err
}
