package orderStatus

import "github.com/biezhi/gorm-paginator/pagination"

type OrderStatusService interface {
	CreateNewOrderStatus(newOrderStatus *OrderStatus)(*OrderStatus,error)
	GetOrderStatus(limit int, page int)(*pagination.Paginator,error)
	GetOrderStatusById(id int)(*OrderStatus,error)
	UpdateOrderStatus(updateOrderStatus *OrderStatus)(*OrderStatus,error)
	DeleteOrderStatus(id int)(bool,error)
}


