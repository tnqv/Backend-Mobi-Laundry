package models

import "github.com/jinzhu/gorm"

type UserShippingLocation struct {
	gorm.Model
	PhoneNumber			string		`form:"phone_number" json:"phone_number"`
	ShippingAddress		string 		`form:"shipping_address" json:"shipping_address" sql:"type:VARCHAR CHARACTER SET utf8 COLLATE utf8_general_ci"`
	Longitude			uint		`form:"longitude" json:"longitude"`
	Latitude			uint		`form:"latitude" json:"latitude"`
	ReceiverName		string		`form:"receiver_name" json:"receiver_name" sql:"type:VARCHAR CHARACTER SET utf8 COLLATE utf8_general_ci"`
	UserId				uint		`json:"user_id"`
}