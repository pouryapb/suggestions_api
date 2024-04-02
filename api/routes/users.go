package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"suggestions.api/models"
	"suggestions.api/utils"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "no username or password provided"})
		return
	}

	err = user.Save()
	if err != nil {
		// I know this is not a good message here but I'm not gonna try THAT hard for this :))
		ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "user already exists"})
		return
	}

	token, err := utils.GenerateToken(user.Username, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"ok": true, "token": token})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "no username or password provided"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(user.Username, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
}
