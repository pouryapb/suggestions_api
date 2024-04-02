package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"suggestions.api/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized."})
		return
	}

	username, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized."})
		return
	}

	ctx.Set("username", username)
	ctx.Next()
}
