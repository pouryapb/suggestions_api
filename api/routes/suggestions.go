package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"suggestions.api/models"
)

func createSuggestion(ctx *gin.Context) {
	var suggestion models.Suggestion

	err := ctx.ShouldBind(&suggestion)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "no text provided"})
		return
	}

	username := ctx.GetString("username")
	suggestion.Username = username

	err = suggestion.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "no text provided"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"ok": true})
}

func getSuggestions(ctx *gin.Context) {
	suggestions, err := models.GetAllSuggestions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch suggestions."})
		return
	}

	ctx.JSON(http.StatusOK, suggestions)
}
