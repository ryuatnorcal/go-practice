package middlewares

import (
	"net/http"

	"example.com/final_project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// you can attach custom value in context
	context.Set("userId", userId)

	// next request handle will get called
	context.Next()
}
