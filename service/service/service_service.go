package service

import (
	"d2d-backend/service"
	"github.com/biezhi/gorm-paginator/pagination"
)

type serviceService struct {
	serviceRepos service.ServiceRepository
}

func NewServiceService(serviceRepository service.ServiceRepository) service.ServiceService {
	return &serviceService{serviceRepository}
}

func (serviceService *serviceService) GetServiceById(id string) (*service.Service, error) {
	service, err := serviceService.serviceRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (serviceService *serviceService) GetServiceDetailByName(name string) (*service.Service, error) {
	return nil, nil
}

func (serviceService *serviceService) CreateNewService(newService *service.Service) (*service.Service, error) {
	_, err := serviceService.serviceRepos.Create(newService)
	if err != nil {
		return nil, err
	}
	return newService, nil
}

func (serviceService *serviceService) GetServices(limit int, page int) (*pagination.Paginator, error) {
	paginate, err := serviceService.serviceRepos.FindAll(limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (serviceService *serviceService) UpdateService(updateService *service.Service) (*service.Service, error) {
	updateService, err := serviceService.serviceRepos.Update(updateService)
	if err != nil {
		return nil, err
	}
	return updateService, nil
}

func (serviceService *serviceService) DeleteService(id int) (bool, error) {
	bool,err := serviceService.serviceRepos.Delete(id)
	if err != nil {
		return bool,err
	}
	return bool,nil
}

