package category

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model				`json:"-"`
	Name 		string		`form:"name" json:"name" binding:"exists"`
	Description string 		`form:"description" json:"description" binding:"exists"`
}
