package placedOrder

import "github.com/biezhi/gorm-paginator/pagination"

type PlacedOrderService interface {
	CreateNewPlacedOrder(newPlacedOrder *PlacedOrder)(*PlacedOrder, error)
	GetPlacedOrders(limit int, page int)(*pagination.Paginator, error)
	GetPlacedOrderById(id int)(*PlacedOrder, error)
	UpdatePlacedOrder(updatePlacedOrder *PlacedOrder)(*PlacedOrder, error)
	DeletePlacedOrder(id int)(bool, error)
	GetListOrdersByUserId(limit int, page int, id int) (*pagination.Paginator, error)
}
