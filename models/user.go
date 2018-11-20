package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name 				string		`form:"name" json:"name" binding:"exists"`
	PhoneNumber			string		`gorm:"not null;unique" form:"phone_number" json:"phone_number" binding:"exists"`
	ShippingAddress		string 		`form:"shipping_address" json:"shipping_address"`
	Longitude			uint		`form:"longitude" json:"longitude"`
	Latitude			uint		`form:"latitude" json:"latitude"`
	RoleId				uint		`form:"role_id" json:"role_id"`
	Role				Role		`gorm:"auto_preload" json:"role" gorm:"save_associations:false"`
	AccountId			uint 		`gorm:"not null;unique" form:"account_id" json:"-"`
	Account				Account		`json:"account" gorm:"save_associations:false"`
	StoreId				uint		`form:"storeid" json:"store_id"`
	Store 				Store		`json:"store,omitempty" gorm:"save_associations:false"`
	Address 			string		`form:"address" json:"address"`
	IdentifyNumber		string		`form:"identify_number" json:"identify_number"`
	Capacity			uint		`form:"capacity" json:"capacity"`
	Imageurl			string		`form:"imageurl" json:"imageurl"`
}
