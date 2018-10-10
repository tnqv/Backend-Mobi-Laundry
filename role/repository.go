package role

import "github.com/biezhi/gorm-paginator/pagination"

type RoleRepository interface{
	Find(id int) (*Role, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(role *Role) (*Role,error)
	Update(role *Role) (*Role, error)
	Delete(id int) (bool,error)
}