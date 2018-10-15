package handler

import (
	"d2d-backend/common"
	"d2d-backend/orderStatus"
	"d2d-backend/placedOrder"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	orderStatusRepository "d2d-backend/orderStatus/repository"
	orderStatusService "d2d-backend/orderStatus/service"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpPlacedOrderHandler struct {
	placedOrderService placedOrder.PlacedOrderService
}

type HttpOrderStatusHandler struct {
	orderStatusService orderStatus.OrderStatusService
}

func NewPlacedOrderHttpHandler(e *gin.RouterGroup, service placedOrder.PlacedOrderService) (*HttpPlacedOrderHandler){
	handler := &HttpPlacedOrderHandler{
		placedOrderService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpPlacedOrderHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllPlacedOrders)
	e.GET("/:id", s.GetPlacedOrderById)
	e.POST("/", s.CreatePlacedOrder)
	e.PUT("/:id",s.UpdatePlacedOrder)
	e.DELETE("/:id", s.DeletePlacedOrder)
}

func (s *HttpPlacedOrderHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){

}

func (s *HttpPlacedOrderHandler) GetAllPlacedOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listStore, err := s.placedOrderService.GetPlacedOrders(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listStore)
}

func (s *HttpPlacedOrderHandler) GetPlacedOrderById(c *gin.Context){
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
	placedOrder, err := s.placedOrderService.GetPlacedOrderById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, placedOrder)
}

func (s *HttpPlacedOrderHandler) CreatePlacedOrder(c *gin.Context){
	var placedOrder placedOrder.PlacedOrder
	err := common.Bind(c, &placedOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	placedOrder.TimePlaced = time.Now()
	placedOrder.OrderCode = time.Now().Format("20060102150405")
	var tempOrderStatus orderStatus.OrderStatus
	tempOrderStatus.StatusID = 1
	tempOrderStatus.UserID = placedOrder.UserID
	tempOrderStatus.StatusChangedTime = time.Now()
	orderStatusRepository := orderStatusRepository.NewMysqlOrderStatusRepository()
	orderStatusService := orderStatusService.NewOrderStatusService(orderStatusRepository)
	newOrderStatus, err := orderStatus.OrderStatusService.CreateNewOrderStatus(orderStatusService, &tempOrderStatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	placedOrder.OrderStatusID = newOrderStatus.ID
	_, err = s.placedOrderService.CreateNewPlacedOrder(&placedOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, placedOrder)
}

func  (s *HttpPlacedOrderHandler) UpdatePlacedOrder(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var placedOrder placedOrder.PlacedOrder
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	placedOrder.ID = uint(idNum)
	err = common.Bind(c, &placedOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.placedOrderService.UpdatePlacedOrder(&placedOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&placedOrder)
}

func (s *HttpPlacedOrderHandler) DeletePlacedOrder(c *gin.Context){
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
	bool,err := s.placedOrderService.DeletePlacedOrder(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}

