package repository

import (
	"d2d-backend/common"
	"d2d-backend/role"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlRoleRepository() role.RoleRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.Role, error) {
	var roleModel models.Role
	err := r.Conn.First(&roleModel, id).Error
	if err != nil {
		return nil, err
	}
	return &roleModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var roles []*models.Role
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &roles)

	return paginator,nil
}

func (r *repo) Create(role *models.Role) (*models.Role, error) {
	role.DeletedAt = nil
	err := r.Conn.Create(role).Error
	if err != nil {
		return nil,err
	}
	return role,nil
}

func (r *repo) Update(updateRole *models.Role) (*models.Role, error) {
	var tempRole models.Role
	err := r.Conn.First(&tempRole,updateRole.ID).Error
	if err != nil{
		return nil, err
	}
	updateRole.DeletedAt = nil
	err = r.Conn.Model(&updateRole).Update(&updateRole).Error
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempRole models.Role
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

