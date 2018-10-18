package serviceOrder

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ServiceOrderService interface {
	CreateNewServiceOrder(newServiceOrder *models.ServiceOrder) (*models.ServiceOrder, error)
	GetServiceOrders(limit int, page int) (*pagination.Paginator, error)
	GetServiceOrderById(id int) (*models.ServiceOrder, error)
	UpdateServiceOrder(updateServiceOrder *models.ServiceOrder) (*models.ServiceOrder, error)
	DeleteServiceOrder(id int) (bool, error)
}