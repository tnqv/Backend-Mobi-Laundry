package repository

import (
	"d2d-backend/common"
	"d2d-backend/user"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository() user.UserRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.User, error) {
	var userModel models.User
	err := r.Conn.Preload("Account").Preload("Role").Preload("Store").First(&userModel,id).Error
	if err != nil {
		return nil,err
	}
	return &userModel, nil
}

func (r *repo) FindUserByPhoneNumber(phoneNumber string)(*models.User,error){
	var userTemp models.User

	if err := r.Conn.Where("phone_number = ?",phoneNumber).First(&userTemp).Error; err != nil {
		return nil, err
	}

	return &userTemp,nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var userModels []models.User
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &userModels)
	return paginator,nil
}

func (r *repo) Create(userModel *models.User) (*models.User, error) {

	err := r.Conn.Create(userModel).Preload("Role").First(userModel).Error
	if err != nil {
		return nil,err
	}

	return userModel,nil
}

func (r *repo) Update(updateUser *models.User) (*models.User, error) {
	var tempUser models.User
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
	var tempUser models.User
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


func (r *repo) FindUserByAccountId(accountId uint)(*models.User,error){
	var tempUser models.User
	if err := r.Conn.Where("account_id = ?",accountId).
						Preload("Role").
						Preload("Store").
						Preload("ShippingLocations").
						First(&tempUser).Error; err != nil{
		return nil,err
	}
	return &tempUser,nil
}

func (r *repo) SaveNewUserLocation(location *models.UserShippingLocation) (*models.UserShippingLocation,error) {
	location.DeletedAt = nil
	err := r.Conn.Create(location).Error
	if err != nil {
		return nil,err
	}
	return location,nil
}

func (r *repo) UpdateUserLocation(location *models.UserShippingLocation) (*models.UserShippingLocation,error) {
	var locationTemp models.UserShippingLocation
	err := r.Conn.First(&locationTemp, location.ID).Error
	if err != nil{
		return nil, err
	}
	location.DeletedAt = nil
	err = r.Conn.Model(&location).Update(&location).First(&locationTemp).Error
	if err != nil {
		return nil, err
	}
	return &locationTemp,nil
}