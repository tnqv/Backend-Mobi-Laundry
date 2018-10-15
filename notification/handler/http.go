package handler

import (
	"d2d-backend/common"
	"d2d-backend/notification"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpNotificationHandler struct {
	notificationService notification.NotificationService
}

func NewNotificationHttpHandler(e *gin.RouterGroup, service notification.NotificationService) *HttpNotificationHandler {
	handler := &HttpNotificationHandler{
		notificationService: service,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpNotificationHandler) UnauthorizedRoutes(e *gin.RouterGroup) {
	e.GET("/", s.GetAllNotifications)
	e.GET("/:id", s.GetNotificationById)
	e.POST("/", s.CreateNotification)
	e.PUT("/:id", s.UpdateNotification)
	e.DELETE("/:id", s.DeleteNotification)
}

func (s *HttpNotificationHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup) {

}

func (s *HttpNotificationHandler) GetAllNotifications(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listStore, err := s.notificationService.GetNotifications(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, listStore)
}

func (s *HttpNotificationHandler) GetNotificationById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	idNum, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	notification, err := s.notificationService.GetNotificationById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, notification)
}

func  (s *HttpNotificationHandler) CreateNotification(c *gin.Context){
	var notification notification.Notification
	err:= common.Bind(c,&notification)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if notification.Content == "" || strings.TrimSpace(notification.Content) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty content",errors.New("Content is empty")))
		return
	}
	_,err = s.notificationService.CreateNewNotification(&notification)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,notification)
}

func  (s *HttpNotificationHandler) UpdateNotification(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var notification notification.Notification
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	notification.ID = uint(idNum)
	err = common.Bind(c,&notification)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if notification.Content == "" || strings.TrimSpace(notification.Content) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty content",errors.New("Content is empty")))
		return
	}
	_,err = s.notificationService.UpdateNotification(&notification)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&notification)
}

func (s *HttpNotificationHandler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	idNum, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	bool, err := s.notificationService.DeleteNotification(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK, ResponseError{Message: strconv.FormatBool(bool)})
}

