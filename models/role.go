package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model				`json:"-"`
	Name		string		`form:"name" json:"name" binding:"exists"`
	Description string		`form:"description" json:"description" binding:"exists"`
}
//1 : customer
//2 : delivery giao
//3 : delievery nhận
//4 : store
//5 : dieu phoi