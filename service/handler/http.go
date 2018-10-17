package handler

import (
	"d2d-backend/common"
	"d2d-backend/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpServiceHandler struct {
	serviceService service.ServiceService
}

func NewServiceHttpHandler(e *gin.RouterGroup, service service.ServiceService)(*HttpServiceHandler) {
	handler := &HttpServiceHandler{
		serviceService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpServiceHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllServices)
	e.GET("/:id", s.GetServiceById)
}

func (s *HttpServiceHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.POST("/", s.CreateService)
	e.PUT("/:id",s.UpdateService)
	e.DELETE("/:id", s.DeleteService)
}

func (s *HttpServiceHandler) GetAllServices(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listService, err := s.serviceService.GetServices(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, listService)
}

func  (s *HttpServiceHandler) GetServiceById(c *gin.Context){
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
	service, err := s.serviceService.GetServiceById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, service)
}

func  (s *HttpServiceHandler) CreateService(c *gin.Context){
	var service service.Service
	err := common.Bind(c, &service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.serviceService.CreateNewService(&service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, service)
}

func  (s *HttpServiceHandler) UpdateService(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var service service.Service
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	service.ID = uint(idNum)
	err = common.Bind(c, &service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.serviceService.UpdateService(&service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&service)
}

func (s *HttpServiceHandler) DeleteService(c *gin.Context){
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
	bool,err := s.serviceService.DeleteService(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}