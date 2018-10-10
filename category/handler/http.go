package handler

import (
	"d2d-backend/category"
	"d2d-backend/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpCategoryHandler struct {
	categoryService category.CategoryService
}

func NewStoreHttpHandler(e *gin.RouterGroup, service category.CategoryService) (*HttpCategoryHandler){
	handler := &HttpCategoryHandler{
		categoryService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpCategoryHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllCategory)
	e.GET("/:id", s.GetCategoryById)
	e.POST("/", s.CreateCategory)
	e.PUT("/:id",s.UpdateCategory)
	e.DELETE("/:id", s.DeleteCategory)
}

func (s *HttpCategoryHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){


	/*e.POST("/", s.CreateCategory)
	e.GET("/:id", s.GetCategoryById)
	e.PUT("/:id",s.UpdateCategory)
	e.DELETE("/:id", s.DeleteCategory)*/
}

func (s *HttpCategoryHandler) GetAllCategory(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listCategory, err := s.categoryService.GetCategory(limit,page)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listCategory)
}

func  (s *HttpCategoryHandler) GetCategoryById(c *gin.Context){
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
	category,err := s.categoryService.GetCategoryById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,category)
}

func  (s *HttpCategoryHandler) CreateCategory(c *gin.Context){
	var category category.Category
	err:= common.Bind(c,&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if category.Name == "" || strings.TrimSpace(category.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if category.Description == "" || strings.TrimSpace(category.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.categoryService.CreateNewCategory(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,category)
}

func  (s *HttpCategoryHandler) UpdateCategory(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var category category.Category
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	category.ID = uint(idNum)
	err = common.Bind(c,&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if category.Name == "" || strings.TrimSpace(category.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if category.Description == "" || strings.TrimSpace(category.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.categoryService.UpdateCategory(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&category)
}

func (s *HttpCategoryHandler) DeleteCategory(c *gin.Context){
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
	bool,err := s.categoryService.DeleteCategory(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}

