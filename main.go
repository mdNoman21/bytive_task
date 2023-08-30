package main

import (
	"bytive-task/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.New()

	if port == "" {
		port = "8000"
	}
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Run(":" + port)

}
