package account

import (
"d2d-backend/common"
"github.com/gin-gonic/gin"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type AccountModelValidator struct {
	Account struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"account"`
	accountModel Account `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *AccountModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.accountModel.Email = self.Account.Email


	if self.Account.Password != common.NBRandomPassword {
		self.accountModel.SetPassword(self.Account.Password)
	}
	return nil
}

// You can put the default value of a Validator here
func NewAccountModelValidator() AccountModelValidator {
	userModelValidator := AccountModelValidator{}
	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel Account) AccountModelValidator {
	userModelValidator := NewAccountModelValidator()
	userModelValidator.Account.Email = userModel.Email
	userModelValidator.Account.Password = common.NBRandomPassword
	return userModelValidator
}


type LoginValidator struct {
	Account struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"account"`
	AccountModel Account `json:"-"`
}


func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.AccountModel.Email = self.Account.Email
	return nil
}


type FBLoginValidator struct {
	Account struct {
		Provider string	`form:"provider" json:"provider" binding:"exists"`
		AccessToken string `form:"fb_access_token" json:"fb_access_token"`
	} `json:"account"`
	AccountModel Account `json:"-"`
}

func (self *FBLoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.AccountModel.AccessToken = self.Account.AccessToken
	return nil
}

type FcmTokenValidator struct {
	Account struct {
		FcmToken string `form:"fcm_token" json:"fcm_token" binding:"exists"`
	} `json:"account"`
	AccountModel Account `json:"-"`
}

func (self *FcmTokenValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.AccountModel.FcmToken = self.Account.FcmToken
	return nil
}


// You can put the default value of a Validator here
func NewFcmTokenValidator() FcmTokenValidator{
	fcmtokenValidator := FcmTokenValidator{}
	return fcmtokenValidator
}
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}

func FBNewLoginValidator() FBLoginValidator {
	fbLoginValidator := FBLoginValidator{}
	return fbLoginValidator
}