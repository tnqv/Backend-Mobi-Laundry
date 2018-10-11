package service

import (
	"d2d-backend/category"
	"github.com/biezhi/gorm-paginator/pagination"
)

type categoryService struct {
	categoryRepos	category.CategoryRepository
}

func NewCategoryService(categoryRepository category.CategoryRepository) category.CategoryService {
	return &categoryService{categoryRepository}
}

func (categoryService *categoryService) GetCategoryDetailByName(name string) (*category.Category, error) {
	panic("implement me")
}

func (categoryService *categoryService) CreateNewCategory(newCategory *category.Category) (*category.Category, error) {
	_,err := categoryService.categoryRepos.Create(newCategory)
	if err != nil {
		return nil,err
	}
	return newCategory,nil
}

func (categoryService *categoryService) GetCategory(limit int, page int) (*pagination.Paginator, error) {
	paginate,err := categoryService.categoryRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (categoryService *categoryService) GetCategoryById(id int) (*category.Category, error) {
	category,err := categoryService.categoryRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return category,nil
}

func (categoryService *categoryService) UpdateCategory(updateCategory *category.Category) (*category.Category, error) {
	updateCategory,err := categoryService.categoryRepos.Update(updateCategory)
	if err != nil {
		return nil,err
	}
	return updateCategory,nil
}

func (categoryService *categoryService) DeleteCategory(id int) (bool, error) {
	bool,err := categoryService.categoryRepos.Delete(id)
	if err != nil {
		return bool,err
	}
	return bool,nil
}

