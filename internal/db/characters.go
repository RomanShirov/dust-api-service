package db

import (
	"context"
	"dust-api-service/internal/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var charactersDb *mongo.Collection

func AddCharacter(username, title string, description interface{}) error {
	user := models.CreateCharacter(username, title, description)
	_, err := charactersDb.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func UpdateCharacter(username, title string, description interface{}) error {
	result, err := charactersDb.UpdateOne(context.TODO(), bson.D{{"upload_by", username}, {"title", title}},
		bson.D{
			{"$set", bson.D{{"description", description}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

func GetAllCharacters() []bson.M {
	cursor, err := charactersDb.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var characters []bson.M
	if err = cursor.All(context.TODO(), &characters); err != nil {
		log.Fatal(err)
	}
	return characters
}

func GetAllUserCharacters(username string) []bson.M {
	cursor, err := charactersDb.Find(context.TODO(), bson.D{{"upload_by", username}})
	if err != nil {
		log.Fatal(err)
	}
	var characters []bson.M
	if err = cursor.All(context.TODO(), &characters); err != nil {
		log.Fatal(err)
	}
	return characters
}

func RemoveCharacter(username, title string) error {
	_, err := charactersDb.DeleteOne(context.TODO(), bson.D{{"upload_by", username}, {"title", title}})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
