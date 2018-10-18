package category

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type CategoryService interface {
	GetCategoryDetailByName(name string) (*models.Category,error)
	CreateNewCategory(newCategory *models.Category)(*models.Category,error)
	GetCategory(limit int, page int)(*pagination.Paginator,error)
	GetCategoryById(id int)(*models.Category,error)
	UpdateCategory(updateCategory *models.Category)(*models.Category,error)
	DeleteCategory(id int)(bool,error)
}


