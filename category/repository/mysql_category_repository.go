package repository

import (
	"d2d-backend/category"
	"d2d-backend/common"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository() category.CategoryRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*category.Category, error) {
	var category category.Category
	err := r.Conn.First(&category,id).Error
	if err != nil {
		return nil,err
	}
	return &category, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var category []category.Category
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &category)
	return paginator,nil
}

func (r *repo) Create(category *category.Category) (*category.Category, error) {
	err := r.Conn.Create(category).Error
	if err != nil {
		return nil,err
	}
	return category,nil
}

func (r *repo) Update(updateCategory *category.Category) (*category.Category, error) {
	var tempCategory category.Category
	err := r.Conn.First(&tempCategory,updateCategory.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updateCategory).Error
	if err != nil {
		return nil, err
	}
	return updateCategory,nil
}

func (r *repo) Delete(id int) (bool, error) {
	var tempCategory category.Category
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

