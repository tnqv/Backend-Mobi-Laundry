package review

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/accounts"
	"d2d-backend/common"
)

type Review struct {
	gorm.Model						`json:"-"`
	Content 	   string			`json:"content"`
	Rate           int		   		`json:"rate"`
	UserRate       accounts.User  	`json:"user"`
	UserID		   int				`json:"-"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Review{})
}