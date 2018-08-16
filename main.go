package main

import (
  "./config"
)

func main(){
  config.InitApp()
  config.RegisterStaticFiles()
  config.RegisterControllers()
  config.StartApp(":8080")
}
