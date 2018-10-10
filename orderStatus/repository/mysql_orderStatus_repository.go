package repository

import (
	"d2d-backend/common"
	"d2d-backend/orderStatus"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlOrderStatusRepository() orderStatus.OrderStatusRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*orderStatus.OrderStatus, error) {
	var orderStatus orderStatus.OrderStatus
	err := r.Conn.First(&orderStatus,id).Error
	if err != nil {
		return nil,err
	}
	return &orderStatus, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var orderStatus []orderStatus.OrderStatus
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &orderStatus)
	return paginator,nil
}

func (r *repo) Create(orderStatus *orderStatus.OrderStatus) (*orderStatus.OrderStatus, error) {
	err := r.Conn.Create(orderStatus).Error
	if err != nil {
		return nil,err
	}
	return orderStatus,nil
}

func (r *repo) Update(updateOrderStatus *orderStatus.OrderStatus) (*orderStatus.OrderStatus, error) {
	var tempOrderStatus orderStatus.OrderStatus
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
	var tempOrderStatus orderStatus.OrderStatus
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
