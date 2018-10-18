package orderStatus

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type OrderStatusService interface {
	CreateNewOrderStatus(newOrderStatus *models.OrderStatus)(*models.OrderStatus,error)
	GetOrderStatus(limit int, page int)(*pagination.Paginator,error)
	GetOrderStatusById(id int)(*models.OrderStatus,error)
	UpdateOrderStatus(updateOrderStatus *models.OrderStatus)(*models.OrderStatus,error)
	DeleteOrderStatus(id int)(bool,error)
}


