package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func InitDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(err)
	}
	usersDb = client.Database("dust").Collection("users")
	charactersDb = client.Database("dust").Collection("characters")
}
