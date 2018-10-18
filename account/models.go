package account

import (
	"github.com/jinzhu/gorm"
	"d2d-backend/common"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type Account struct {
	gorm.Model			`json:"-"`
	Email 		string	`gorm:"not null;unique" form:"email" json:"email"`
	Username 	string	`form:"username" json:"username"`
	Password 	string	`form:"password" json:"password"`
	Provider 	string	`form:"provider" json:"provider"`
	AccessToken string	`form:"access_token" json:"access_token"`
	FcmToken 	string	`form:"fcm_token" json:"fcm_token"`
	ApnToken 	string	`form:"apn_token" json:"apn_token"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Account{})
}


// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *Account) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *Account) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}