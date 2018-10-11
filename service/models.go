package service

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model					`json:"-"`
	Name 			string		`form:"name" json:"name"`
	Price 			int64		`form:"price" json:"price"`
	Description 	string		`form:"description" json:"description"`
	ImageUrl		string		`json:"-"`
	CategoryId 		uint		`json:"-"`
}