package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/services"
)

func GetUserListHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Unautorized to fetch all user list.",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		user, err := services.GetUserListService(db)
		if err != nil {
			response := gin.H{
				"code":  http.StatusOK,
				"error": "Users not found",
				"issue": err.Error(),
			}
			c.JSON(http.StatusOK, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "Fetched User list.",
			"data":    user,
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid User ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		user, err := services.GetUserByIdService(db, id)
		if err != nil {
			response := gin.H{
				"code":  http.StatusOK,
				"error": "User not found",
				"issue": err.Error(),
			}
			c.JSON(http.StatusOK, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "User details found.",
			"data":    user,
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateUserProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid User ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}
		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Something went wrong.",
				"issue": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// username := c.GetInt("username")
		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")
		// fmt.Println("username", username)
		// fmt.Println("userID", userID)
		// fmt.Println("id", id)
		if !isAdmin && userID != id {
			response := gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Unauthorized to update this user profile.",
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		updateUser, err := services.UpdateUserProfileService(db, id, userUpdate.Username, userUpdate.Email)

		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error updating user profile.",
				"issue": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":    http.StatusOK,
			"message": "User Profile updated successfully.",
			"data":    updateUser,
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateUserProfilePictureHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid User ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin && userID != id {
			response := gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Unauthorized to update this user profile.",
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		file, err := c.FormFile("profile_picture")

		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid User ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if err := os.MkdirAll(os.Getenv("UPLOAD_DIR"), os.ModePerm); err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error creating upload directory.",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		filename := fmt.Sprintf("%d-%s", id, filepath.Base(file.Filename))
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error saving uploaded file.",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		err = services.UpdateUserProfilePictureService(db, id, filename)

		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Error updating profile picture in database.",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":    http.StatusOK,
			"message": "Profile picture updated successfully.",
		}
		c.JSON(http.StatusOK, response)
	}
}

func DeleteUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			response := gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Unautorized access.",
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid User ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		currentUserID := c.GetInt("userID")
		if currentUserID == id {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "You cannot delete yourself.",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		err = services.DeleteUserByIdService(c, db, id)
		if err != nil {
			if err.Error() == "User not found." {
				response := gin.H{
					"code":  http.StatusInternalServerError,
					"error": "User not found.",
				}
				c.JSON(http.StatusInternalServerError, response)
				return
			}
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": fmt.Sprintf("Error deleting user: %v", err),
				"issue": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "User and associated data deleted successfully.",
		}
		c.JSON(http.StatusOK, response)
	}
}
