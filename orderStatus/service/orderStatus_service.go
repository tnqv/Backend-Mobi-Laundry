package service

import (
	"d2d-backend/orderStatus"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type orderStatusService struct {
	orderStatusRepos	orderStatus.OrderStatusRepository
}

func NewOrderStatusService(orderStatusRepository orderStatus.OrderStatusRepository) orderStatus.OrderStatusService {
	return &orderStatusService{orderStatusRepository}
}


func (orderStatusService *orderStatusService) CreateNewOrderStatus(newOrderStatus *models.OrderStatus) (*models.OrderStatus, error) {
	_,err := orderStatusService.orderStatusRepos.Create(newOrderStatus)

	if err != nil {
		return nil,err
	}
	return newOrderStatus,nil
}

func (orderStatusService *orderStatusService) GetOrderStatus(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := orderStatusService.orderStatusRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (orderStatusService *orderStatusService) GetOrderStatusById(id int) (*models.OrderStatus, error) {
	orderStatusModel,err := orderStatusService.orderStatusRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return orderStatusModel,nil
}

func (orderStatusService *orderStatusService) UpdateOrderStatus(updateOrderStatus *models.OrderStatus) (*models.OrderStatus, error) {
	updateOrderStatus,err := orderStatusService.orderStatusRepos.Update(updateOrderStatus)
	if err != nil {
		return nil,err
	}
	return updateOrderStatus,nil
}

func (orderStatusService *orderStatusService) DeleteOrderStatus(id int) (bool, error) {
	isDeletedSuccess,err := orderStatusService.orderStatusRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess,err
	}
	return isDeletedSuccess,nil
}

