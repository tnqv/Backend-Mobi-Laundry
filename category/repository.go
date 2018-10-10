package category

import "github.com/biezhi/gorm-paginator/pagination"

type CategoryRepository interface{
	Find(id int) (*Category, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(category *Category) (*Category,error)
	Update(category *Category) (*Category, error)
	Delete(id int) (bool,error)
}
