package repository

import (
	"d2d-backend/common"
	"d2d-backend/placedOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlPlacedOrderRepository() placedOrder.PlacedOrderRepository{
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.PlacedOrder, error) {
	var placedOrderModel models.PlacedOrder
	//err := r.Conn.Preload("OrderStatuses").
	//			  Preload("ServiceOrders").
	//			  Preload("ServiceOrders.Service").
	//	          Preload("Store").First(&placedOrderModel, id).Error
	//err := r.Conn.Preload("Store").Preload("User").Preload("User.Role").First(&placedOrderModel, id).Error
	err := r.Conn.Preload("User").
				  Preload("Store").
				  Preload("User.Account").
		          Preload("Delivery").
				  Preload("Delivery.Account").
		     	  First(&placedOrderModel, id).Error
	if err != nil {
		return nil, err
	}
	return &placedOrderModel, nil
}

func (r *repo) FindByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	var placedOrders []*models.PlacedOrder
	db := r.Conn
	db = db.Where("user_id = ?", id)
	paginator := pagination.Pagging(&pagination.Param{
		DB: db.Preload("ServiceOrders").Preload("ServiceOrders.Service").Preload("OrderStatuses", func(db *gorm.DB) *gorm.DB{
				return db.Order("status_id DESC")
		}),
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &placedOrders)
	return paginator,nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var placedOrders []*models.PlacedOrder
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &placedOrders)

	return paginator,nil
}

func (r *repo) FindPlacedOrderByOrderCode(orderCode string)(*models.PlacedOrder,error){
	var placeOrder models.PlacedOrder

	err := r.Conn.Where("order_code = ?",orderCode).
	Preload("OrderStatuses",func(db *gorm.DB) *gorm.DB{
		return db.Order("status_id DESC")
	}).
	Preload("ServiceOrders").
	Preload("ServiceOrders.Service").
	Preload("User").
	Preload("Store").
	First(&placeOrder).Error

	if err != nil {
		return nil,err
	}

	return &placeOrder,nil

}


func (r *repo) Create(placedOrder *models.PlacedOrder) (*models.PlacedOrder, error) {
	err := r.Conn.
		Create(placedOrder).Preload("User").
		Preload("Store").
		Preload("User.Account").
		Preload("Delivery").
		Preload("Delivery.Account").First(placedOrder).Error
	if err != nil {
		return nil,err
	}
	return placedOrder,nil
}

func (r *repo) Update(updatePlacedOrder *models.PlacedOrder) (*models.PlacedOrder, error) {
	//var tempPlacedOrder models.Role
	//err := r.Conn.Preload("OrderStatuses").First(&tempPlacedOrder,updatePlacedOrder.ID).Error
	//if err != nil{
	//	return nil, err
	//}

	err := r.Conn.Save(updatePlacedOrder).Error
	if err != nil {
		return nil, err
	}
	if updatePlacedOrder.DeliveryID != 0 {
		r.Conn.Preload("Delivery").
			Preload("Delivery.Account").First(updatePlacedOrder)
	}
	if updatePlacedOrder.StoreID != 0 {
		r.Conn.Preload("Store").First(updatePlacedOrder)
	}
	return updatePlacedOrder, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempPlacedOrder models.PlacedOrder
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


func (r *repo) UpdateOrderStatusId(placedOrder *models.PlacedOrder) (*models.PlacedOrder, error){

	//r.Conn.Model(&placedOrder).Select("order_status_id").Updates(map[string]interface{}{"name": "hello",})
	return nil,nil
}
