package main

import (
	"rest-api/config"
	"rest-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	db := config.SetupModels()

	// Provide db to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.InitializemerchantRoutes(r)

	// Run the server
	r.Run(":3000")
}
