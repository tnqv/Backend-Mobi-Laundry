package account

import (
"d2d-backend/common"
"github.com/gin-gonic/gin"
)

type AccountSerializer struct {
	C *gin.Context
}

type AccountResponse struct {
	Email    string  `json:"email"`
	Token    string  `json:"token"`
	ID		 uint	  `json:"-"`
}

func (self *AccountSerializer) Response() AccountResponse {
	accountModel := self.C.MustGet("user_model").(Account)
	user := AccountResponse{
		Email:    accountModel.Email,
		Token:    common.GenToken(accountModel.ID),
		ID:		  accountModel.ID,
	}
	return user
}