package models

import (
	"fmt"
	"time"
)

type CharacterRequest struct {
	Title       string      `json:"title"`
	Description interface{} `json:"description"`
}

type CharacterData struct {
	Username string `json:"username"`
	CharacterRequest
}

type Character struct {
	UploadBy    string      `json:"upload_by" bson:"upload_by"`
	Title       string      `json:"title" bson:"title"`
	Description interface{} `json:"description" bson:"description"`
	ImagePath   string      `json:"image_path" bson:"image_path"`
	SkinPath    string      `json:"skin_path" bson:"skin_path"`
	MusicPath   string      `json:"music_path" bson:"music_path"`
	CreatedAt   string      `json:"created_at" bson:"created_at"`
	IsModerated bool        `json:"is_moderated" bson:"is_moderated"`
	ModeratedBy string      `json:"moderated_by" bson:"moderated_by"`
}

type RemoveCharacterRequest struct {
	Username string `json:"username"`
	Title    string `json:"title"`
}

func CreateCharacter(username string, title string, description interface{}) *Character {
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
