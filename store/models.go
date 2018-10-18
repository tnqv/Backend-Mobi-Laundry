package store

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
)

type Store struct {
	gorm.Model					`json:"-"`
	Name           string       `form:"name" json:"name" binding:"exists"`
	Description    string       `form:"description" json:"description" binding:"exists"`
	Longitude      float32		`form:"longitude" json:"longitude" binding:"exists"`
	Latitude 	   float32		`form:"latitude" json:"latitude" binding:"exists"`
	Address        string   	`form:"address" json:"address" binding:"exists"`
	PhoneNumber    string  		`form:"phone_number" json:"phone_number" binding:"exists"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Store{})
}