package orderStatus

import "github.com/biezhi/gorm-paginator/pagination"

type OrderStatusRepository interface{
	Find(id int) (*OrderStatus, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *OrderStatus) (*OrderStatus,error)
	Update(category *OrderStatus) (*OrderStatus, error)
	Delete(id int) (bool,error)
}
