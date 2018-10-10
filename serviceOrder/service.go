package serviceOrder

import "github.com/biezhi/gorm-paginator/pagination"

type ServiceOrderService interface {
	CreateNewServiceOrder(newServiceOrder *ServiceOrder) (*ServiceOrder, error)
	GetServiceOrders(limit int, page int) (*pagination.Paginator, error)
	GetServiceOrderById(id int) (*ServiceOrder, error)
	UpdateServiceOrder(updateServiceOrder *ServiceOrder) (*ServiceOrder, error)
	DeleteServiceOrder(id int) (bool, error)
}