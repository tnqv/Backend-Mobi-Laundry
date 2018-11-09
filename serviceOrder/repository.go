package serviceOrder

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ServiceOrderRepository interface{
	Find(id int) (*models.ServiceOrder, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(serviceOrder *models.ServiceOrder) (*models.ServiceOrder, error)
	CreateServiceOrders(serviceorders []*models.ServiceOrder)([]*models.ServiceOrder,error)
	Update(serviceOrder *models.ServiceOrder) (*models.ServiceOrder, error)
	Delete(id int) (bool, error)
}
