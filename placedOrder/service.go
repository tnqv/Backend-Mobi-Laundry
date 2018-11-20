package placedOrder

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type PlacedOrderService interface {
	CreateNewPlacedOrder(newPlacedOrder *models.PlacedOrder)(*models.PlacedOrder, error)
	GetPlacedOrders(limit int, page int)(*pagination.Paginator, error)
	GetPlacedOrderById(id int)(*models.PlacedOrder, error)
	UpdatePlacedOrder(updatePlacedOrder *models.PlacedOrder)(*models.PlacedOrder, error)
	DeletePlacedOrder(id int)(bool, error)
	GetListOrdersByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	GetListActiveOrdersByDeliveryId(deliveryId uint,limit int, page int) (*pagination.Paginator, error)
	GetListActivePlacedOrdersByStoreId(storeId uint) ([]*models.PlacedOrder, error)
	GetPlacedOrderByOrderCode(orderCode string)(*models.PlacedOrder,error)
	GetInStorePlacedOrdersByDeliveryId(deliveryId uint,limit int,page int)(*pagination.Paginator, error)
	UpdatePlacedOrderAndCreateNewOrderStatus(statusId uint,userId uint,description string,order *models.PlacedOrder)(*models.PlacedOrder,error)
}
