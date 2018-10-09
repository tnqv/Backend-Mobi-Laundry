package repository

import (
	"d2d-backend/store"
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlStoreRepository() store.StoreRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(id string) (*store.Store, error) {

	return nil, nil
}

func (r *repo) FindByStoreName(name string) (*store.Store, error){
	return nil,nil
}


func (r *repo) FindAll() ([]*store.Store, error){
	return nil,nil
}

func (r *repo) Create(store *store.Store) (string,error){
	return "",nil
}

func (r *repo) Update(store *store.Store) error{
	return nil
}
func (r *repo) Delete(id int) (bool,error){
	return false,nil
}