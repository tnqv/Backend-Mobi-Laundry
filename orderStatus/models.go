package orderStatus

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderStatus struct {
	gorm.Model						`json:"-"`
	StatusID 			uint		`form:"status_id" json:"status_id"`
	UserID				uint		`form:"user_id" json:"user_id"`
	UserModel			uint		`json:"-"`
	StatusChangedTime 	time.Time	`form:"status_changed_time" json:"status_changed_time"`
	Description 		string 		`form:"description" json:"description"`
}



