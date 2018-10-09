package store

import "github.com/biezhi/gorm-paginator/pagination"

type StoreRepository interface{
	Find(store *Store) (*Store, error)
	FindByStoreName(name string) (*Store, error)
	FindAll(limit int,page int) (*pagination.Paginator, error)
	Create(store *Store) (*Store,error)
	Update(store *Store) (*Store, error)
	Delete(id int) (bool,error)
}