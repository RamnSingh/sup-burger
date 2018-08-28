package models

import(
  "../db"
  "fmt"
)

type BurgerStuff struct{
  BurgerId int
  StuffId int
  Burger Burger `json:"burger"`
  Stuff Stuff `json:"stuff"`
}

type BurgerStuffArr [] BurgerStuff

func (burgerStuffArr BurgerStuffArr)AddN () error {

  query := "INSERT INTO burger_stuff (burger_id, stuff_id) VALUES "

  args := make([]interface{}, 0)

  for _, bs := range burgerStuffArr {
    query += "(?, ?),"
    args = append(args, bs.Burger.ID, bs.Stuff.ID)
  }

  query = query[0:len(query)-1]
  fmt.Println(query)
  _, err := db.Insert(query, args...)

  if err != nil {
    return err
  }
  return nil
}

func GetBurgerStuff() ([]BurgerStuff, error){
  rows, err := db.Select("SELECT * FROM burger_stuff")
  burgerStuff := make([]BurgerStuff,0)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  for rows.Next() {
    var burger_id, stuff_id int
    if err := rows.Scan(&burger_id, &stuff_id); err !=  nil {
      return nil, err
    }
    dbBurgerStuff := BurgerStuff {
      BurgerId : burger_id,
      StuffId : stuff_id,
    }
    burgerStuff = append(burgerStuff, dbBurgerStuff)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return burgerStuff, nil
}

func (burgerStuff *BurgerStuff) Add() error {
  query := "INSERT INTO burger_stuff (burger_id, stuff_id) VALUES (?, ?)"
  args := []interface{}{burgerStuff.BurgerId, burgerStuff.StuffId}

  _, err := db.Insert(query, args...)

  return err
}

func (burgerStuff *BurgerStuff) Delete() error {
  query := "DELETE FROM burger_stuff where burger_id = ? AND stuff_id = ?"
  args := []interface{}{burgerStuff.BurgerId, burgerStuff.StuffId}

  _, err := db.Delete(query, args...)

  return err
}
