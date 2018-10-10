package service

import (
	"d2d-backend/role"
	"github.com/biezhi/gorm-paginator/pagination"
)

type roleService struct {
	roleRepos role.RoleRepository
}

func NewRoleService(roleRepository role.RoleRepository) role.RoleService {
	return &roleService{roleRepository}
}

func (roleService *roleService) GetRole(name string) (*role.Role, error) {
	panic("implement me")
}

func (roleService *roleService) CreateNewRole(newRole *role.Role) (*role.Role, error) {
	_, err := roleService.roleRepos.Create(newRole)
	if err != nil {
		return nil, err
	}
	return newRole, nil
}

func (roleService *roleService) GetRoles(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := roleService.roleRepos.FindAll(limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (roleService *roleService) GetRoleById(id int) (*role.Role, error) {
	role, err := roleService.roleRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (roleService *roleService) UpdateRole(updateRole *role.Role) (*role.Role, error) {
	updateRole, err := roleService.roleRepos.Update(updateRole)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (roleService *roleService) DeleteRole(id int) (bool, error) {
	bool, err := roleService.roleRepos.Delete(id)
	if err != nil {
		return bool, err
	}
	return bool, nil
}

