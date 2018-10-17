package handler

import (
	"d2d-backend/common"
	"d2d-backend/role"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)
type ResponseError struct {
	Message string `json:"message"`
}

type HttpRoleHandler struct {
	roleService role.RoleService
}

func NewRoleHttpHandler(e *gin.RouterGroup, service role.RoleService) (*HttpRoleHandler){
	handler := &HttpRoleHandler{
		roleService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpRoleHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpRoleHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllRoles)
	e.GET("/:id", s.GetRoleById)
	e.POST("/", s.CreateRole)
	e.PUT("/:id",s.UpdateRole)
	e.DELETE("/:id", s.DeleteRole)
}

func (s *HttpRoleHandler) GetAllRoles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listStore, err := s.roleService.GetRoles(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listStore)
}

func (s *HttpRoleHandler) GetRoleById(c *gin.Context){
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
	role, err := s.roleService.GetRoleById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, role)
}

func (s *HttpRoleHandler) CreateRole(c *gin.Context){
	var role role.Role
	err := common.Bind(c, &role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if role.Name == "" || strings.TrimSpace(role.Name) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if role.Description == "" || strings.TrimSpace(role.Description) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_, err = s.roleService.CreateNewRole(&role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, role)
}

func  (s *HttpRoleHandler) UpdateRole(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var role role.Role
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	role.ID = uint(idNum)
	err = common.Bind(c, &role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if role.Name == "" || strings.TrimSpace(role.Name) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if role.Description == "" || strings.TrimSpace(role.Description) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_, err = s.roleService.UpdateRole(&role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&role)
}

func (s *HttpRoleHandler) DeleteRole(c *gin.Context){
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
	bool,err := s.roleService.DeleteRole(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(bool)})
}