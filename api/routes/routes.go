package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"suggestions.api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", getStatus)

	server.POST("/signup", signup)
	server.POST("/login", login)

	suggestions := server.Group("/suggestions")
	suggestions.Use(middlewares.Authenticate)

	suggestions.POST("/", createSuggestion)
	suggestions.GET("/", getSuggestions)
}

func getStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
