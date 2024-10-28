package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/cutomers", func(c *gin.Context) {})
	r.GET("/customers/:customerId", func(c *gin.Context) {})
	r.PUT("/customers/:customerId", func(c *gin.Context) {})
	r.DELETE("/customers/:customerId", func(c *gin.Context) {})

	r.Run("localhost:8080")
}
