package accounts

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
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
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
	self.accountModel.Username = self.Account.Username
	self.accountModel.Email = self.Account.Email


	if self.Account.Password != common.NBRandomPassword {
		self.accountModel.setPassword(self.Account.Password)
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
	userModelValidator.Account.Username = userModel.Username
	userModelValidator.Account.Email = userModel.Email
	userModelValidator.Account.Password = common.NBRandomPassword
	return userModelValidator
}


type LoginValidator struct {
	Account struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password"json:"password" binding:"exists,min=8,max=255"`
	} `json:"account"`
	accountModel Account `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.accountModel.Email = self.Account.Email
	return nil
}

// You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}