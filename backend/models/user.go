package models

import (
	"github.com/hyhkjiy/devin_test/backend/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  UserResponse `json:"user"`
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	var passwordHash string
	query := `SELECT id, username, password_hash FROM users WHERE username = ?`
	err := database.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &passwordHash)
	user.Password = passwordHash
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
	}
}

func (u *User) Create() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password_hash) VALUES (?, ?)`
	result, err := database.DB.Exec(query, u.Username, string(hashedPassword))
	if err != nil {
		return err
	}

	u.ID, _ = result.LastInsertId()
	return nil
}
