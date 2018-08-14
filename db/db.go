package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func GetDatabase() *sql.DB {
  db, err := sql.Open("mysql", "dodo:password@tcp(127.0.0.1:3306)/test")
  if err != nil{
    panic(err.Error())
  }
  err = db.Ping()
  if err != nil{
    panic(err.Error())
  }
  return db
}

//
// func Populate(database *sql.DB) {
//   GetDatabase()
//   data, err := ioutil.ReadFile("./db/populate.sql")
//
//   if err != nil{
//     panic(err.Error())
//   }
//   dataString := string(data[:])
//
//
//   stmt, err := database.Prepare(dataString)sss
//
//   if err != nil {
//     panic(err.Error())
//   }
//
//   fmt.Println(stmt)
//   res, err := stmt.Exec()
//
//   if err != nil{
//     panic(err.Error())
//   }
//   fmt.Println(res)
// }
