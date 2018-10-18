package store

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"d2d-backend/models"
)

type StoreRepository interface{
	Find(store *models.Store) (*models.Store, error)
	FindByStoreName(name string) (*models.Store, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(store *models.Store) (*models.Store, error)
	Update(store *models.Store) (*models.Store, error)
	Delete(id int) (bool, error)
}