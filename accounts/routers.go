package accounts

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"d2d-backend/common"
	"errors"
	"github.com/huandu/facebook"
	"fmt"
)

const (
	FacebookProvider = "FACEBOOK"
	NormalProvider = "NORMAL"
)

func AccountsRouterRegister(router *gin.RouterGroup){
	router.POST("/login",AccountsLogin)
	router.POST("/",AccountsRegistration)
	router.POST("/facebook/auth",FacebookAccountsLogin)
}

func FacebookAccountsLogin(c *gin.Context){
	fbLoginValidator := FBNewLoginValidator()
	if err := fbLoginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if fbLoginValidator.accountModel.AccessToken == ""{
		c.JSON(http.StatusForbidden,common.NewError("fblogin",errors.New("Access token not found")))
		return
	}

	resp,err := facebook.Get("/me",facebook.Params{
		"fields": "id,first_name,last_name,picture,name,email",
		"access_token": fbLoginValidator.accountModel.AccessToken,
	})

	if err != nil {
		if e, ok := err.(*facebook.Error); ok {
			message := fmt.Sprintf("facebook error. [message:%v] [type:%v] [code:%v] [subcode:%v] [trace:%v]",
				e.Message, e.Type, e.Code, e.ErrorSubcode, e.TraceID)
			c.JSON(http.StatusForbidden,common.NewError("fblogin",errors.New(message)))
			return
		}
		return
	}

	var fbEmail string
	var fbName string
	resp.DecodeField("name",&fbName)
	err = resp.DecodeField("email",&fbEmail)

	if err != nil || fbEmail == "" {
		c.JSON(http.StatusForbidden,common.NewError("fblogin",errors.New("Facebook access token does not required email scope")))
		return
	}

	accountModel,err := FindOneUser(&Account{Email: fbEmail})

	if err != nil {
		var newAccount Account
		newAccount.Email = fbEmail
		newAccount.Username = fbName
		newAccount.Provider = FacebookProvider
		newAccount.AccessToken = fbLoginValidator.accountModel.AccessToken

		if err := SaveOne(&newAccount); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}

		c.Set("user_model", newAccount)
		serializer := AccountSerializer{c}
		c.JSON(http.StatusCreated, gin.H{"account": serializer.Response()})
	}else{
		c.Set("user_model", accountModel)
		serializer := AccountSerializer{c}
		c.JSON(http.StatusOK, gin.H{"account": serializer.Response()})
	}
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