package main

import (
  "./config"
)

func main(){
  config.InitApp()
  config.RegisterControllers()
  config.StartApp(":8080")
}
