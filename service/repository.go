package service

import "github.com/biezhi/gorm-paginator/pagination"

type ServiceRepository interface{
	Find(id int) (*Service, error)
	FindAll(limit int, page int) (*pagination.Paginator, error)
	Create(review *Service) (*Service,error)
	Update(service *Service) (*Service,error)
	Delete(id int) (bool,error)
}