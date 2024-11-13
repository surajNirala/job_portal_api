package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/pkg/utils"
)

func AuthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response := gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Missing Authorization Header...",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(token)
		if err != nil {
			response := gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token.",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("isAdmin", claims.IsAdmin)
		c.Next()

	}
}
