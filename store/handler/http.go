package handler

import (
	"github.com/gin-gonic/gin"
	"d2d-backend/store"
	"strconv"
	"d2d-backend/common"
	"net/http"
	"errors"
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

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listStore, err := s.storeService.GetStores(limit,page)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK,listStore)

}

func  (s *HttpStoreHandler) GetStoreById(c *gin.Context){

	id := c.Param("id")

	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	var store store.Store

	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}


	store.ID = uint(idNum)

	_,err = s.storeService.GetStoreById(&store)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK,store)


}

func  (s *HttpStoreHandler) CreateStore(c *gin.Context){

	var store store.Store

	err:= common.Bind(c,&store)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_,err = s.storeService.CreateNewStore(&store)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK,store)
}

func  (s *HttpStoreHandler) UpdateStore(c *gin.Context){
	id := c.Param("id")

	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	var store store.Store

	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}


	store.ID = uint(idNum)

	err = common.Bind(c,&store)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}


	_,err = s.storeService.UpdateStore(&store)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	c.JSON(http.StatusOK,&store)

}

func (s *HttpStoreHandler) DeleteStore(c *gin.Context){
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

	bool,err := s.storeService.DeleteStore(int(idNum))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}

	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})

}