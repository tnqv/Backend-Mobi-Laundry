package models

import (

	"github.com/jinzhu/gorm"
	"time"
)

type PlacedOrder struct {
	gorm.Model									`json:"-"`
	StoreID 			uint					`form:"store_id" json:"store_id"`
	StoreModel 			Store				`json:"-"`
	TimePlaced 			time.Time				`form:"time_placed" json:"time_placed"`
	Detail 				string					`form:"detail" json:"detail"`
	//OrderStatusID 		uint					`form:"order_status_id" json:"order_status_id"`
	ListOrderStatusModel 	[]OrderStatus	`json:"order_status"`
	//Customer
	UserID 				uint					`form:"user_id" json:"user_id"`
	UserModel 			User				`json:"-"`
	//OrderInformation
	Capacity 			float32					`form:"capacity" json:"capacity"`
	EstimatedCapacity 	float32					`form:"estimated_capacity" json:"estimated_capacity"`
	DeliveryAddress 	string					`form:"delivery_address" json:"delivery_address"`
	DeliveryLatitude 	float32					`form:"delivery_latitude" json:"delivery_latitude"`
	DeliveryLongitude 	float32					`form:"delivery_longitude" json:"delivery_longitude"`
	ServiceTotalPrice 	float32					`form:"total" json:"total"`
	Priority 			int						`form:"priority" json:"priority"`
	OrderCode 			string					`form:"order_code" json:"order_code"`
	//Review
	ReviewID 			uint					`form:"" json:"review_id"`
	ReviewModel 		Review			`json:"-"`
}