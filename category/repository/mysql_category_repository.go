package repository

import (
	"d2d-backend/category"
	"d2d-backend/common"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository() category.CategoryRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.Category, error) {
	var categoryModel models.Category
	err := r.Conn.First(&categoryModel,id).Error
	if err != nil {
		return nil,err
	}
	return &categoryModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var categoryModels []models.Category

	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn.Preload("Services"),
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &categoryModels)
	return paginator,nil
}

func (r *repo) Create(category *models.Category) (*models.Category, error) {
	err := r.Conn.Create(&category).Error
	if err != nil {
		return nil,err
	}
	return category,nil
}

func (r *repo) Update(updateCategory *models.Category) (*models.Category, error) {
	var tempCategory models.Category
	err := r.Conn.First(&tempCategory,updateCategory.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Model(&updateCategory).Update(&updateCategory).Error
	if err != nil {
		return nil, err
	}
	return updateCategory,nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempCategory models.Category
	err := r.Conn.First(&tempCategory,id).Error
	if err != nil {
		return false,err
	}
	err = r.Conn.Delete(&tempCategory).Error
	if err != nil {
		return false,err
	}
	return true,nil
}

