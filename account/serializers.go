package account

import (
	"github.com/gin-gonic/gin"
	"d2d-backend/models"
	"d2d-backend/common"
)

type DriverAccountSerializer struct {
	C *gin.Context
}

type DriverAccountResponse struct {
	Username    string  `json:"username"`
	Token    string  `json:"token"`
	ID		 uint	  `json:"-"`
}

func (self *DriverAccountSerializer) Response() *DriverAccountResponse {
	accountModel := self.C.MustGet("user_model").(*models.Account)
	//userModel,err := self.userService.GetUserByAccountId(accountModel.ID)

	//if err != nil {
	//	return nil
	//}

	user := DriverAccountResponse{
		Username:    accountModel.Username,
		Token:    common.GenToken(accountModel.ID),
	}
	return &user
}

type AccountSerializer struct {
	C *gin.Context
}

type AccountResponse struct {
	Email    string  `json:"email"`
	Token    string  `json:"token"`
	ID		 uint	  `json:"-"`
}

func (self *AccountSerializer) Response() *AccountResponse {
	accountModel := self.C.MustGet("user_model").(*models.Account)
	//userModel,err := self.userService.GetUserByAccountId(accountModel.ID)

	//if err != nil {
	//	return nil
	//}

	user := AccountResponse{
		Email:    accountModel.Email,
		Token:    common.GenToken(accountModel.ID),
	}
	return &user
}