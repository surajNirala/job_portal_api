package services

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/repository"
	"github.com/surajNirala/job_portal_api/pkg/utils"
)

func GetUserByIdService(db *sql.DB, id int) (*models.User, error) {
	return repository.GetUserByIdRepository(db, id)
}

func UpdateUserProfileService(db *sql.DB, id int, username, email string) (*models.User, error) {
	user := models.User{ID: id, Username: username, Email: email}
	return repository.UpdateUserProfileRepository(db, &user)
}

func UpdateUserProfilePictureService(db *sql.DB, id int, profilePicture string) error {
	return repository.UpdateUserProfilePictureRepository(db, id, profilePicture)
}

func GetUserListService(db *sql.DB) ([]*models.User, error) {
	return repository.GetUserListRepository(db)
}

func DeleteUserByIdService(c *gin.Context, db *sql.DB, id int) error {
	// Start a transaction
	tx, err := db.BeginTx(c.Request.Context(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback() // Rollback if not committed

	// Delete user and associated data
	profilePicture, err := repository.DeleteUserByIdRepository(tx, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("User not found.")
		}
		return fmt.Errorf("error deleting user: %v", err)
	}

	// Commit the trasaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	// Delete Profile picture after successful transaction if it exists
	if profilePicture != "" {
		filePathImage := filepath.Join(os.Getenv("UPLOAD_DIR"), profilePicture)
		err = utils.DeleteFileIfExists(filePathImage)
		if err != nil {
			return fmt.Errorf("error deleting profile picture: %v", err)
		}
	}
	return nil
}
