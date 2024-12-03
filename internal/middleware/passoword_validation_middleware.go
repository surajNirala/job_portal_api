package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/pkg/utils"
)

func PasswordValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response := gin.H{
				"code":    http.StatusBadRequest,
				"message": "Error reading request body.",
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}
		// Create new reader with the bytes
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		//Parse the request
		var req models.ChangePassword
		if err := c.ShouldBindJSON(&req); err != nil {
			response := gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid body request.",
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		// Restore the request body for the next middleware/handler
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		isValid, errors := utils.ValidatePasswordStrength(req.NewPassword)
		if !isValid {
			response := gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Password validation failed.",
				"details": errors,
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}
		c.Next()
	}
}
