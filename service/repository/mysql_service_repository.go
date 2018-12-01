package repository

import (
	"d2d-backend/common"
	"d2d-backend/service"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"d2d-backend/models"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlServiceRepository() service.ServiceRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*models.Service, error) {
	var serviceModel models.Service
	err := r.Conn.First(&serviceModel, id).Error
	if err != nil {
		return nil, err
	}
	return &serviceModel, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var services []models.Service
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &services)
	return paginator, nil
}

func (r *repo) Create(service *models.Service) (*models.Service, error) {
	service.DeletedAt = nil
	err := r.Conn.Create(service).Error
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *repo) Update(updatedService *models.Service) (*models.Service, error) {
	var serviceTemp models.Service
	err := r.Conn.First(&serviceTemp, updatedService.ID).Error
	if err != nil{
		return nil, err
	}
	updatedService.DeletedAt = nil
	err = r.Conn.Model(&updatedService).Update(&updatedService).Error
	if err != nil {
		return nil, err
	}
	return updatedService, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var serviceTemp models.Service
	err := r.Conn.First(&serviceTemp, id).Error
	if err != nil {
		return false, err
	}
	err = r.Conn.Delete(&serviceTemp).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
