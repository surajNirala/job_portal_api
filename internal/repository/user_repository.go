package repository

import (
	"database/sql"
	"fmt"

	"github.com/surajNirala/job_portal_api/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`insert into users (username,email,password) values(?,?,?)`, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByIdRepository(db *sql.DB, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString // Use sql.Nullstring to handle NULL Values
	err := db.QueryRow(`Select * from users where id = ?`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}
	return &user, nil
}

func GetUserByUserNameRepository(db *sql.DB, username string) (*models.User, error) {
	var user models.User
	err := db.QueryRow(`Select * from users where username = ?`, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserProfileRepository(db *sql.DB, user *models.User) (*models.User, error) {
	_, err := db.Exec(`Update users set username = ?, email = ?  where id = ?`, user.Username, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserProfilePictureRepository(db *sql.DB, id int, profilePicture string) error {
	_, err := db.Exec(`Update users set profile_picture = ?  where id = ?`, profilePicture, id)
	if err != nil {
		return err
	}
	return nil
}

func GetUserListRepository(db *sql.DB) ([]*models.User, error) {
	var users []*models.User
	rows, err := db.Query(`Select * from users order by created_at desc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var profilePicture sql.NullString // Use sql.NullString to handle NULL values

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
		if err != nil {
			return nil, err
		}

		if profilePicture.Valid {
			user.ProfilePicture = &profilePicture.String
		} else {
			user.ProfilePicture = nil
		}
		users = append(users, &user)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return users, nil
}

func UpdateUserPasswordRepository(db *sql.DB, user *models.User) error {
	_, err := db.Exec(`Update users set password = ?  where id = ?`, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByIdRepository(tx *sql.Tx, userID int) (string, error) {
	_, err := tx.Exec(`Delete from jobs where user_id = ?`, userID)
	if err != nil {
		return "", fmt.Errorf("error deleting user's jobs: %v", err)
	}

	// Get user's profile picture before deleting user
	var profilePicture sql.NullString // Use sql.Nullstring to handle NULL Values
	err = tx.QueryRow("Select profile_picture from users where id = ?", userID).Scan(&profilePicture)
	if err != nil {
		return "", fmt.Errorf("error fetching user's profile picture: %v", err)
	}

	//Delete the user
	result, err := tx.Exec(`Delete from users where id = ?`, userID)
	if err != nil {
		return "", fmt.Errorf("error deleting user: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}
	return profilePicture.String, nil

}

// func ChangePasswordRepository(db *sql.DB, userID int, OldPassword string, NewPassword string) (string, error) {
// 	user, err := GetUserByIdRepository(db, userID)
// 	if err != nil {
// 		return "", err
// 	}
// 	user.Password = NewPassword
// 	UpdateUserPasswordRepository(db, user)
// }
