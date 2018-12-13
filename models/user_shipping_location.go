package models

import "github.com/jinzhu/gorm"

type UserShippingLocation struct {
	gorm.Model
	PhoneNumber			string		`form:"phone_number" json:"phone_number"`
	ShippingAddress		string 		`form:"shipping_address" json:"shipping_address"`
	Longitude			float32		`form:"longitude" json:"longitude"`
	Latitude			float32		`form:"latitude" json:"latitude"`
	ReceiverName		string		`form:"receiver_name" json:"receiver_name"`
	UserId				uint		`json:"user_id"`
}