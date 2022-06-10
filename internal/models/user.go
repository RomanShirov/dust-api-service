package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Information, written when creating a user
type UserRequest struct {
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	Role         string `json:"role" bson:"role"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
}

// Model for serialize change_role JSON request
type ChangeRoleRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

// User information, used for verification in db.users
type UserResponse struct {
	Username     string `bson:"username"`
	PasswordHash string `bson:"password_hash"`
	Role         string `bson:"role"`
}

func MakeUser(username string, password string) *UserRequest {
	user := new(UserRequest)
	user.Username = username
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	user.CreatedAt = time.Now().String()
	return user
}
