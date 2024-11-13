package services

import (
	"database/sql"

	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/repository"
	"github.com/surajNirala/job_portal_api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user *models.User) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	return repository.CreateUser(db, user)

}

func LoginUserService(db *sql.DB, username string, password string) (string, error) {
	user, err := repository.GetUserByUserNameRepository(db, username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "Password is not correct", err
	}
	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}

func ForgotPasswordService(db *sql.DB, username string) (string, error) {

	user, err := repository.GetUserByUserNameRepository(db, username)
	if err != nil {
		return "", err
	}
	randomPassword := utils.GenerateFromPassword(6)
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hasedPassword)
	err = repository.UpdateUserPasswordRepository(db, user)
	if err != nil {
		return "", err
	}
	return randomPassword, nil

}
