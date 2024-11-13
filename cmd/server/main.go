package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surajNirala/job_portal_api/internal/repository"
	"github.com/surajNirala/job_portal_api/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("Database error: ", err)
	}
	defer db.Close()
	r := gin.Default()
	routes.InitRoutes(r, db)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
