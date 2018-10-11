package placedOrder

import "github.com/biezhi/gorm-paginator/pagination"

type PlacedOrderRepository interface{
	Find(id int) (*PlacedOrder, error)
	FindByUserId(limit int, page int, id int) (*pagination.Paginator, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(placedOrder *PlacedOrder) (*PlacedOrder,error)
	Update(placedOrder *PlacedOrder) (*PlacedOrder, error)
	Delete(id int) (bool,error)
}
