package repository

import (
	"d2d-backend/store"
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
	"errors"
	"github.com/biezhi/gorm-paginator/pagination"
)

type repo struct {
	Conn *gorm.DB
}

func NewMysqlStoreRepository() store.StoreRepository {

	return &repo{common.GetDB()}
}

func (r *repo) Find(store *store.Store) (*store.Store, error) {

	err := r.Conn.First(&store,store.ID).Error

	if err != nil {
		 return nil,err
	}

	return store, nil
}

func (r *repo) FindByStoreName(name string) (*store.Store, error){
	if name == ""{
		return nil, errors.New("Invalid name")
	}

	var store store.Store

	err := r.Conn.First(&store,"name = ?",name).Error

	if err != nil {
		return nil,err
	}

	return &store,nil
}


func (r *repo) FindAll(limit int,page int) (*pagination.Paginator, error){

	var stores []*store.Store

	paginator := pagination.Pagging(&pagination.Param{
		DB: r.Conn,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &stores)

	return paginator,nil
}

func (r *repo) Create(store *store.Store) (*store.Store,error){
	err := r.Conn.Create(store).Error
	if err != nil {
		return nil,err
	}

	return store,nil
}

func (r *repo) Update(storeUpdate *store.Store) (*store.Store, error){

	var storeTemp store.Store

	err := r.Conn.First(&storeTemp,storeUpdate.ID).Error

	if err != nil{
		return nil,err
	}

	if storeUpdate.Name != ""{
		storeTemp.Name = storeUpdate.Name
	}

	if storeUpdate.Description != ""{
		storeTemp.Description = storeUpdate.Description
	}

	if storeUpdate.Address != ""{
		storeTemp.Address = storeUpdate.Address

	}
	if storeUpdate.PhoneNumber != ""{
		storeTemp.PhoneNumber = storeUpdate.PhoneNumber
	}

	if storeUpdate.Latitude != 0 {
		storeTemp.Latitude = storeUpdate.Latitude
	}

	if storeUpdate.Longitude != 0 {
		storeTemp.Longitude = storeUpdate.Longitude
	}

	err = r.Conn.Model(&store.Store{}).Save(&storeTemp).Error

	if err != nil {
		return nil,err
	}

	return &storeTemp,nil
}
func (r *repo) Delete(id int) (bool,error){
	var storeTemp store.Store

	err := r.Conn.First(&storeTemp,id).Error

	if err != nil {
		return false,err
	}

	err = r.Conn.Delete(&storeTemp).Error

	if err != nil {
		return false,err
	}

	return true,nil
}