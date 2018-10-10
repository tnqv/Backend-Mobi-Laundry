package repository

import (
	"d2d-backend/common"
	"d2d-backend/serviceOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlServiceOrderRepository() serviceOrder.ServiceOrderRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*serviceOrder.ServiceOrder, error) {
	var serviceOrder serviceOrder.ServiceOrder
	err := r.Conn.First(&serviceOrder, id).Error
	if err != nil {
		return nil, err
	}
	return &serviceOrder, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var serviceOrders []*serviceOrder.ServiceOrder
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &serviceOrders)
	return paginator, nil
}

func (r *repo) Create(serviceOrder *serviceOrder.ServiceOrder) (*serviceOrder.ServiceOrder, error) {
	err := r.Conn.Create(serviceOrder).Error
	if err != nil {
		return nil, err
	}
	return serviceOrder, nil
}

func (r *repo) Update(updateServiceOrder *serviceOrder.ServiceOrder) (*serviceOrder.ServiceOrder, error) {
	var tempServiceOrder serviceOrder.ServiceOrder
	err := r.Conn.First(&tempServiceOrder,updateServiceOrder.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateServiceOrder).Error
	if err != nil {
		return nil, err
	}
	return updateServiceOrder, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempServiceOrder serviceOrder.ServiceOrder
	err := r.Conn.First(&tempServiceOrder, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempServiceOrder).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

