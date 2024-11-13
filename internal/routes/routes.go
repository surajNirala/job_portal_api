package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/auth"
	handler "github.com/surajNirala/job_portal_api/internal/handlers"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	// AUTH ROUTES
	r.POST("/login", handler.LoginHandler(db))
	r.POST("/register", handler.RegisterHandler(db))
	r.GET("/all-jobs", handler.GetJobListHandler(db))
	r.POST("/forgotpassword", handler.ForgotPasswordHandler(db))

	//User Routes
	authenticated := r.Group("/")
	authenticated.Use(auth.AuthMiddlware())
	authenticated.GET("/users", handler.GetUserListHandler(db))
	authenticated.GET("/users/:id", handler.GetUserByIdHandler(db))
	authenticated.PUT("/users/:id", handler.UpdateUserProfileHandler(db))
	authenticated.POST("/users/:id/profile-picture", handler.UpdateUserProfilePictureHandler(db))

	//Job Routes
	authenticated.GET("/jobs-by-user", handler.GetAllJobByUserHandler(db))
	authenticated.POST("/jobs", handler.CreateJobHandler(db))
	authenticated.GET("/jobs/:id", handler.GetJobByIdHandler(db))
	authenticated.PUT("/jobs/:id", handler.UpdateJobByIdHandler(db))
	authenticated.DELETE("/jobs/:id", handler.DeleteJobByIdHandler(db))

}
