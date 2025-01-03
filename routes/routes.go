package routes

import (
  "github.com/gin-gonic/gin"
  "myproject/handlers"
)

func SetupRoutes(r *gin.Engine) {
  r.GET("/ping", handlers.PingHandler)
  r.GET("/user", handlers.UserHandler)
  r.GET("/product", handlers.ProductHandler)
  r.GET("/hello", handlers.HelloHandler)
}


