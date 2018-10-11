package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model						`json:"-"`
	Name 				string		`form:"name" json:"name" binding:"exists"`
	PhoneNumber			string		`form:"phone_number" json:"phone_number" binding:"exists"`
	ShippingAddress		string 		`form:"shipping_address" json:"shipping_address"`
	Longitude			uint		`form:"longitude" json:"longitude"`
	Latitude			uint		`form:"latitude" json:"latitude"`
	RoleId				uint		`form:"role_id" json:"role_id"`
	AccountId			uint 		`form:"account_id" json:"account_id"`
	StoreId				uint		`form:"storeid" json:"store_id"`
	Address 			string		`form:"address" json:"address"`
	IdentifyNummber		uint		`form:"identify_number" json:"identify_nummber"`
	Capacity			uint		`form:"capacity" json:"capacity"`
	Imageurl			string		`form:"imageurl" json:"imageurl"`
}
