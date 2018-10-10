package handler

import (
	"d2d-backend/common"
	"d2d-backend/service"
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
}

func (s *HttpServiceHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.POST("/", s.CreateService)
	e.GET("/:id", s.GetServiceById)
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

}

func  (s *HttpServiceHandler) CreateService(c *gin.Context){

}

func  (s *HttpServiceHandler) UpdateService(c *gin.Context){

}

func (s *HttpServiceHandler) DeleteService(c *gin.Context){

}