package store

import "github.com/biezhi/gorm-paginator/pagination"

type StoreService interface {
	GetStoreDetailByName(name string) (*Store,error)
	CreateNewStore(newStore *Store)(*Store,error)
	GetStores(limit int, page int)(*pagination.Paginator,error)
	GetStoreById(store *Store)(*Store,error)
	UpdateStore(updateStore *Store)(*Store,error)
	DeleteStore(id int)(bool,error)
}

