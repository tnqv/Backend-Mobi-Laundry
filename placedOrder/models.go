package placedOrder

import (
	"d2d-backend/orderStatus"
	"d2d-backend/review"
	"d2d-backend/store"
	"d2d-backend/user"
	"github.com/jinzhu/gorm"
	"time"
)

type PlacedOrder struct {
	gorm.Model									`json:"-"`
	StoreID 			uint					`json:"store_id"`
	StoreModel 			store.Store				`json:"-"`
	TimePlaced 			time.Time				`json:"time_placed"`
	Detail 				string					`json:"detail"`
	OrderStatusID 		uint					`json:"order_status_id"`
	OrderStatusModel 	orderStatus.OrderStatus	`json:"-"`
	//Customer
	UserID 				uint					`json:"user_id"`
	UserModel 			user.User				`json:"-"`
	//OrderInformation
	Capacity 			float32					`json:"capacity"`
	EstimatedCapacity 	float32					`json:"estimated_capacity"`
	DeliveryAddress 	string					`json:"delivery_address"`
	DeliveryLatitude 	float32					`json:"delivery_latitude"`
	DeliveryLongitude 	float32					`json:"delivery_longitude"`
	ServiceTotalPrice 	float32					`json:"total"`
	Priority 			int						`json:"priority"`
	OrderCode 			string					`json:"order_code"`
	//Review
	ReviewID 			uint					`json:"review_id""`
	ReviewModel 		review.Review			`json:"-"`
}
