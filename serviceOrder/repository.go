package serviceOrder

import "github.com/biezhi/gorm-paginator/pagination"

type ServiceOrderRepository interface{
	Find(id int) (*ServiceOrder, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(serviceOrder *ServiceOrder) (*ServiceOrder, error)
	Update(serviceOrder *ServiceOrder) (*ServiceOrder, error)
	Delete(id int) (bool, error)
}
