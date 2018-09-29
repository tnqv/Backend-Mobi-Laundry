package accounts

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
)

type Account struct {
	gorm.Model
	Email string
	Username string
	Password string
	Salt string
	Provider string
	AccessToken string
	FcmToken string
}

type Customer struct {
	gorm.Model
	Name string
	ShippingAddress string
	Longitude float32
	Latitude float32
	AvatarUrl string
	AccountID uint
	AccountInfo Account
}

type Delivery struct {
	gorm.Model
	Name string
	Capacity uint
	IdentifyNumber string
	Address string
	PhoneNumber string
	ImageUrl string
	AccountID uint
	AccountInfo Account
}

type StoreEmployee struct {
	gorm.Model
	Name string
	AccountID uint
	AccountInfo Account
}

type Store struct {
	gorm.Model
	Name string
	Description string
	Longitude float32
	Latitude float32
	Address string
	PhoneNumber string
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Delivery{})
	db.AutoMigrate(&StoreEmployee{})
	db.AutoMigrate(&Store{})

}