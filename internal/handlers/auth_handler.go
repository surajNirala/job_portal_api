package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/services"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		token, err := services.LoginUserService(db, user.Username, user.Password)

		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Invalid Credentials",
				"issue": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		//Given Output in JSON Format
		response := gin.H{
			"code":    http.StatusOK,
			"message": "Login Successfully.",
			"token":   token,
		}
		c.JSON(http.StatusOK, response)
	}
}

func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		err := services.RegisterUser(db, &user)
		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error creating user.",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := gin.H{
			"code":    http.StatusCreated,
			"message": "User created successfully.",
		}
		c.JSON(http.StatusCreated, response)

	}
}

func ForgotPasswordHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ForgotPasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		password, err := services.ForgotPasswordService(db, req.Username)
		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":             http.StatusCreated,
			"message":          "Password reset successfully.",
			"updated_password": password,
		}
		c.JSON(http.StatusOK, response)
	}
}

func ChangePasswordHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ChangePassword
		if err := c.ShouldBindJSON(&req); err != nil {
			response := gin.H{
				"code":     http.StatusBadRequest,
				"error":    "Something went wrong...",
				"issuesss": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userID := c.GetInt("userID")
		result, err := services.ChangePasswordService(db, userID, req.OldPassword, req.NewPassword)
		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":    http.StatusOK,
			"message": result,
		}
		c.JSON(http.StatusOK, response)
	}
}
