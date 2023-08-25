package middleware

import (
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/utility"
	"github.com/gin-gonic/gin"
)

func JWTAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt token from header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		// remove bearer
		if len(token) < 7 || token[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		token = token[7:]

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		// verify token
		userID, ok := utility.ParseToken(token)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		// set userID to context
		c.Set("userID", userID)
		c.Next()
	}
}
