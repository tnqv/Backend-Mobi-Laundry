package handler

import (
	"d2d-backend/category"
	"d2d-backend/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"d2d-backend/models"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpCategoryHandler struct {
	categoryService category.CategoryService
}

func NewCategoryHttpHandler(e *gin.RouterGroup, service category.CategoryService) (*HttpCategoryHandler){
	handler := &HttpCategoryHandler{
		categoryService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpCategoryHandler) UnauthorizedRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllCategory)
	e.GET("/:id", s.GetCategoryById)
}

func (s *HttpCategoryHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){

	e.POST("/", s.CreateCategory)
	e.PUT("/:id", s.UpdateCategory)
	e.DELETE("/:id", s.DeleteCategory)


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
	categoryModel,err := s.categoryService.GetCategoryById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,categoryModel)
}

func  (s *HttpCategoryHandler) CreateCategory(c *gin.Context){
	var categoryModel models.Category
	err:= common.Bind(c,&categoryModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if categoryModel.Name == "" || strings.TrimSpace(categoryModel.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if categoryModel.Description == "" || strings.TrimSpace(categoryModel.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.categoryService.CreateNewCategory(&categoryModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,categoryModel)
}

func  (s *HttpCategoryHandler) UpdateCategory(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var categoryModel models.Category
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	categoryModel.ID = uint(idNum)
	err = common.Bind(c,&categoryModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if categoryModel.Name == "" || strings.TrimSpace(categoryModel.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if categoryModel.Description == "" || strings.TrimSpace(categoryModel.Description) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.categoryService.UpdateCategory(&categoryModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&categoryModel)
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
	isDeleted,err := s.categoryService.DeleteCategory(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(isDeleted)})
}