package handler

import (
	"d2d-backend/common"
	"d2d-backend/orderStatus"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpOrderStatusHandler struct {
	orderStatusService orderStatus.OrderStatusService
}

func NewOrderStatusHttpHandler(e *gin.RouterGroup, service orderStatus.OrderStatusService) (*HttpOrderStatusHandler){
	handler := &HttpOrderStatusHandler{
		orderStatusService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpOrderStatusHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllOrderStatus)
	e.GET("/:id", s.GetOrderStatusById)
	e.POST("/", s.CreateOrderStatus)
	e.PUT("/:id", s.UpdateOrderStatus)
	e.DELETE("/:id", s.DeleteOrderStatus)
}

func (s *HttpOrderStatusHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	/*e.POST("/", s.CreateCategory)
	e.GET("/:id", s.GetCategoryById)
	e.PUT("/:id",s.UpdateCategory)
	e.DELETE("/:id", s.DeleteCategory)*/
}

func (s *HttpOrderStatusHandler) GetAllOrderStatus(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listOrderStatus, err := s.orderStatusService.GetOrderStatus(limit,page)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listOrderStatus)
}

func  (s *HttpOrderStatusHandler) GetOrderStatusById(c *gin.Context){
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
	orderStatus,err := s.orderStatusService.GetOrderStatusById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,orderStatus)
}

func  (s *HttpOrderStatusHandler) CreateOrderStatus(c *gin.Context){
	var orderStatus orderStatus.OrderStatus
	err:= common.Bind(c,&orderStatus)
	orderStatus.StatusChangedTime = time.Now()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if orderStatus.Description == "" || strings.TrimSpace(orderStatus.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.orderStatusService.CreateNewOrderStatus(&orderStatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,orderStatus)
}

func  (s *HttpOrderStatusHandler) UpdateOrderStatus(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var orderStatus orderStatus.OrderStatus
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	orderStatus.ID = uint(idNum)
	err = common.Bind(c,&orderStatus)
	orderStatus.StatusChangedTime = time.Now()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if orderStatus.Description == "" || strings.TrimSpace(orderStatus.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.orderStatusService.UpdateOrderStatus(&orderStatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&orderStatus)
}

func (s *HttpOrderStatusHandler) DeleteOrderStatus(c *gin.Context){
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
	bool,err := s.orderStatusService.DeleteOrderStatus(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}

