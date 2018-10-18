package role

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type RoleService interface {
	GetRole(name string) (*models.Role, error)
	CreateNewRole(newRole *models.Role)(*models.Role, error)
	GetRoles(limit int, page int)(*pagination.Paginator, error)
	GetRoleById(id int)(*models.Role, error)
	UpdateRole(updateRole *models.Role)(*models.Role, error)
	DeleteRole(id int)(bool, error)
}