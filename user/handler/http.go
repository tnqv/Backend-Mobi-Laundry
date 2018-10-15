package handler

import (
	"d2d-backend/common"
	"d2d-backend/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"d2d-backend/notification"
	"d2d-backend/placedOrder"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpUserHandler struct {
	userService user.UserService
	notificationService notification.NotificationService
	placedOrderService placedOrder.PlacedOrderService
}

func NewUserHttpHandler(e *gin.RouterGroup,
						service user.UserService,
						notifService notification.NotificationService,
						orderService placedOrder.PlacedOrderService	) (*HttpUserHandler){
	handler := &HttpUserHandler{
		userService: service,
		notificationService: notifService,
		placedOrderService: orderService,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpUserHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllUser)
	e.GET("/:id", s.GetUserById)
	e.GET("/:id/notifications", s.GetNotificationsByUserId)
	e.GET("/:id/placedorders", s.GetPlacedOrdersByUserId)
	e.POST("/", s.CreateUser)
	e.PUT("/:id", s.UpdateUser)
	e.DELETE("/:id", s.DeleteUser)
}

func (s *HttpUserHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){


}

func (s *HttpUserHandler) GetAllUser(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listUser, err := s.userService.GetUser(limit,page)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listUser)
}

func  (s *HttpUserHandler) GetUserById(c *gin.Context){
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
	user,err := s.userService.GetUserById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,user)
}

func  (s *HttpUserHandler) CreateUser(c *gin.Context){
	var user user.User
	err:= common.Bind(c,&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if user.Name == "" || strings.TrimSpace(user.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if user.PhoneNumber == "" || strings.TrimSpace(user.PhoneNumber) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Phone number is empty")))
		return
	}
	_,err = s.userService.CreateNewUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,user)
}

func  (s *HttpUserHandler) UpdateUser(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var user user.User
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	user.ID = uint(idNum)
	err = common.Bind(c,&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if user.Name == "" || strings.TrimSpace(user.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if user.PhoneNumber == "" || strings.TrimSpace(user.PhoneNumber) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&user)
}

func (s *HttpUserHandler) DeleteUser(c *gin.Context){
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
	bool,err := s.userService.DeleteUser(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}


func (s *HttpUserHandler) GetNotificationsByUserId(c *gin.Context)  {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.notificationService.GetNotificationByUserId(limit, page, int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}


func (s *HttpUserHandler) GetPlacedOrdersByUserId(c *gin.Context)  {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.placedOrderService.GetListOrdersByUserId(limit, page, int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}