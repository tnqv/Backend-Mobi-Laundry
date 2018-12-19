package handler

import (
	"d2d-backend/account"
	"d2d-backend/common"
	middlewares "d2d-backend/middlewares"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"d2d-backend/user"
	"github.com/huandu/facebook"
	"fmt"
	"d2d-backend/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpAccountHandler struct {
	accountService account.AccountService
	userService user.UserService
}

func NewAccountHttpHandler(e *gin.RouterGroup,
						   service account.AccountService,
						   	uService user.UserService) (*HttpAccountHandler){
	handler := &HttpAccountHandler{
		accountService: service,
		userService: uService,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpAccountHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.POST("/", s.CreateAccount)
	e.POST("/store", s.CreateAccountStore)
	e.POST("/delivery", s.CreateAccountDelivery)
	e.POST("/login",s.AccountsLogin)
	e.POST("/facebook/auth",s.FacebookAccountsLogin)
	e.POST("/driver/login",s.DriverLogin)
	e.POST("/store/login",s.StoreLogin)
}

func (s *HttpAccountHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllAccounts)
	e.GET("/:id", s.GetAccountById)
	e.PUT("/:id",s.UpdateAccount)
	e.DELETE("/:id", s.DeleteAccount)
	e.PUT("/:id/token/refresh",s.FcmTokenUpdate)

}

func (s *HttpAccountHandler) GetAllAccounts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	list, err := s.accountService.GetAccounts(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func (s *HttpAccountHandler) GetAccountById(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	account, err := s.accountService.GetAccountById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, account)
}
//
//func (s *HttpAccountHandler) CreateDriverAccount(c *gin.Context){
//
//	driverValidator := account.NewDriverLoginValidator()
//	if err := driverValidator.Bind(c); err != nil {
//		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
//		return
//	}
//
//	if driverValidator.Account.Username == ""{
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo Email")))
//		return
//	}
//
//	if driverValidator.Account.Password == ""{
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo password")))
//		return
//	}
//
//	driverValidator.AccountModel.Provider = common.NormalProvider
//	driverValidator.AccountModel.DeletedAt = nil
//
//	if driverValidator.AccountModel.Password != common.NBRandomPassword {
//		driverValidator.AccountModel.SetPassword(driverValidator.AccountModel.Password)
//	}
//
//	var user models.User
//	user.Name = c.PostForm("name")
//	if user.Name == ""{
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tên")))
//		return
//	}
//	user.PhoneNumber = c.PostForm("phone_number")
//	if user.PhoneNumber == ""{
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Số điện thoại không hợp lệ")))
//		return
//	}
//
//	userTemp := &models.User{PhoneNumber:""}
//	accountTemp := &models.Account{Username: ""}
//
//	userTemp,_ = s.userService.GetUserByPhoneNumber(user.PhoneNumber)
//
//	if userTemp != nil{
//		c.JSON(http.StatusForbidden, common.NewError("database", errors.New("Số điện thoại đã bị trùng")))
//		return
//	}
//
//	accountTemp,_ = s.accountService.GetAccountByUsername(driverValidator.Account.Username)
//
//	if accountTemp != nil{
//		c.JSON(http.StatusForbidden,common.NewError("database",errors.New("Tài khoản đã bị trùng")))
//		return
//	}
//
//	_, err = s.accountService.CreateNewAccount(&accountModel)
//	if err != nil {
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
//		return
//	}
//
//	user.AccountId = accountModel.ID
//	user.RoleId = 1
//
//	_,err = s.userService.CreateNewUser(&user)
//	if err != nil {
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
//		return
//	}
//
//	c.Set("user_model", accountModel)
//
//	serializer := account.AccountSerializer{c}
//	c.JSON(http.StatusOK, gin.H{
//		"account": serializer.Response(),
//		"user": user,
//	})
//
//
//
//
//	c.JSON(http.StatusOK, gin.H{
//		"account": serializer.Response(),
//		"user": userRequested ,
//	})
//
//}

func (s *HttpAccountHandler) CreateAccount(c *gin.Context){
	var accountModel models.Account
	err := common.Bind(c, &accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}

	if accountModel.Email == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tài khoản ")))
		return
	}

	if accountModel.Password == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo mật khẩu")))
		return
	}

	accountModel.Provider = common.NormalProvider
	accountModel.DeletedAt = nil

	if accountModel.Password != common.NBRandomPassword {
		accountModel.SetPassword(accountModel.Password)
	}

	var user models.User
	user.Name = c.PostForm("name")
	if user.Name == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tên")))
		return
	}
	user.PhoneNumber = c.PostForm("phone_number")
	if user.PhoneNumber == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Số điện thoại không hợp lệ")))
		return
	}

	userTemp := &models.User{PhoneNumber:""}
	accountTemp := &models.Account{Email: ""}

	userTemp,_ = s.userService.GetUserByPhoneNumber(user.PhoneNumber)

	if userTemp != nil{
		c.JSON(http.StatusForbidden, common.NewError("database", errors.New("Số điện thoại đã bị trùng")))
		return
	}

	accountTemp,_ = s.accountService.GetAccountByEmail(accountModel.Email)

	if accountTemp != nil{
		c.JSON(http.StatusForbidden,common.NewError("database",errors.New("Email đã bị trùng")))
		return
	}

	_, err = s.accountService.CreateNewAccount(&accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	user.AccountId = accountModel.ID
	user.RoleId = 1

	_,err = s.userService.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("user_model", &accountModel)

	serializer := account.AccountSerializer{c}
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": user,
	})
}


func (s *HttpAccountHandler) CreateAccountDelivery(c *gin.Context){
	var accountModel models.Account
	err := common.Bind(c, &accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}

	if accountModel.Username == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tài khoản")))
		return
	}

	if accountModel.Password == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo mật khẩu")))
		return
	}

	accountModel.Provider = common.NormalProvider
	accountModel.DeletedAt = nil

	if accountModel.Password != common.NBRandomPassword {
		accountModel.SetPassword(accountModel.Password)
	}

	var user models.User
	user.Name = c.PostForm("name")
	if user.Name == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tên")))
		return
	}
	user.PhoneNumber = c.PostForm("phone_number")
	if user.PhoneNumber == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Số điện thoại không hợp lệ")))
		return
	}

	user.IdentifyNumber = c.PostForm("identify_number")
	user.Address = c.PostForm("address")
	user.Imageurl = c.PostForm("imageurl")

	userTemp := &models.User{PhoneNumber:""}
	accountTemp := &models.Account{Username: ""}

	userTemp,_ = s.userService.GetUserByPhoneNumber(user.PhoneNumber)

	if userTemp != nil{
		c.JSON(http.StatusForbidden, common.NewError("database", errors.New("Số điện thoại đã bị trùng")))
		return
	}

	accountTemp,_ = s.accountService.FindOneAccount(accountModel.Username)

	if accountTemp != nil{
		c.JSON(http.StatusForbidden,common.NewError("database",errors.New("Tài khoản đã bị trùng")))
		return
	}

	_, err = s.accountService.CreateNewAccount(&accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	user.AccountId = accountModel.ID
	user.RoleId = 2

	_,err = s.userService.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("user_model", &accountModel)

	serializer := account.AccountSerializer{c}
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": user,
	})
}


func (s *HttpAccountHandler) CreateAccountStore(c *gin.Context){
	var accountModel models.Account
	err := common.Bind(c, &accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}

	if accountModel.Username == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tài khoản")))
		return
	}

	if accountModel.Password == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo mật khẩu")))
		return
	}

	accountModel.Provider = common.NormalProvider
	accountModel.DeletedAt = nil

	if accountModel.Password != common.NBRandomPassword {
		accountModel.SetPassword(accountModel.Password)
	}

	var user models.User
	user.Name = c.PostForm("name")
	if user.Name == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Chưa khai báo tên")))
		return
	}
	user.PhoneNumber = c.PostForm("phone_number")
	if user.PhoneNumber == ""{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", errors.New("Số điện thoại không hợp lệ")))
		return
	}

	user.IdentifyNumber = c.PostForm("identify_number")
	user.Address = c.PostForm("address")
	user.Imageurl = c.PostForm("imageurl")

	userTemp := &models.User{PhoneNumber:""}
	accountTemp := &models.Account{Username: ""}

	userTemp,_ = s.userService.GetUserByPhoneNumber(user.PhoneNumber)

	if userTemp != nil{
		c.JSON(http.StatusForbidden, common.NewError("database", errors.New("Số điện thoại đã bị trùng")))
		return
	}

	accountTemp,_ = s.accountService.FindOneAccount(accountModel.Username)

	if accountTemp != nil{
		c.JSON(http.StatusForbidden,common.NewError("database",errors.New("Tài khoản đã bị trùng")))
		return
	}

	storeNum,err := strconv.ParseUint(c.PostForm("store_id"),10,64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("error", errors.New("Sai thông tin cửa hàng")))
		return
	}

	_, err = s.accountService.CreateNewAccount(&accountModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	user.AccountId = accountModel.ID
	user.RoleId = 3
	user.StoreId = uint(storeNum)

	_,err = s.userService.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("user_model", &accountModel)

	serializer := account.AccountSerializer{c}
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": user,
	})
}

func  (s *HttpAccountHandler) UpdateAccount (c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var account models.Account
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	account.ID = uint(idNum)
	err = common.Bind(c, &account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.accountService.UpdateAccount(&account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&account)
}

func (s *HttpAccountHandler) DeleteAccount(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	bool,err := s.accountService.DeleteAccount(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}


func (s *HttpAccountHandler) FacebookAccountsLogin(c *gin.Context){
	fbLoginValidator := account.FBNewLoginValidator()
	if err := fbLoginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if fbLoginValidator.AccountModel.AccessToken == ""{
		c.JSON(http.StatusForbidden,common.NewError("fblogin",errors.New("Access token not found")))
		return
	}

	resp,err := facebook.Get("/me",facebook.Params{
		"fields": "id,first_name,last_name,picture,name,email",
		"access_token": fbLoginValidator.AccountModel.AccessToken,
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
	imageUrl := make(map[string]interface{})
	dataImage := make(map[string]interface{})
	resp.DecodeField("name",&fbName)
	resp.DecodeField("picture",&imageUrl)
	dataImage = imageUrl["data"].(map[string]interface{})
	err = resp.DecodeField("email",&fbEmail)
	if err != nil || fbEmail == "" {
		c.JSON(http.StatusForbidden,common.NewError("fblogin",errors.New("Facebook access token does not required email scope or facebook account does not have email")))
		return
	}

	accountModel,err := s.accountService.FindOneAccount(&models.Account{Email: fbEmail})

	if err != nil {
		var newAccount models.Account
		newAccount.Email = fbEmail
		newAccount.Provider = common.FacebookProvider
		newAccount.AccessToken = fbLoginValidator.AccountModel.AccessToken

		if _,err := s.accountService.CreateNewAccount(&newAccount); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}

		var user models.User

		user.AccountId = newAccount.ID
		user.RoleId = 1
		if dataImage["url"].(string) != ""{
			user.Imageurl = dataImage["url"].(string)
		}


		_,err = s.userService.CreateNewUser(&user)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}


		c.Set("user_model", &newAccount)
		serializer := account.AccountSerializer{C: c}
		c.JSON(http.StatusCreated, gin.H{"account": serializer.Response(),"user": user})
	}else{

		userModel,err := s.userService.GetUserByAccountId(accountModel.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}

		c.Set("user_model", accountModel)
		serializer := account.AccountSerializer{c}
		c.JSON(http.StatusOK, gin.H{"account": serializer.Response(),"user": userModel,})
	}
}

func (s *HttpAccountHandler) StoreLogin(c *gin.Context){
	loginValidator := account.NewDriverLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	accountModel,err := s.accountService.FindOneAccount(&models.Account{Username: loginValidator.AccountModel.Username})

	if err != nil {
		c.JSON(http.StatusForbidden,common.NewError("login",errors.New("Tài khoản hoặc mật khẩu không hợp lệ")))
		return
	}

	if accountModel.CheckPassword(loginValidator.Account.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Tài khoản chưa đăng ký hoặc sai mật khẩu")))
		return
	}

	userModel,err := s.userService.GetUserByAccountId(accountModel.ID)
	if userModel.RoleId != 3 {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Invalid credential")))
		return
	}

	middlewares.UpdateContextUserModel(c, accountModel.ID)
	serializer := account.DriverAccountSerializer{c}
	userRequested,err := s.userService.GetUserByAccountId(accountModel.ID)
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": userRequested ,
	})

}


func (s *HttpAccountHandler) DriverLogin(c *gin.Context){
	loginValidator := account.NewDriverLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	accountModel,err := s.accountService.FindOneAccount(&models.Account{Username: loginValidator.AccountModel.Username})

	if err != nil {
		c.JSON(http.StatusForbidden,common.NewError("login",errors.New("Tài khoản hoặc mật khẩu không hợp lệ")))
		return
	}

	if accountModel.CheckPassword(loginValidator.Account.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Tài khoản chưa đăng ký hoặc sai mật khẩu")))
		return
	}

	userModel,err := s.userService.GetUserByAccountId(accountModel.ID)
	if userModel.RoleId != 2 {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Invalid credential")))
		return
	}

	middlewares.UpdateContextUserModel(c, accountModel.ID)
	serializer := account.DriverAccountSerializer{c}
	userRequested,err := s.userService.GetUserByAccountId(accountModel.ID)
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": userRequested ,
	})

}

func (s *HttpAccountHandler) AccountsLogin(c *gin.Context){
	loginValidator := account.NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	accountModel,err := s.accountService.FindOneAccount(&models.Account{Email: loginValidator.AccountModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden,common.NewError("login",errors.New("Email hoặc mật khẩu không hợp lệ")))
		return
	}

	if accountModel.CheckPassword(loginValidator.Account.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Email chưa đăng ký hoặc sai mật khẩu")))
		return
	}
	middlewares.UpdateContextUserModel(c, accountModel.ID)
	serializer := account.AccountSerializer{c}
	userRequested,err := s.userService.GetUserByAccountId(accountModel.ID)
	c.JSON(http.StatusOK, gin.H{
		"account": serializer.Response(),
		"user": userRequested ,
	})

}



func (s *HttpAccountHandler) FcmTokenUpdate(c *gin.Context){
	id := c.Params.ByName(`id`)

	if id == ""{
		c.JSON(http.StatusUnprocessableEntity, errors.New("Account not specified"))
		return
	}
	fcmTokenValidator := account.NewFcmTokenValidator()
	if err := fcmTokenValidator.Bind(c);err != nil{
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if fcmTokenValidator.AccountModel.FcmToken == ""{
		c.JSON(http.StatusForbidden,common.NewError("token",errors.New("FCM token not found")))
		return

	}

	idNum,_ := strconv.Atoi(id)

	err := s.accountService.UpdateAccountFcmToken(idNum,fcmTokenValidator.AccountModel.FcmToken)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Update token successfully",
	})



}
