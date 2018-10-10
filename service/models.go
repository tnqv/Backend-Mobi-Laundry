package service

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model					`json:"-"`
	Name 			string		`json:"name"`
	Price 			int64		`json:"price"`
	Description 	string		`json:"description"`
	ImageUrl		string		`json:"-"`
	CategoryId 		uint		`json:"-"`
}