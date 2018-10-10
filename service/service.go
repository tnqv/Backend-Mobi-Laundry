package service

import "github.com/biezhi/gorm-paginator/pagination"

type ServiceService interface {
	GetServiceDetailByName(name string) (*Service, error)
	CreateNewService(newService *Service)(*Service, error)
	GetServices(limit int, page int)(*pagination.Paginator, error)
	GetServiceById(id string)(*Service, error)
	UpdateService(updateService *Service)(*Service, error)
	DeleteService(id int)(bool, error)
}