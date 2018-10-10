package role

import "github.com/biezhi/gorm-paginator/pagination"

type RoleService interface {
	GetRole(name string) (*Role, error)
	CreateNewRole(newRole *Role)(*Role, error)
	GetRoles(limit int, page int)(*pagination.Paginator, error)
	GetRoleById(id int)(*Role, error)
	UpdateRole(updateRole *Role)(*Role, error)
	DeleteRole(id int)(bool, error)
}