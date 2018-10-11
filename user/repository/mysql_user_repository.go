package repository

import (
	"d2d-backend/common"
	"d2d-backend/user"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository() user.UserRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*user.User, error) {
	var user user.User
	err := r.Conn.First(&user,id).Error
	if err != nil {
		return nil,err
	}
	return &user, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var user []user.User
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &user)
	return paginator,nil
}

func (r *repo) Create(user *user.User) (*user.User, error) {
	err := r.Conn.Create(user).Error
	if err != nil {
		return nil,err
	}
	return user,nil
}

func (r *repo) Update(updateUser *user.User) (*user.User, error) {
	var tempUser user.User
	err := r.Conn.First(&tempUser,updateUser.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateUser).Error
	if err != nil {
		return nil, err
	}
	return updateUser,nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempUser user.User
	err := r.Conn.First(&tempUser,id).Error
	if err != nil {
		return false,err
	}
	err = r.Conn.Delete(&tempUser).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

