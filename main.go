package main

import (
  "os"
  "fmt"
  "./models"
)

func main(){
  burger := new(models.Burger)
  burger.ID = 5
  fmt.Println(os.Args[1])
  fmt.Println(burger)
}
