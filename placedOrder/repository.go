package placedOrder

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type PlacedOrderRepository interface{
	Find(id int) (*models.PlacedOrder, error)
	FindByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(placedOrder *models.PlacedOrder) (*models.PlacedOrder,error)
	Update(placedOrder *models.PlacedOrder) (*models.PlacedOrder, error)
	Delete(id int) (bool,error)
	FindPlacedOrderByOrderCode(orderCode string)(*models.PlacedOrder,error)
}
