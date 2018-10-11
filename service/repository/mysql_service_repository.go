package repository

import (
	"d2d-backend/common"
	"d2d-backend/service"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlServiceRepository() service.ServiceRepository {
	return &repo{common.GetDB()}
}

func (r *repo) Find(id int) (*service.Service, error) {
	var service service.Service
	err := r.Conn.First(&service, id).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *repo) FindAll(limit int, page int) (*pagination.Paginator, error) {
	var services []service.Service
	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &services)
	return paginator, nil
}

func (r *repo) Create(service *service.Service) (*service.Service, error) {
	err := r.Conn.Create(service).Error
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *repo) Update(updatedService *service.Service) (*service.Service, error) {
	var serviceTemp service.Service
	err := r.Conn.First(&serviceTemp, updatedService.ID).Error
	if err != nil{
		return nil, err
	}
	err = r.Conn.Save(updatedService).Error
	if err != nil {
		return nil, err
	}
	return updatedService, nil
}

func (r *repo) Delete(id int) (bool, error) {
	var serviceTemp service.Service
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
