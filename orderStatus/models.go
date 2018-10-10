package orderStatus

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderStatus struct {
	gorm.Model						`json:"-"`
	StatusID 			uint		`form:"statusid" json:"status_id"`
	StatusChangedTime 	time.Time	`form:"statuschangedtime" json:"status_changed_time"`
	Description 		string 		`form:"description" json:"description"`
}



