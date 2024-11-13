package services

import (
	"database/sql"

	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/repository"
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
