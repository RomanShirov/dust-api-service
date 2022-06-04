package db

import (
	"context"
	"dust-api-service/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var usersDb *mongo.Collection
var charactersDb *mongo.Collection
var ctx = context.TODO()

func InitDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(err)
	}
	usersDb = client.Database("dust").Collection("users")
	charactersDb = client.Database("dust").Collection("characters")
}

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
	var result bson.D
	err := usersDb.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Error(err)
	}
	return true
}
