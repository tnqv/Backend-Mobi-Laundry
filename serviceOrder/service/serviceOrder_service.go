package service

import (
	"d2d-backend/serviceOrder"
	"github.com/biezhi/gorm-paginator/pagination"
)

type serviceOrderService struct {
	serviceOrderRepos serviceOrder.ServiceOrderRepository
}

func NewServiceOrderService(serviceOrderRepository serviceOrder.ServiceOrderRepository) serviceOrder.ServiceOrderService {
	return &serviceOrderService{serviceOrderRepository}
}

func (serviceOrderService *serviceOrderService) CreateNewServiceOrder(newServiceOrder *serviceOrder.ServiceOrder) (*serviceOrder.ServiceOrder, error) {
	_, err := serviceOrderService.serviceOrderRepos.Create(newServiceOrder)
	if err != nil {
		return nil, err
	}
	return newServiceOrder, nil
}

func (serviceOrderService *serviceOrderService) GetServiceOrders(limit int, page int) (*pagination.Paginator, error) {
	paginate, err := serviceOrderService.serviceOrderRepos.FindAll(limit, page)
	if err != nil {
		return nil, err
	}
	return paginate, nil
}

func (serviceOrderService *serviceOrderService) GetServiceOrderById(id int) (*serviceOrder.ServiceOrder, error) {
	serviceOrder, err := serviceOrderService.serviceOrderRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return serviceOrder, nil
}

func (serviceOrderService *serviceOrderService) UpdateServiceOrder(updateServiceOrder *serviceOrder.ServiceOrder) (*serviceOrder.ServiceOrder, error) {
	updateServiceOrder, err := serviceOrderService.serviceOrderRepos.Update(updateServiceOrder)
	if err != nil {
		return nil, err
	}
	return updateServiceOrder, nil
}

func (serviceOrderService *serviceOrderService) DeleteServiceOrder(id int) (bool, error) {
	bool, err := serviceOrderService.serviceOrderRepos.Delete(id)
	if err != nil {
		return bool, err
	}
	return bool, nil
}
