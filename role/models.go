package role

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
)

type Role struct {
	gorm.Model				`json:"-"`
	Name		string		`form:"name" json:"name" binding:"exists"`
	Description string		`form:"description" json:"description" binding:"exists"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Role{})
}