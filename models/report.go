package models

import "github.com/jinzhu/gorm"

type Report struct{
	gorm.Model
	Content 		string 	`form:"content" json:"content"`
	PlacedOrderId 	uint	`json:"placed_order_id"`
	IsResolved		bool	`form:"is_resolved" json:"is_resolved"`
}