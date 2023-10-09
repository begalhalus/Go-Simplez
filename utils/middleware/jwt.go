package middleware

import (
	"go-simple/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	berarerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(berarerToken, "Bearer") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	token := strings.Replace(berarerToken, "Bearer ", "", 1)

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	credential, err := helper.DecodeToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid Token",
		})
		return
	}

	ctx.Set("credential", credential)

	ctx.Next()
}
