package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	tokenizer, err := service.NewTokenizer()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "tokenizer unavailable"})
		return
	}

	tokenData, err := tokenizer.DecodeToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	context.Set("userId", tokenData.UserId)

	context.Next()
}
