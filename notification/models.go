package notification

import (
	"d2d-backend/accounts"

	"github.com/jinzhu/gorm"
	"d2d-backend/common"
)

type Notification struct {
	gorm.Model         						`json:"-"`
	NotificationTypeID 		uint   			`form:"notifucation_type_id" json:"notification_type_id" binding:"exists"`
	Read               		bool   			`form:"read" json:"read" binding:"exists"`
	Content            		string 			`form:"content" json:"content" binding:"exists"`
	//Customer
	UserID    				uint          	`form:"user_id" json:"user_id" binding:"exists"`
	UserModel 				accounts.User 	`json:"-"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Notification{})
}