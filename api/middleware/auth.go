package middleware

import (
	"nexora_backend/pkg/utils"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Next() 
	}
}
