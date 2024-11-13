package repository

import (
	"database/sql"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", os.Getenv("DB_PATH"))
	if err != nil {
		return nil, err
	}

	if err = UserTable(db); err != nil {
		return nil, fmt.Errorf("error creating users table: %w", err)
	}
	if err = UserInsertAdmin(db); err != nil {
		return nil, fmt.Errorf("error insert admin in users table: %w", err)
	}

	if err = JobTable(db); err != nil {
		return nil, fmt.Errorf("error creating jobs table: %w", err)
	}

	return db, nil
}

func UserTable(db *sql.DB) error {
	_, err := db.Exec(`create table if not exists users(
		id integer primary key autoincrement,
		username text not null unique,
		email text not null,
		password text not null,
		created_at datetime default current_timestamp,
		updated_at datetime default current_timestamp,
		is_admin boolean default 0,
		profile_picture text 
	)`)

	return err
}

func UserInsertAdmin(db *sql.DB) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte("admin@123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	bcrypt_hasedPassword := string(hasedPassword)
	_, err = db.Exec(`insert or ignore into users (username,email,password,is_admin) values (?,?,?,?)`, "admin", "admin@gmail.com", bcrypt_hasedPassword, true)
	return err
}

func JobTable(db *sql.DB) error {
	_, err := db.Exec(`create table if not exists jobs(
		id integer primary key autoincrement,
		title text not null,
		description text not null,
		company text not null,
		location text not null,
		salary text not null,
		user_id integer not null,
		created_at datetime default current_timestamp,
		updated_at datetime default current_timestamp,
		foreign key (user_id) references users(id)
	)`)
	return err
}
