package handler

import (
	"d2d-backend/account"
	"d2d-backend/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpAccountHandler struct {
	accountService account.AccountService
}

func NewAccountHttpHandler(e *gin.RouterGroup, service account.AccountService) (*HttpAccountHandler){
	handler := &HttpAccountHandler{
		accountService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpAccountHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllAccounts)
	e.GET("/:id", s.GetAccountById)
	e.POST("/", s.CreateAccount)
	e.PUT("/:id",s.UpdateAccount)
	e.DELETE("/:id", s.DeleteAccount)
}

func (s *HttpAccountHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){

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

func (s *HttpAccountHandler) CreateAccount(c *gin.Context){
	var account account.Account
	err := common.Bind(c, &account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.accountService.CreateNewAccount(&account)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, account)
}

func  (s *HttpAccountHandler) UpdateAccount (c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var account account.Account
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