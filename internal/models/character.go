package models

import (
	"fmt"
	"time"
)

type Character struct {
	UploadBy    string `json:"upload_by" bson:"upload_by"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	SkinPath    string `json:"skin_path" bson:"skin_path"`
	MusicPath   string `json:"music_path" bson:"music_path"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	IsModerated bool   `json:"is_moderated" bson:"is_moderated"`
	ModeratedBy string `json:"moderated_by" bson:"moderated_by"`
}

func CreateCharacter(username, title, description string) *Character {
	character := new(Character)
	character.UploadBy = username
	character.Title = title
	character.Description = description
	character.SkinPath = fmt.Sprintf("skins/%s/%s/skin.png", username, title)
	character.MusicPath = fmt.Sprintf("music/%s/%s/music.mp3", username, title)
	character.CreatedAt = time.Now().String()
	character.IsModerated = false
	character.ModeratedBy = ""
	return character
}
