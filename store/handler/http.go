package handler

import (
	"github.com/gin-gonic/gin"
	"d2d-backend/store"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpStoreHandler struct {
	storeService store.StoreService
}

func NewStoreHttpHandler(e *gin.RouterGroup, service store.StoreService) (*HttpStoreHandler){
	handler := &HttpStoreHandler{
		storeService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpStoreHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpStoreHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){

	e.GET("/", s.GetAllStore)
	e.POST("/", s.CreateStore)
	e.GET("/:id", s.GetStoreById)
	e.PUT("/:id",s.UpdateStore)
	e.DELETE("/:id", s.DeleteStore)
}

func (s *HttpStoreHandler) GetAllStore(c *gin.Context) {

}

func  (s *HttpStoreHandler) GetStoreById(c *gin.Context){

}

func  (s *HttpStoreHandler) CreateStore(c *gin.Context){

}

func  (s *HttpStoreHandler) UpdateStore(c *gin.Context){

}

func (s *HttpStoreHandler) DeleteStore(c *gin.Context){

}