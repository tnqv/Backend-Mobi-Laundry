package repository

import (
	"d2d-backend/common"
	"d2d-backend/role"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlRoleRepository() role.RoleRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*role.Role, error) {
	var role role.Role
	err := r.Conn.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var roles []*role.Role
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &roles)

	return paginator,nil
}

func (r *repo) Create(role *role.Role) (*role.Role, error) {
	err := r.Conn.Create(role).Error
	if err != nil {
		return nil,err
	}
	return role,nil
}

func (r *repo) Update(updateRole *role.Role) (*role.Role, error) {
	var tempRole role.Role
	err := r.Conn.First(&tempRole,updateRole.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateRole).Error
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempRole role.Role
	err := r.Conn.First(&tempRole, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempRole).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

