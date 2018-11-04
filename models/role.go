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
//2 : delivery
//3 : store
//4 : dieu phoi/admin