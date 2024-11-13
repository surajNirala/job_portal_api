package models

type User struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	IsAdmin        bool    `json:"is_admin"`
	ProfilePicture *string `json:"profile_picture"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
