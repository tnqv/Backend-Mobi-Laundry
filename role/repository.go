package role

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type RoleRepository interface{
	Find(id int) (*models.Role, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(role *models.Role) (*models.Role,error)
	Update(role *models.Role) (*models.Role, error)
	Delete(id int) (bool,error)
}