package service

import (
	"d2d-backend/role"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type roleService struct {
	roleRepos role.RoleRepository
}

func NewRoleService(roleRepository role.RoleRepository) role.RoleService {
	return &roleService{roleRepository}
}

func (roleService *roleService) GetRole(name string) (*models.Role, error) {
	panic("implement me")
}

func (roleService *roleService) CreateNewRole(newRole *models.Role) (*models.Role, error) {
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

func (roleService *roleService) GetRoleById(id int) (*models.Role, error) {
	roleModel, err := roleService.roleRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return roleModel, nil
}

func (roleService *roleService) UpdateRole(updateRole *models.Role) (*models.Role, error) {
	updateRole, err := roleService.roleRepos.Update(updateRole)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (roleService *roleService) DeleteRole(id int) (bool, error) {
	isDeletedSuccess, err := roleService.roleRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess, err
	}
	return isDeletedSuccess, nil
}

