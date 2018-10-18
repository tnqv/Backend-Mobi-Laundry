package category

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type CategoryRepository interface{
	Find(id int) (*models.Category, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *models.Category) (*models.Category,error)
	Update(category *models.Category) (*models.Category, error)
	Delete(id int) (bool,error)
}
