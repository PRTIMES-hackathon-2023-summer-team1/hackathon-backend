package middleware

import (
	"log"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/controllers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		log.Println(err)
		if err != nil {
			apierror := err.Meta.(controllers.APIError)
			c.AbortWithStatusJSON(apierror.StatusCode, gin.H{
				"error": apierror.ErrorMessage,
			})

		}

	}
}
