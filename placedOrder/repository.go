package placedOrder

import (
	"d2d-backend/models"
	"d2d-backend/common"
)

type PlacedOrderRepository interface{
	Find(id int) (*models.PlacedOrder, error)
	FindByUserId(limit int, page int, id int) (*common.Paginator, error)
	FindAll(limit int,page int) (*common.Paginator, error)
	Create(placedOrder *models.PlacedOrder) (*models.PlacedOrder,error)
	Update(placedOrder *models.PlacedOrder) (*models.PlacedOrder, error)
	UpdateOrderStatusId(placedOrder *models.PlacedOrder) (*models.PlacedOrder, error)
	Delete(id int) (bool,error)
	FindActivePlacedOrdersByDeliveryId(deliveryId uint,limit int,page int)(*common.Paginator, error)
	FindActivePlacedOrdersByStoreId(storeId uint)([]*models.PlacedOrder,error)
	FindInStorePlacedOrdersByDeliveryId(deliveryId uint,limit int,page int)(*common.Paginator, error)
	FindPlacedOrderByOrderCode(orderCode string)(*models.PlacedOrder,error)
}
