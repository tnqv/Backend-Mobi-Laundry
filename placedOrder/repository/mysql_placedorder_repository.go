package repository

import (
	"d2d-backend/common"
	"d2d-backend/placedOrder"
	"d2d-backend/role"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlPlacedOrderRepository() placedOrder.PlacedOrderRepository{
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*placedOrder.PlacedOrder, error) {
	var placedOrder placedOrder.PlacedOrder
	err := r.Conn.First(&placedOrder, id).Error
	if err != nil {
		return nil, err
	}
	return &placedOrder, nil
}

func (r *repo) FindByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	var placedOrders []*placedOrder.PlacedOrder
	db := r.Conn
	db = db.Where("user_id = ?", id)
	paginator := pagination.Pagging(&pagination.Param{
		DB: db,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &placedOrders)
	return paginator,nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var placedOrders []*placedOrder.PlacedOrder
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &placedOrders)

	return paginator,nil
}

func (r *repo) Create(placedOrder *placedOrder.PlacedOrder) (*placedOrder.PlacedOrder, error) {
	err := r.Conn.Create(placedOrder).Error
	if err != nil {
		return nil,err
	}
	return placedOrder,nil
}

func (r *repo) Update(updatePlacedOrder *placedOrder.PlacedOrder) (*placedOrder.PlacedOrder, error) {
	var tempPlacedOrder role.Role
	err := r.Conn.First(&tempPlacedOrder,updatePlacedOrder.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updatePlacedOrder).Error
	if err != nil {
		return nil, err
	}
	return updatePlacedOrder, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempPlacedOrder placedOrder.PlacedOrder
	err := r.Conn.First(&tempPlacedOrder, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&tempPlacedOrder).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

