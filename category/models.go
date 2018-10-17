package category

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/service"
)

type Category struct {
	gorm.Model				`json:"-"`
	Name 		string		`form:"name" json:"name" binding:"exists"`
	Description string 		`form:"description" json:"description" binding:"exists"`
	Services 	[]service.Service `json:"services"`
}
