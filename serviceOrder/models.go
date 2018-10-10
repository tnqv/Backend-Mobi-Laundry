package serviceOrder

import (
	"d2d-backend/orders"
	"d2d-backend/service"
	"github.com/jinzhu/gorm"
)

type ServiceOrder struct {
	gorm.Model							`json:"-"`
	PlacedOrderID	uint				`form:"placed_order_id" json:"placed_order_id"`
	PlacedOrder 	orders.PlacedOrder	`json:"-"`
	ServiceId		uint				`form:"service_id" json:"service_id"`
	Service 		service.Service		`json:"-"`
	Description 	string				`form:"description" json:"description"`
	Quantity 		uint				`form:"quantity" json:"quantity"`
	Price 			uint				`form:"price" json:"price"`
}
