package handler

import (
	"d2d-backend/common"
	"d2d-backend/serviceOrder"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpServiceOrderHandler struct {
	serviceOrderService serviceOrder.ServiceOrderService
}

func NewServiceOrderHttpHandler(e *gin.RouterGroup, service serviceOrder.ServiceOrderService) (*HttpServiceOrderHandler){
	handler := &HttpServiceOrderHandler{
		serviceOrderService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpServiceOrderHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllServiceOrders)
	e.GET("/:id", s.GetServiceOrderById)
	e.POST("/", s.CreateServiceOrder)
	e.PUT("/:id",s.UpdateServiceOrder)
	e.DELETE("/:id", s.DeleteServiceOrder)
}

func (s *HttpServiceOrderHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){

}

func (s *HttpServiceOrderHandler) GetAllServiceOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	list, err := s.serviceOrderService.GetServiceOrders(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func (s *HttpServiceOrderHandler) GetServiceOrderById(c *gin.Context){
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
	serviceOrder, err := s.serviceOrderService.GetServiceOrderById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, serviceOrder)
}

func (s *HttpServiceOrderHandler) CreateServiceOrder(c *gin.Context){
	var serviceOrder serviceOrder.ServiceOrder
	err := common.Bind(c, &serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.serviceOrderService.CreateNewServiceOrder(&serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, serviceOrder)
}

func  (s *HttpServiceOrderHandler) UpdateServiceOrder(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var serviceOrder serviceOrder.ServiceOrder
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	serviceOrder.ID = uint(idNum)
	err = common.Bind(c, &serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.serviceOrderService.UpdateServiceOrder(&serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&serviceOrder)
}

func (s *HttpServiceOrderHandler) DeleteServiceOrder(c *gin.Context){
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
	bool,err := s.serviceOrderService.DeleteServiceOrder(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}



