package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/surajNirala/job_portal_api/internal/models"
	"github.com/surajNirala/job_portal_api/internal/repository"
)

func CreateJobService(db *sql.DB, job *models.Job) (*models.Job, error) {
	return repository.CreateJobRepository(db, job)
}

func GetJobByIdService(db *sql.DB, id int) (*models.Job, error) {
	return repository.GetJobByIdRepository(db, id)
}

func GetJobListService(db *sql.DB) ([]*models.Job, error) {
	return repository.GetJobListRepository(db)
}

func GetAllJobByUserService(db *sql.DB, userID int) ([]*models.Job, error) {
	return repository.GetAllJobByUserRepository(db, userID)
}

func UpdateJobByIdService(db *sql.DB, job *models.Job, userID int, isAdmin bool) (*models.Job, error) {
	existingJob, err := repository.GetJobByIdRepository(db, job.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println("existingJob.UserID : ", existingJob.UserID)
	fmt.Println("existingJob.userID : ", userID)
	if !isAdmin && existingJob.UserID != userID {
		return nil, errors.New("unauthorized to update this job")
	}
	return repository.UpdateJobByIdRepository(db, job)
}

func DeleteJobByIdService(db *sql.DB, id int, userID int, isAdmin bool) error {
	existingJob, err := repository.GetJobByIdRepository(db, id)
	if err != nil {
		return err
	}
	if !isAdmin && existingJob.UserID != userID {
		return errors.New("unauthorized to update this job")
	}
	return repository.DeleteJobByIdRepository(db, id)
}
