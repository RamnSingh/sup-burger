package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func getDatabase() *sql.DB {

  db, err := sql.Open("mysql", "dodo:password@tcp(127.0.0.1:3306)/test")
  defer db.Close()
  if err != nil{
    panic(err.Error())
  }
  err = db.Ping()
  if err != nil{
    panic(err.Error())
  }

  return db
}

// func init
//
//
//
//
//   rows, err := db.Query("Select * from users")
//
//   if err != nil {
//     panic("rows error")
//   }
//
//   var id int64
//   var username, password string
//
//   for rows.Next() {
//     err := rows.Scan(&id, &username, &password)
//     if err != nil {
//       panic(err.Error())
//     }
//
//     fmt.Println(id, username, password)
//   }
//   err = rows.Err()
//   if err != nil {
//     panic(err.Error())
//   }
//
//   fmt.Println("Everything is ok !")
//
//
// }
