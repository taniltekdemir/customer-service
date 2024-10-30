package main

import (
	"github.com/gin-gonic/gin"

	"customer-service/db"
	"customer-service/service"
)

func main() {

	secret := db.GetSecretValue()
	db := db.GetDB(secret)
	a := service.GetApp(db)

	r := gin.Default()

	r.POST("/cutomers", a.PostHandler)
	r.GET("/customers/:customerId", a.GetHandler)
	r.PUT("/customers/:customerId", a.PutHandler)
	r.DELETE("/customers/:customerId", a.DeleteHandler)

	r.Run("localhost:8080")
}
