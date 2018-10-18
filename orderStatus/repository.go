package orderStatus

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type OrderStatusRepository interface{
	Find(id int) (*models.OrderStatus, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *models.OrderStatus) (*models.OrderStatus,error)
	Update(category *models.OrderStatus) (*models.OrderStatus, error)
	Delete(id int) (bool,error)
}
