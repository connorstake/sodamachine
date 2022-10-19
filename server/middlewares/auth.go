package middlewares

import (
	"github.com/mvpmatch/server/auth"

	"github.com/gin-gonic/gin"
)

// Auth checks the validity of a users access token and passes the claims in the context
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.AddParam("username", claims.Username)
		context.Next()
	}
}
