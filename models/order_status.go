package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


//Status ID
//--1
// 1 : Order just Created
// 2 : Accepted orders
//--2
// 3 : Delivery verified orders
// 4 : Pick up order to warehouse
//--3
// 5 : Store xx received order
// 6 : Order is in process
// 7 : Order is finished laundry
//--4
// 8 : Order is delivering

//--5
// 9 : Order completed

type OrderStatus struct {
	gorm.Model						`json:"-"`
	StatusID 			uint		`form:"status_id" json:"status_id"`
	//Who change status
	UserId				uint		`form:"user_id" json:"user_id"`
	User				User		`json:"-"`
	StatusChangedTime 	time.Time	`form:"status_changed_time" json:"status_changed_time"`
	Description 		string 		`form:"description" json:"description"`
	PlacedOrderID		uint 		`form:"place_order_id" json:"-"`
	PlacedOrder 		PlacedOrder `json:"-"`
}
