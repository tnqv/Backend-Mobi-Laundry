package service

import (
	"d2d-backend/store"
	"github.com/biezhi/gorm-paginator/pagination"
)

type storeService struct {
	storeRepos   store.StoreRepository
}

func NewStoreService(storeRepository store.StoreRepository) store.StoreService {
	return &storeService{storeRepository}
}

func (storeService *storeService) GetStoreById(store *store.Store)(*store.Store,error) {
		store,err := storeService.storeRepos.Find(store)

		if err != nil {
			return nil,err
		}

		return store,nil
}

func (storeService *storeService) GetStores(limit int, page int)(*pagination.Paginator,error){
	paginate,err := storeService.storeRepos.FindAll(limit,page)
	if err != nil {
		return nil,err
	}
	return paginate,nil
}

func (storeService *storeService) GetStoreDetailByName(name string) (*store.Store,error) {
	return nil,nil
}

func (storeService *storeService)  CreateNewStore(newStore *store.Store)(*store.Store,error){
	_,err := storeService.storeRepos.Create(newStore)

	if err != nil {
		return nil,err
	}
	return newStore,nil
}

func (storeService *storeService) UpdateStore(updateStore *store.Store)(*store.Store,error){
	updateStore,err := storeService.storeRepos.Update(updateStore)

	if err != nil {
		return nil,err
	}

	return updateStore,nil
}

func (storeService *storeService) DeleteStore(id int)(bool,error) {
	bool,err := storeService.storeRepos.Delete(id)

	if err != nil {
		return bool,err
	}

	return bool,nil


}