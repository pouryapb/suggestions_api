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

	server.POST("/suggestions", middlewares.Authenticate, createSuggestion)
	server.GET("/suggestions", getSuggestions)
}

func getStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
