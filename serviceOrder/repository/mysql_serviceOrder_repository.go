package repository

import (
	"d2d-backend/common"
	"d2d-backend/serviceOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlServiceOrderRepository() serviceOrder.ServiceOrderRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.ServiceOrder, error) {
	var serviceOrderModel models.ServiceOrder
	err := r.Conn.First(&serviceOrderModel, id).Error
	if err != nil {
		return nil, err
	}
	return &serviceOrderModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var serviceOrders []*models.ServiceOrder
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &serviceOrders)
	return paginator, nil
}

func (r *repo) CreateServiceOrders(serviceorders []*models.ServiceOrder)([]*models.ServiceOrder,error){
	tx := r.Conn.Begin()
	for i:= 0; i < len(serviceorders); i++ {
		if err := tx.Create(&serviceorders[i]).Error; err != nil{
					tx.Rollback()
					return nil,err
		}

	}
	//for k := range serviceorders{
	//	if err := tx.Create(&k).Error; err != nil{
	//		tx.Rollback()
	//		return nil,err
	//	}
	//
	//}

	err := tx.Commit().Error
	if err != nil{
		return nil,err
	}

	return serviceorders,nil
}
func (r *repo) Create(serviceOrder *models.ServiceOrder) (*models.ServiceOrder, error) {
	err := r.Conn.Create(serviceOrder).Error
	if err != nil {
		return nil, err
	}
	return serviceOrder, nil
}

func (r *repo) Update(updateServiceOrder *models.ServiceOrder) (*models.ServiceOrder, error) {
	var tempServiceOrder models.ServiceOrder
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
	var tempServiceOrder models.ServiceOrder
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

