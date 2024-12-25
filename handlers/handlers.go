package handlers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func PingHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func UserHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "List-user",
  })
}

func ProductHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "List-product",
  })
}

func HelloHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "Hello Thanh",
  })
}
