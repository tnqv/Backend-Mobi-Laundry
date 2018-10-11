package service

import (
	"d2d-backend/placedOrder"
	"github.com/biezhi/gorm-paginator/pagination"
)

type placedOrderService struct {
	placedOrderRepos placedOrder.PlacedOrderRepository
}

func NewPlacedOrderService(placedOrderRepository placedOrder.PlacedOrderRepository) placedOrder.PlacedOrderService{
	return &placedOrderService{placedOrderRepository}
}

func (placedOrderService *placedOrderService) CreateNewPlacedOrder(newPlacedOrder *placedOrder.PlacedOrder) (*placedOrder.PlacedOrder, error) {
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

func (placedOrderService *placedOrderService) GetPlacedOrderById(id int) (*placedOrder.PlacedOrder, error) {
	role, err := placedOrderService.placedOrderRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (placedOrderService *placedOrderService) UpdatePlacedOrder(updatePlacedOrder *placedOrder.PlacedOrder) (*placedOrder.PlacedOrder, error) {
	updateRole, err := placedOrderService.placedOrderRepos.Update(updatePlacedOrder)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (placedOrderService *placedOrderService) DeletePlacedOrder(id int) (bool, error) {
	bool, err := placedOrderService.placedOrderRepos.Delete(id)
	if err != nil {
		return bool, err
	}
	return bool, nil
}

func (placedOrderService *placedOrderService) GetListOrdersByUserId(limit int, page int, id int) (*pagination.Paginator, error) {
	paginate,err := placedOrderService.placedOrderRepos.FindByUserId(limit, page, id)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}



