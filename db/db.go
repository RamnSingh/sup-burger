package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type DB struct {
  DB *sql.DB
}

var db DB

func InitDatabase() {
  if db.DB == nil {
    database, err := sql.Open("mysql", "dodo:password@tcp(127.0.0.1:3306)/sup_burger?parseTime=true")
    if err != nil{
      panic(err.Error())
    }
    err = database.Ping()
    if err != nil{
      panic(err.Error())
    }
    db.DB = database
  }
}

func Select (query string, args ...interface{})(*sql.Rows, error) {
  return db.DB.Query(query, args...)
}

func SelectRow (query string, args ...interface{})(*sql.Row) {
  return db.DB.QueryRow(query, args...)
}

func Insert (query string, args ...interface{})(sql.Result, error) {
  return db.DB.Exec(query, args...)
}

func Update (query string, args ...interface{})(sql.Result, error) {
  return db.DB.Exec(query, args...)
}

func Delete (query string, args ...interface{})(sql.Result, error) {
  return db.DB.Exec(query, args...)
}

func Exec(query string, args ...interface{}) error {
  _, err := db.DB.Exec(query, args...)
  return err
}
