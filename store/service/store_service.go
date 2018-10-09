package service

import (
	"d2d-backend/store"
)

type storeService struct {
	storeRepos   store.StoreRepository
}

func NewStoreService(storeRepository store.StoreRepository) store.StoreService {
	return &storeService{storeRepository}
}

func (storeService *storeService) GetStoreDetailByName(name string) (*store.Store,error) {
	return nil,nil
}

func (storeService *storeService)  CreateNewStore(newStore *store.Store)(*store.Store,error){
	return nil,nil
}