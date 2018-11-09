package store

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type StoreService interface {
	GetStoreDetailByName(name string) (*models.Store,error)
	CreateNewStore(newStore *models.Store)(*models.Store,error)
	GetStores(limit int, page int)(*pagination.Paginator,error)
	GetAllStores()([]*models.Store, error)
	GetStoreById(store *models.Store)(*models.Store,error)
	UpdateStore(updateStore *models.Store)(*models.Store,error)
	DeleteStore(id int)(bool,error)
}