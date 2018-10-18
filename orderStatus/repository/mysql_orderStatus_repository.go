package repository

import (
	"d2d-backend/common"
	"d2d-backend/orderStatus"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlOrderStatusRepository() orderStatus.OrderStatusRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.OrderStatus, error) {
	var orderStatusModel models.OrderStatus
	err := r.Conn.First(&orderStatusModel,id).Error
	if err != nil {
		return nil,err
	}
	return &orderStatusModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var orderStatuses []models.OrderStatus
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &orderStatuses)
	return paginator,nil
}

func (r *repo) Create(orderStatus *models.OrderStatus) (*models.OrderStatus, error) {
	err := r.Conn.Create(orderStatus).Error
	if err != nil {
		return nil,err
	}
	return orderStatus,nil
}

func (r *repo) Update(updateOrderStatus *models.OrderStatus) (*models.OrderStatus, error) {
	var tempOrderStatus models.OrderStatus
	err := r.Conn.First(&tempOrderStatus,updateOrderStatus.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateOrderStatus).Error
	if err != nil {
		return nil, err
	}
	return updateOrderStatus,nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempOrderStatus models.OrderStatus
	err := r.Conn.First(&tempOrderStatus,id).Error
	if err != nil {
		return false,err
	}
	err = r.Conn.Delete(&tempOrderStatus).Error
	if err != nil {
		return false,err
	}
	return true,nil
}
