package repository

import (
	"database/sql"

	"github.com/surajNirala/job_portal_api/internal/models"
)

func CreateJobRepository(db *sql.DB, job *models.Job) (*models.Job, error) {
	result, err := db.Exec("insert into jobs (title, description, company, location, salary, user_id) values (?, ?, ?, ?, ?, ?)", job.Title, job.Description, job.Company, job.Location, job.Salary, job.UserID)

	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	job.ID = int(id)
	return job, nil
}

func GetJobByIdRepository(db *sql.DB, id int) (*models.Job, error) {
	var job models.Job
	err := db.QueryRow(`Select id, title, description, company, location, salary, user_id, created_at, updated_at from jobs where id = ?`, id).Scan(&job.ID, &job.Title, &job.Description, &job.Company, &job.Location, &job.Salary, &job.UserID, &job.CreatedAt, &job.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func GetJobListRepository(db *sql.DB) ([]*models.Job, error) {
	var jobs []*models.Job
	rows, err := db.Query(`Select id, title, description, company, location, salary, user_id, created_at, updated_at from jobs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Company, &job.Location, &job.Salary, &job.UserID, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return jobs, nil
}

func GetAllJobByUserRepository(db *sql.DB, userID int) ([]*models.Job, error) {
	var jobs []*models.Job
	rows, err := db.Query(`Select id, title, description, company, location, salary, user_id, created_at, updated_at from jobs where user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Company, &job.Location, &job.Salary, &job.UserID, &job.CreatedAt, &job.UpdatedAt)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return jobs, nil
}

func UpdateJobByIdRepository(db *sql.DB, job *models.Job) (*models.Job, error) {
	_, err := db.Exec("update jobs set title = ?, description = ?, company = ?, location = ?, salary = ? where id = ?", job.Title, job.Description, job.Company, job.Location, job.Salary, job.ID)

	if err != nil {
		return nil, err
	}
	return job, nil
}

func DeleteJobByIdRepository(db *sql.DB, id int) error {
	_, err := db.Exec("delete from jobs where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
