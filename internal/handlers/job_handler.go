package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/services"
)

func CreateJobHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var job models.Job
		if err := c.ShouldBindJSON(&job); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		userID := c.GetInt("userID")
		job.UserID = userID
		createJob, err := services.CreateJobService(db, &job)
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
			"message": "Job created successfully.",
			"data":    createJob,
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetJobByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid Job ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		job, err := services.GetJobByIdService(db, id)
		if err != nil {
			response := gin.H{
				"code":  http.StatusOK,
				"error": "Job not found",
				"issue": err.Error(),
			}
			c.JSON(http.StatusOK, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "Job details found.",
			"data":    job,
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetJobListHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		jobs, err := services.GetJobListService(db)
		if err != nil {
			response := gin.H{
				"code":  http.StatusOK,
				"error": "Jobs list not found",
				"issue": err.Error(),
			}
			c.JSON(http.StatusOK, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "Fetched Jobs List found.",
			"data":    jobs,
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetAllJobByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.GetInt("userID")
		jobs, err := services.GetAllJobByUserService(db, user_id)
		if err != nil {
			response := gin.H{
				"code":  http.StatusOK,
				"error": "Jobs list not found",
				"issue": err.Error(),
			}
			c.JSON(http.StatusOK, response)
			return
		}

		response := gin.H{
			"code":    http.StatusOK,
			"message": "Fetched Jobs List found.",
			"data":    jobs,
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateJobByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid Job ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		var job models.Job
		job.ID = id
		if err := c.ShouldBindJSON(&job); err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")
		updateJob, err := services.UpdateJobByIdService(db, &job, userID, isAdmin)
		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Job is not updated successfully.",
				"issue": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":    http.StatusOK,
			"message": "Job is updated successfully.",
			"data":    updateJob,
		}
		c.JSON(http.StatusOK, response)
	}
}

func DeleteJobByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response := gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid Job ID",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")
		err = services.DeleteJobByIdService(db, id, userID, isAdmin)
		if err != nil {
			response := gin.H{
				"code":  http.StatusInternalServerError,
				"error": "Job is not deleted successfully.",
				"issue": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := gin.H{
			"code":    http.StatusOK,
			"message": "Job is delete successfully.",
		}
		c.JSON(http.StatusOK, response)
	}
}
