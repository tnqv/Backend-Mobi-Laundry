package models

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	Content 	   string			`json:"content"`
	Rate           int		   		`json:"rate"`
	UserRate       User  			`json:"user" gorm:"save_associations:false"`
	UserID		   int				`json:"-"`
	User		   User 			`json:"user" gorm:"save_associations:false"`
}
