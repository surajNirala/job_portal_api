package models

import "time"

type Job struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Company     string    `json:"company"`
	Salary      string    `json:"salary"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
