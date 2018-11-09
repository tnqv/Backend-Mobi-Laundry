package service

import (
	"d2d-backend/serviceOrder"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type serviceOrderService struct {
	serviceOrderRepos serviceOrder.ServiceOrderRepository
}

func NewServiceOrderService(serviceOrderRepository serviceOrder.ServiceOrderRepository) serviceOrder.ServiceOrderService {
	return &serviceOrderService{serviceOrderRepository}
}

func (serviceOrderService *serviceOrderService) CreateListServiceOrders(newServiceOrders []*models.ServiceOrder) ([]*models.ServiceOrder, error) {
	newServiceOrders,err := serviceOrderService.serviceOrderRepos.CreateServiceOrders(newServiceOrders)
	if err != nil {
		return nil,err
	}

	return newServiceOrders,nil
}

func (serviceOrderService *serviceOrderService) CreateNewServiceOrder(newServiceOrder *models.ServiceOrder) (*models.ServiceOrder, error) {
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

func (serviceOrderService *serviceOrderService) GetServiceOrderById(id int) (*models.ServiceOrder, error) {
	serviceOrderModel, err := serviceOrderService.serviceOrderRepos.Find(id)
	if err != nil {
		return nil, err
	}
	return serviceOrderModel, nil
}

func (serviceOrderService *serviceOrderService) UpdateServiceOrder(updateServiceOrder *models.ServiceOrder) (*models.ServiceOrder, error) {
	updateServiceOrder, err := serviceOrderService.serviceOrderRepos.Update(updateServiceOrder)
	if err != nil {
		return nil, err
	}
	return updateServiceOrder, nil
}

func (serviceOrderService *serviceOrderService) DeleteServiceOrder(id int) (bool, error) {
	isDeletedSuccess, err := serviceOrderService.serviceOrderRepos.Delete(id)
	if err != nil {
		return isDeletedSuccess, err
	}
	return isDeletedSuccess, nil
}
