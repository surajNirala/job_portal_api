package services

import (
	"database/sql"
	"fmt"

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

func ChangePasswordService(db *sql.DB, userID int, OldPassword string, NewPassword string) (string, error) {
	user, err := repository.GetUserByIdRepository(db, userID)
	if err != nil {
		return "", fmt.Errorf("user not found.")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(OldPassword)); err != nil {
		return "", fmt.Errorf("Old Password is not correct.")
	}
	hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashNewPassword)
	err = repository.UpdateUserPasswordRepository(db, user)
	if err != nil {
		return "", err
	}
	return "Password updated successfully.", nil
	// response, err := repository.ChangePasswordRepository(db, userID, OldPassword, hashNewPassword)
	// if err != nil {
	// 	return "", err
	// }
	// return response, nil
}
