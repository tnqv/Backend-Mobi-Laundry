package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model						`json:"-"`
	Name 				string		`form:"name" json:"name" binding:"exists"`
	PhoneNumber			string		`form:"phone_number" json:"phone_number" binding:"exists"`
	//ShippingAddress		string 		`form:"shippingaddress" json:"shipping_address"`
	//Longitude			uint		`form:"longitude" json:"longitude"`
	//Latitude			uint		`form:"latitude" json:"latitude"`
	//RoleId				uint		`form:"roleid" json:"roleid"`
	//AccountId			uint 		`form:"accountid" json:"account_id"`
	//StoreId				uint		`form:"storeid" json:"store_id"`
	//Address 			string		`form:"address" json:"address"`
	//Identifynummber		uint		`form:"identifynumber" json:"identifynummber"`
	//Capacity			uint		`form:"capacity" json:"capacity"`
	//Imageurl			string		`form:"imageurl" json:"imageurl"`
	
}
