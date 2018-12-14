package handler

import (
	"d2d-backend/common"
	"d2d-backend/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"d2d-backend/notification"
	"d2d-backend/placedOrder"
	"d2d-backend/models"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpUserHandler struct {
	userService user.UserService
	notificationService notification.NotificationService
	placedOrderService placedOrder.PlacedOrderService
}

func NewUserHttpHandler(e *gin.RouterGroup,
						service user.UserService,
						notifService notification.NotificationService,
						orderService placedOrder.PlacedOrderService	) (*HttpUserHandler){
	handler := &HttpUserHandler{
		userService: service,
		notificationService: notifService,
		placedOrderService: orderService,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpUserHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpUserHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllUser)
	e.GET("/:id", s.GetUserById)
	e.GET("/:id/notifications", s.GetNotificationsByUserId)
	e.GET("/:id/notifications/unread", s.GetTotalUnreadMessage)
	e.GET("/:id/delivery/active",s.GetActivePlacedOrderByDeliveryId)
	e.GET("/:id/delivery/instore",s.GetInStorePlacedOrderByDeliveryId)
	e.GET("/:id/store/active",s.GetActivePlacedOrderByStoreId)
	e.GET("/:id/placedorders", s.GetPlacedOrdersByUserId)
	e.POST("/", s.CreateUser)
	e.PUT("/:id", s.UpdateUser)
	e.DELETE("/:id", s.DeleteUser)
	e.POST("/:id/location",s.AddNewUserShippingLocation)
	e.PUT("/:id/location/:locationId",s.EditUserShippingLocation)
	e.DELETE("/:id/location/:locationId",s.DeleteUserShippingLocation)

}



func (s *HttpUserHandler) GetActivePlacedOrderByDeliveryId(c *gin.Context){
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.placedOrderService.GetListActiveOrdersByDeliveryId(uint(idNum),limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)

}

func (s *HttpUserHandler) GetInStorePlacedOrderByDeliveryId(c *gin.Context){
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.placedOrderService.GetInStorePlacedOrdersByDeliveryId(uint(idNum),limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func (s *HttpUserHandler) GetAllUser(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listUser, err := s.userService.GetUser(limit,page)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listUser)
}

func  (s *HttpUserHandler) GetUserById(c *gin.Context){
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
	userModel,err := s.userService.GetUserById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,userModel)
}

func  (s *HttpUserHandler) CreateUser(c *gin.Context){
	var userModel models.User
	err:= common.Bind(c,&userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if userModel.Name == "" || strings.TrimSpace(userModel.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if userModel.PhoneNumber == "" || strings.TrimSpace(userModel.PhoneNumber) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Phone number is empty")))
		return
	}
	_,err = s.userService.CreateNewUser(&userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,userModel)
}

func  (s *HttpUserHandler) UpdateUser(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var userModel models.User
	idNum,err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	userModel.ID = uint(idNum)
	err = common.Bind(c,&userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	if userModel.Name == "" || strings.TrimSpace(userModel.Name) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty name",errors.New("Name is empty")))
		return
	}
	if userModel.PhoneNumber == "" || strings.TrimSpace(userModel.PhoneNumber) == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("Empty description",errors.New("Description is empty")))
		return
	}
	_,err = s.userService.UpdateUser(&userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&userModel)
}

func (s *HttpUserHandler) DeleteUser(c *gin.Context){
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
	isDeleted,err := s.userService.DeleteUser(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(isDeleted)})
}


func (s *HttpUserHandler) GetNotificationsByUserId(c *gin.Context)  {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.notificationService.GetNotificationByUserId(limit, page, int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func (s *HttpUserHandler) GetTotalUnreadMessage(c *gin.Context)  {
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
	count, err := s.notificationService.GetTotalUnreadNotification(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"count": count,
	})
}

func (s *HttpUserHandler) GetPlacedOrdersByUserId(c *gin.Context)  {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
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
	list, err := s.placedOrderService.GetListOrdersByUserId(limit, page, int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}




func (s *HttpUserHandler) GetActivePlacedOrderByStoreId(c *gin.Context){
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
	list, err := s.placedOrderService.GetListActivePlacedOrdersByStoreId(uint(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,gin.H{"records": list})
}

func (s *HttpUserHandler) AddNewUserShippingLocation(c *gin.Context){
	var userShippingModel models.UserShippingLocation
	err:= common.Bind(c,&userShippingModel)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Lỗi khi đồng bộ địa chỉ")))
		return
	}
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	userId, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã tài khoản khong hợp lệ")))
		return
	}

	userShippingModel.UserId = uint(userId)

	if userShippingModel.Latitude == 0 || userShippingModel.Longitude == 0 || userShippingModel.PhoneNumber  == "" ||
			userShippingModel.ReceiverName == "" || userShippingModel.ShippingAddress == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Địa chỉ không hợp lệ")))
		return
	}

	if strings.TrimSpace(userShippingModel.PhoneNumber)  == "" ||
		strings.TrimSpace(userShippingModel.ReceiverName) == "" || strings.TrimSpace(userShippingModel.ShippingAddress) == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Địa chỉ không hợp lệ")))
		return
	}

	addedUserShippingLocation,err := s.userService.SaveNewShippingLocation(&userShippingModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Địa chỉ không hợp lệ")))
		return
	}

	c.JSON(http.StatusOK, addedUserShippingLocation)
}

func (s *HttpUserHandler) EditUserShippingLocation(c *gin.Context){
	var userShippingModel models.UserShippingLocation
	err:= common.Bind(c,&userShippingModel)
	if err != nil {
		//c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Lỗi khi đồng bộ địa chỉ")))
		return
	}
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	locationId := c.Param("locationId")
	if locationId == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	locationIdNum, err := strconv.ParseUint(locationId,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã tài khoản khong hợp lệ")))
		return
	}

	userShippingModel.ID = uint(locationIdNum)

	updatedModel,err := s.userService.UpdateUserLocation(&userShippingModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", err))
		return
	}

	c.JSON(http.StatusOK, updatedModel)
}

func (s *HttpUserHandler) DeleteUserShippingLocation(c *gin.Context){
	var userShippingModel models.UserShippingLocation
	err:= common.Bind(c,&userShippingModel)
	if err != nil {
		//c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Lỗi khi đồng bộ địa chỉ")))
		return
	}
	id := c.Param("id")
	if id == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	locationId := c.Param("locationId")
	if locationId == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}

	locationIdNum, err := strconv.ParseUint(locationId,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã tài khoản kh6ong hợp lệ")))
		return
	}

	deleted,err := s.userService.DeleteUserLocation(uint(locationIdNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  deleted,
	})
}