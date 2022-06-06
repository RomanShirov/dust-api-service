package api

import (
	"dust-api-service/internal/db"
	"dust-api-service/internal/models"
	"dust-api-service/internal/tokens"
	"errors"
	log "github.com/sirupsen/logrus"
)

func CreateUser(username, password string) (token string, err error) {
	isUserExists := db.CheckUserExists(username, password)
	if isUserExists {
		return "", errors.New("user already exists")
	} else {
		err = db.AddUser(username, password)
		if err != nil {
			return "", err
		}
		return tokens.GenerateUserToken(username), nil
	}
}

func CreateCharacter(character models.CharacterData) error {
	username := character.Username
	title := character.Title
	description := character.Description
	err := db.AddCharacter(username, title, description)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func RemoveCharacter(username, title string) error {
	err := db.RemoveCharacter(username, title)
	if err != nil {
		log.Error(err)
		return err
	} else {
		return nil
	}
}
