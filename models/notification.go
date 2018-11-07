package models

import (
	"d2d-backend/accounts"

	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	NotificationTypeID 		uint   			`form:"notifucation_type_id" json:"notification_type_id" binding:"exists"`
	Read               		bool   			`form:"read" json:"read" binding:"exists"`
	Content            		string 			`form:"content" json:"content" binding:"exists"`
	//Customer
	UserID    				uint          	`form:"user_id" json:"user_id" binding:"exists"`
	User 				accounts.User 	`json:"-" gorm:"save_associations:false"`
}
