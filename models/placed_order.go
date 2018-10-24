package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PlacedOrder struct {
	gorm.Model
	StoreID 				uint					`form:"store_id" json:"store_id"`
	StoreModel 				Store					`json:"-"`
	TimePlaced 				time.Time				`form:"time_placed" json:"time_placed"`
	Detail 					string					`form:"detail" json:"detail"`
	//OrderStatusId 		uint					`form:"order_status_id" json:"order_status_id"`
	OrderStatusId			uint					`json:"current_status_id"`
	//OrderStatus				OrderStatus				`json:"order_status"`
	OrderStatuses 			[]OrderStatus			`form:"-" json:"order_status_list"`
	ServiceOrders 			[]ServiceOrder			`form:"-" json:"order_service_list"`
	//Customer
	UserID 					uint					`form:"user_id" json:"user_id"`
	UserModel 				User					`json:"-"`
	//OrderInformation
	ReceiverName			string 					`form:"receiver_name" json:"receiver_name"`
	ReceiverPhone			string					`form:"receiver_phone" json:"receiver_phone"`
	Note 					string					`form:"note" json:"note"`
	Capacity 				float32					`form:"capacity" json:"capacity"`
	EstimatedCapacity 		float32					`form:"estimated_capacity" json:"estimated_capacity"`
	EstimatedDeliveryTime	time.Time				`form:"-" json:"estimated_delivery_time"`
	ActualDeliveryTime			time.Time				`form:"-" json:"actual_delivery_time"`
	DeliveryAddress 		string					`form:"delivery_address" json:"delivery_address"`
	DeliveryLatitude 		float32					`form:"delivery_latitude" json:"delivery_latitude"`
	DeliveryLongitude 		float32					`form:"delivery_longitude" json:"delivery_longitude"`
	ServiceTotalPrice 		float32					`form:"total" json:"total"`
	Priority 				int						`form:"priority" json:"priority"`
	OrderCode 				string					`form:"order_code" json:"order_code"`
	//Review
	ReviewID 				uint					`form:"" json:"review_id"`
	ReviewModel 			Review					`json:"-"`
}
