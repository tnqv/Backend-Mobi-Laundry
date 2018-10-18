package service

import (
	"d2d-backend/category"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type categoryService struct {
	categoryRepos	category.CategoryRepository
}

func NewCategoryService(categoryRepository category.CategoryRepository) category.CategoryService {
	return &categoryService{categoryRepository}
}

func (categoryService *categoryService) GetCategoryDetailByName(name string) (*models.Category, error) {
	panic("implement me")
}

func (categoryService *categoryService) CreateNewCategory(newCategory *models.Category) (*models.Category, error) {
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

func (categoryService *categoryService) GetCategoryById(id int) (*models.Category, error) {
	categoryModel,err := categoryService.categoryRepos.Find(id)
	if err != nil {
		return nil,err
	}
	return categoryModel,nil
}

func (categoryService *categoryService) UpdateCategory(updateCategory *models.Category) (*models.Category, error) {
	updateCategory,err := categoryService.categoryRepos.Update(updateCategory)
	if err != nil {
		return nil,err
	}
	return updateCategory,nil
}

func (categoryService *categoryService) DeleteCategory(id int) (bool, error) {
	isDeletedSuccess,err := categoryService.categoryRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess,err
	}
	return isDeletedSuccess,nil
}

