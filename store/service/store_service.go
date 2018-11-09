package service

import (
	"d2d-backend/store"
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type storeService struct {
	storeRepos   store.StoreRepository
}

func NewStoreService(storeRepository store.StoreRepository) store.StoreService {
	return &storeService{storeRepository}
}

func (storeService *storeService) GetStoreById(store *models.Store)(*models.Store,error) {
		storeModel,err := storeService.storeRepos.Find(store)

		if err != nil {
			return nil,err
		}

		return storeModel,nil
}

func (storeService *storeService) GetAllStores()([]*models.Store, error){
	stores,err := storeService.storeRepos.FindAllStore()
	if err != nil {
		return nil,err
	}
	return stores,nil
}

func (storeService *storeService) GetStores(limit int, page int)(*pagination.Paginator,error){
	paginate,err := storeService.storeRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (storeService *storeService) GetStoreDetailByName(name string) (*models.Store,error) {
	return nil,nil
}

func (storeService *storeService)  CreateNewStore(newStore *models.Store)(*models.Store,error){
	_,err := storeService.storeRepos.Create(newStore)

	if err != nil {
		return nil,err
	}
	return newStore,nil
}

func (storeService *storeService) UpdateStore(updateStore *models.Store)(*models.Store,error){
	updateStore,err := storeService.storeRepos.Update(updateStore)

	if err != nil {
		return nil,err
	}

	return updateStore,nil
}

func (storeService *storeService) DeleteStore(id int)(bool,error) {
	isDeletedSuccess,err := storeService.storeRepos.Delete(id)

	if err != nil {
		return isDeletedSuccess,err
	}

	return isDeletedSuccess,nil


}