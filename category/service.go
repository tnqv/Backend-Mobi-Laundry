package category

import "github.com/biezhi/gorm-paginator/pagination"

type CategoryService interface {
	GetCategoryDetailByName(name string) (*Category,error)
	CreateNewCategory(newCategory *Category)(*Category,error)
	GetCategory(limit int, page int)(*pagination.Paginator,error)
	GetCategoryById(id int)(*Category,error)
	UpdateCategory(updateCategory *Category)(*Category,error)
	DeleteCategory(id int)(bool,error)
}


