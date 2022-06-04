package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	Role         string `json:"role" bson:"role"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
}

func MakeUser(username string, password string) *User {
	user := new(User)
	user.Username = username
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	user.CreatedAt = time.Now().String()
	return user
}
