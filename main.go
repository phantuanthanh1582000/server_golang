package main

import (
  "github.com/gin-gonic/gin"
  "myproject/routes"
)

func main() {
  r := gin.Default()
  
  routes.SetupRoutes(r)

  r.Run(":3000")
}
