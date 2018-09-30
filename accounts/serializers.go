package accounts

import (
	"d2d-backend/common"
	"github.com/gin-gonic/gin"
)

type AccountSerializer struct {
	c *gin.Context
}

type AccountResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Token    string  `json:"token"`
}

func (self *AccountSerializer) Response() AccountResponse {
	accountModel := self.c.MustGet("user_model").(Account)
	user := AccountResponse{
		Username: accountModel.Username,
		Email:    accountModel.Email,
		Token:    common.GenToken(accountModel.ID),
	}
	return user
}