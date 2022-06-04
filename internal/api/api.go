package api

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/tokens"
)

func CreateUser(username, password string) (token string, err error) {
	isUserExists := db.CheckUserExists(username, password)
	if isUserExists {
		err = db.AddUser(username, password)
		if err != nil {
			return "", err
		}
		return tokens.GenerateUserToken(username), nil
	} else {
		return "", err
	}

}
