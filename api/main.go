package main

import (
	"github.com/gin-gonic/gin"
	"suggestions.api/db"
	"suggestions.api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run("0.0.0.0:80")
}
