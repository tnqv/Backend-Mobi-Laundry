package accounts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"d2d-backend/common"
	"errors"
)

func AccountsRegister(router *gin.RouterGroup){
	router.POST("/login",AccountsLogin)
	router.POST("/",AccountsRegistration)
}

func AccountsLogin(c *gin.Context){
		loginValidator := NewLoginValidator()
		if err := loginValidator.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
			return
		}

		accountModel,err := FindOneUser(&Account{Email: loginValidator.accountModel.Email})

		if err != nil {
			c.JSON(http.StatusForbidden,common.NewError("login",errors.New("Username or password is not valid")))
			return
		}

		if accountModel.checkPassword(loginValidator.Account.Password) != nil {
			c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
			return
		}
		UpdateContextUserModel(c, accountModel.ID)
		serializer := AccountSerializer{c}
		c.JSON(http.StatusOK, gin.H{"account": serializer.Response()})
}


func AccountsRegistration(c *gin.Context) {
	accountModelValidator := NewAccountModelValidator()
	if err := accountModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&accountModelValidator.accountModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("user_model", accountModelValidator.accountModel)
	serializer := AccountSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"account": serializer.Response()})
}