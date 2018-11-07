package models

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/accounts"
)

type Review struct {
	gorm.Model						`json:"-"`
	Content 	   string			`json:"content"`
	Rate           int		   		`json:"rate"`
	UserRate       accounts.User  	`json:"user"`
	UserID		   int				`json:"-"`
	User		   User 			`json:"user" gorm:"save_associations:false"`
}
