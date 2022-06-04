package db

import (
	"context"
	"dust-api-service/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserResult struct {
	Username     string `bson:"username"`
	PasswordHash string `bson:"password_hash"`
}

var usersDb *mongo.Collection

func AddUser(username, password string) error {
	user := models.MakeUser(username, password)
	_, err := usersDb.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CheckUserExists(username, password string) bool {
	user := models.MakeUser(username, password)
	var result UserResult
	err := usersDb.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Error(err)
	}
	return true
}

func ValidateUser(username, password string) bool {
	user := models.MakeUser(username, password)
	var result UserResult
	err := usersDb.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	passwordHashIsValid := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(password))
	if err == mongo.ErrNoDocuments {
		return false
	} else {
		if passwordHashIsValid == nil {
			return true
		}
		return false
	}
}
