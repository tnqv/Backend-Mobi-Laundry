package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name 				string					`form:"name" json:"name" binding:"exists"`
	ShippingLocations	[]UserShippingLocation	`json:"shipping_locations"`
	PhoneNumber			string					`json:"phone_number"`
	RoleId				uint					`form:"role_id" json:"role_id"`
	Role				Role					`gorm:"auto_preload" json:"role" gorm:"save_associations:false"`
	AccountId			uint 					`gorm:"not null;unique" form:"account_id" json:"account_id"`
	Account				Account					`json:"account" gorm:"save_associations:false"`
	StoreId				uint					`form:"store_id" json:"store_id"`
	Store 				Store					`json:"store,omitempty" gorm:"save__associations:false"`
	Address 			string					`form:"address" json:"address"`
	IdentifyNumber		string					`form:"identify_number" json:"identify_number"`
	Capacity			uint					`form:"capacity" json:"capacity"`
	Imageurl			string					`form:"imageurl" json:"imageurl"`
}
