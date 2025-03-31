package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/tokenizer"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if len(token) < 7 || token[:7] != "Bearer " {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			model.NewErrorApiResponse(
				"",
				"Unauthorized",
				http.StatusUnauthorized,
			),
		)
		return
	}

	token = token[7:]

	tokenizer, err := tokenizer.New()

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			model.NewErrorApiResponse(
				"",
				"Tokenizer unavailable",
				http.StatusUnauthorized,
			),
		)
		return
	}

	tokenData, err := tokenizer.DecodeToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			model.NewErrorApiResponse(
				"",
				"Unauthorized",
				http.StatusUnauthorized,
			),
		)
		return
	}

	context.Set("userId", tokenData.UserId)
	context.Set("userEmail", tokenData.UserEmail)

	context.Next()
}
