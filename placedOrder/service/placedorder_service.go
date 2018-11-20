package service

import (
	"d2d-backend/placedOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
	"time"
	"d2d-backend/orderStatus"
)

type placedOrderService struct {
	placedOrderRepos placedOrder.PlacedOrderRepository
	orderStatusRepos orderStatus.OrderStatusRepository
}

func NewPlacedOrderService(placedOrderRepository placedOrder.PlacedOrderRepository,orderStatusRepository orderStatus.OrderStatusRepository) placedOrder.PlacedOrderService{
	return &placedOrderService{placedOrderRepository,orderStatusRepository}
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

func (placedOrderService *placedOrderService) GetPlacedOrderByOrderCode(orderCode string)(*models.PlacedOrder,error){
	placedOrderModel,err := placedOrderService.placedOrderRepos.FindPlacedOrderByOrderCode(orderCode)
	if err != nil {
		return  nil,err
	}
	return placedOrderModel,nil
}

func (placedOrderService *placedOrderService) UpdatePlacedOrderAndCreateNewOrderStatus(statusId uint,userId uint,description string,order *models.PlacedOrder)(*models.PlacedOrder,error){
	var newOrderStatus models.OrderStatus
	newOrderStatus = models.OrderStatus{StatusID:statusId,Description: description ,UserId: userId,StatusChangedTime:time.Now(),PlacedOrderID: order.ID}
	order.OrderStatusId = statusId

	placedOrderService.orderStatusRepos.Create(&newOrderStatus)

	order,err := placedOrderService.UpdatePlacedOrderStatus(order)
	if err != nil{
		return nil,err
	}
	return order,nil
}

func (placedOrderService *placedOrderService) UpdatePlacedOrderStatus(updatePlacedOrder *models.PlacedOrder) (*models.PlacedOrder, error) {
	updateRole, err := placedOrderService.placedOrderRepos.Update(updatePlacedOrder)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (placedOrderService *placedOrderService) GetPlacedOrderById(id int) (*models.PlacedOrder, error) {
	placedOrderModel, err := placedOrderService.placedOrderRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return placedOrderModel, nil
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

func (placedOrderService *placedOrderService) GetListActiveOrdersByDeliveryId(deliveryId uint, limit int, page int) (*pagination.Paginator, error){
	paginate,err := placedOrderService.placedOrderRepos.FindActivePlacedOrdersByDeliveryId(deliveryId,limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (placedOrderService *placedOrderService) GetListActivePlacedOrdersByStoreId(storeId uint) ([]*models.PlacedOrder, error){
	placedOrders,err := placedOrderService.placedOrderRepos.FindActivePlacedOrdersByStoreId(storeId)
	if err != nil {
		return nil, err
	}
	return placedOrders, nil
}

func (placedOrderService *placedOrderService) GetInStorePlacedOrdersByDeliveryId(deliveryId uint,limit int,page int)(*pagination.Paginator, error){
	paginator,err := placedOrderService.placedOrderRepos.FindInStorePlacedOrdersByDeliveryId(deliveryId,limit,page)
	if err != nil {
		return nil, err
	}
	return paginator, nil
}

