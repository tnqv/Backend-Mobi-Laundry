package store

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Longitude      float32		`json:"longitude"`
	Latitude 	   float32		`json:"latitude"`
	Address        string   	`json:"address"`
	PhoneNumber    string  		`json:"phone_number"`
}