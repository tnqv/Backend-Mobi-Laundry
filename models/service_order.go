package models

import (
	"github.com/jinzhu/gorm"
)

type ServiceOrder struct {
	gorm.Model							`json:"-"`
	PlacedOrderID	uint				`form:"placed_order_id" json:"placed_order_id"`
	PlacedOrder 	PlacedOrder			`json:"-"  gorm:"save_associations:false"`
	ServiceID		uint				`form:"service_id" json:"service_id"`
	Service 		Service				`json:"service"  gorm:"save_associations:false"`
	Description 	string				`form:"description" json:"description"`
	Quantity 		uint				`form:"quantity" json:"quantity"`
	Price 			uint				`form:"price" json:"price"`
}
