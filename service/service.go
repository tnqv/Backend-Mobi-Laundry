package service

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ServiceService interface {
	GetServiceDetailByName(name string) (*models.Service, error)
	CreateNewService(newService *models.Service)(*models.Service, error)
	GetServices(limit int, page int)(*pagination.Paginator, error)
	GetServiceById(id int)(*models.Service, error)
	UpdateService(updateService *models.Service)(*models.Service, error)
	DeleteService(id int)(bool, error)
}