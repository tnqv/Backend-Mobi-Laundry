package service

import (
	"d2d-backend/placedOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type placedOrderService struct {
	placedOrderRepos placedOrder.PlacedOrderRepository
}

func NewPlacedOrderService(placedOrderRepository placedOrder.PlacedOrderRepository) placedOrder.PlacedOrderService{
	return &placedOrderService{placedOrderRepository}
}

func (placedOrderService *placedOrderService) CreateNewPlacedOrder(newPlacedOrder *models.PlacedOrder) (*models.PlacedOrder, error) {
	_, err := placedOrderService.placedOrderRepos.Create(newPlacedOrder)
	if err != nil {
		return nil, err
	}
	return newPlacedOrder, nil
}

func (placedOrderService *placedOrderService) GetPlacedOrders(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := placedOrderService.placedOrderRepos.FindAll(limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (placedOrderService *placedOrderService) GetPlacedOrderById(id int) (*models.PlacedOrder, error) {
	role, err := placedOrderService.placedOrderRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (placedOrderService *placedOrderService) UpdatePlacedOrder(updatePlacedOrder *models.PlacedOrder) (*models.PlacedOrder, error) {
	updateRole, err := placedOrderService.placedOrderRepos.Update(updatePlacedOrder)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (placedOrderService *placedOrderService) DeletePlacedOrder(id int) (bool, error) {
	isDeleted, err := placedOrderService.placedOrderRepos.Delete(id)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}

func (placedOrderService *placedOrderService) GetListOrdersByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	paginate,err := placedOrderService.placedOrderRepos.FindByUserId(limit, page, id)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}



