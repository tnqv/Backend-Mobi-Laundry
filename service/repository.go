package service

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type ServiceRepository interface{
	Find(id int) (*models.Service, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(review *models.Service) (*models.Service,error)
	Update(service *models.Service) (*models.Service,error)
	Delete(id int) (bool,error)
}