package accounts

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type Account struct {
	gorm.Model
	Email string		`form:"email"`
	Username string		`form:"username"`
	Password string		`form:"password"`
	Salt string			`form:"-"`
	Provider string		`form:"provider"`
	AccessToken string	`form:"access_token"`
	FcmToken string		`form:"fcm_token"`
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

func FindOneUser(condition interface{}) (Account, error) {
	db := common.GetDB()
	var model Account
	err := db.Where(condition).First(&model).Error
	return model, err
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *Account) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *Account) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// You could input an UserModel which will be saved in database returning with error info
// 	if err := SaveOne(&userModel); err != nil { ... }
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}