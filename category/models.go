package category

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/service"
	"d2d-backend/common"
)

type Category struct {
	gorm.Model				`json:"-"`
	Name 		string		`form:"name" json:"name" binding:"exists"`
	Description string 		`form:"description" json:"description" binding:"exists"`
	Services 	[]service.Service `json:"services"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Category{})
}