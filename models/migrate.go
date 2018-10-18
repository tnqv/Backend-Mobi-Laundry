package models

import "d2d-backend/common"

func AutoMigrate(){
	db := common.GetDB()
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Notification{})
	db.AutoMigrate(&OrderStatus{})
	db.AutoMigrate(&PlacedOrder{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&ServiceOrder{})
	db.AutoMigrate(&Store{})
	db.AutoMigrate(&User{})
}