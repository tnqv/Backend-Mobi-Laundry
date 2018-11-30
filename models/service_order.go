package models

import (
	"github.com/jinzhu/gorm"
)

type ServiceOrder struct {
	gorm.Model
	PlacedOrderID	uint				`form:"placed_order_id" json:"placed_order_id"`
	PlacedOrder 	PlacedOrder			`json:"-"  gorm:"save_associations:false"`
	ServiceID		uint				`form:"service_id" json:"service_id"`
	Service 		Service				`json:"service"  gorm:"save_associations:false"`
	Description 	string				`form:"description" json:"description"`
	Quantity 		uint				`form:"quantity" json:"quantity"`
	Price 			float32				`form:"price" json:"price"`
}
