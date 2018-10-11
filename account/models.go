package account

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model			`json:"-"`
	Email 		string	`form:"email" json:"email"`
	Username 	string	`form:"username" json:"username"`
	Password 	string	`form:"password" json:"password"`
	Provider 	string	`form:"provider" json:"provider"`
	AccessToken string	`form:"access_token" json:"access_token"`
	FcmToken 	string	`form:"fcm_token" json:"fcm_token"`
	ApnToken 	string	`form:"apn_token" json:"apn_token"`
}
