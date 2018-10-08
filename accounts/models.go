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
	Password string		`form:"password"`
	Salt string			`form:"-"`
	Provider string		`form:"provider"`
	AccessToken string	`form:"access_token"`
	FcmToken string		`form:"fcm_token"`
}

type User struct {
	gorm.Model			`json:"-"`
	Name string			`json:"name"`
	Address string		`json:"-"`
	PhoneNumber string	`json:"phone_number"`
	Longitude float32	`json:"-"`
	Latitude float32	`json:"-"`
	RoleID uint			`json:"-"`
	AvatarUrl string	`json:"-"`
	AccountID uint		`json:"-"`
	AccountInfo Account	`json:"-"`
}

//Role : 1 : Customer
//       2 : Delivery
//		 3 : StoreEmployeee
type Role struct {
	gorm.Model
	Name string
	Description string
}

type Store struct {
	gorm.Model				`json:"-"`
	Name string				`json:"name"`
	Description string		`json:"description"`
	Longitude float32		`json:"longitude"`
	Latitude float32		`json:"latitude"`
	Address string			`json:"address"`
	PhoneNumber string		`json:"phone_number"`
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
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

// get user information by accountID
func GetUserInformations(accountID uint) (User) {
	db := common.GetDB()
	var user User
	db.Find(&user, "account_id = ?", accountID)
	return user
}

func CreateNewUser(user *User) error {
	db := common.GetDB()
	err := db.Save(user).Error
	return err
}

