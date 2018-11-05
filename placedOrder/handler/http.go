package handler

import (
	"d2d-backend/common"
	"d2d-backend/orderStatus"
	"d2d-backend/placedOrder"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"d2d-backend/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpPlacedOrderHandler struct {
	placedOrderService placedOrder.PlacedOrderService
	orderStatusService orderStatus.OrderStatusService
}

type HttpOrderStatusHandler struct {
	orderStatusService orderStatus.OrderStatusService
}

func NewPlacedOrderHttpHandler(e *gin.RouterGroup,
							   service placedOrder.PlacedOrderService,
							   	osService orderStatus.OrderStatusService) (*HttpPlacedOrderHandler){
	handler := &HttpPlacedOrderHandler{
		placedOrderService: service,
		orderStatusService: osService,
	}
	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpPlacedOrderHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpPlacedOrderHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllPlacedOrders)
	e.GET("/:id", s.GetPlacedOrderById)
	//e.GET("/order-code/:orderCode",s.GetPlacedOrderByOrderCode)
	e.POST("/", s.CreatePlacedOrder)
	e.PUT("/:id",s.UpdatePlacedOrder)
	e.DELETE("/:id", s.DeletePlacedOrder)
	e.PUT("/:id/status/:statusId",s.UpdateStatusPlacedOrder)
}



func (s *HttpPlacedOrderHandler) GetAllPlacedOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	listStore, err := s.placedOrderService.GetPlacedOrders(limit, page)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK,listStore)
}

func (s *HttpPlacedOrderHandler) GetPlacedOrderByOrderCode(c *gin.Context){
	orderCode := c.Param("orderCode");
	if orderCode == ""{
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid orderCode")))
		return
	}

	placedOrderModel,err := s.placedOrderService.GetPlacedOrderByOrderCode(orderCode)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, placedOrderModel)
}

func (s *HttpPlacedOrderHandler) GetPlacedOrderById(c *gin.Context){
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
	placedOrderModel, err := s.placedOrderService.GetPlacedOrderById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, placedOrderModel)
}

func (s *HttpPlacedOrderHandler) CreatePlacedOrder(c *gin.Context){
	var placedOrderModel models.PlacedOrder
	err := common.Bind(c, &placedOrderModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}

	placedOrderModel.DeletedAt = nil

	if placedOrderModel.UserID == 0{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Invalid params", errors.New("User không hợp lệ")))
		return
	}

	if placedOrderModel.DeliveryAddress == "" || placedOrderModel.DeliveryLongitude == 0 || placedOrderModel.DeliveryLatitude == 0{
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Invalid params", errors.New("Địa điểm không hợp lệ")))
		return
	}

	placedOrderModel.TimePlaced = time.Now()
	placedOrderModel.OrderCode = time.Now().Format("20060102150405")
	var tempOrderStatus models.OrderStatus
	tempOrderStatus.StatusID = 1
	tempOrderStatus.UserId = placedOrderModel.UserID
	tempOrderStatus.StatusChangedTime = time.Now()
	newOrderStatusModel , err := orderStatus.OrderStatusService.CreateNewOrderStatus(s.orderStatusService, &tempOrderStatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	placedOrderModel.OrderStatusId = newOrderStatusModel.ID
	_, err = s.placedOrderService.CreateNewPlacedOrder(&placedOrderModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	newOrderStatusModel.PlacedOrderID = placedOrderModel.ID

	_,err = s.orderStatusService.UpdateOrderStatus(newOrderStatusModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	//Insert to firebase store
	// push notification

	c.JSON(http.StatusOK, placedOrderModel)
}

func  (s *HttpPlacedOrderHandler) UpdatePlacedOrder(c *gin.Context){
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid id")))
		return
	}
	var placedOrderModel models.PlacedOrder
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}
	placedOrderModel.ID = uint(idNum)
	err = common.Bind(c, &placedOrderModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Error binding", err))
		return
	}
	_, err = s.placedOrderService.UpdatePlacedOrder(&placedOrderModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,&placedOrderModel)
}

func (s *HttpPlacedOrderHandler) DeletePlacedOrder(c *gin.Context){
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
	isDeleted,err := s.placedOrderService.DeletePlacedOrder(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Database", err))
		return
	}
	c.JSON(http.StatusOK,ResponseError{Message: strconv.FormatBool(isDeleted)})
}

func (s *HttpPlacedOrderHandler) UpdateStatusPlacedOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã order không hợp lệ")))
		return
	}

	statusId := c.Param("statusId")
	if statusId == "" {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Mã trạng thái không hợp lệ ")))
		return
	}

	var placedOrderModel models.PlacedOrder
	idNum, err := strconv.ParseUint(id,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format id")))
		return
	}

	idStatusNum,err := strconv.ParseUint(id,10,32)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format status id")))
		return
	}

	placedOrderModel.ID = uint(idNum)

	placedOrderUpdate,err := s.placedOrderService.GetPlacedOrderById(int(idNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Đơn hàng không tồn tại")))
		return
	}
	userId := c.PostForm("user_id")
	userIdNum,err := strconv.ParseUint(userId,10,32)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format status id")))
		return

	}

	switch(idStatusNum){
		case 2:

				//Store accept order
				s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(2,uint(userIdNum),placedOrderUpdate)
				//insert to firebase to delivery
				//Push notification

		case 3:
			//Delivery take order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(3,uint(userIdNum),placedOrderUpdate)
			// delete from firebase
			// push notification to user
		case 4:
			//Delivery confirm order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(4,uint(userIdNum),placedOrderUpdate)
			// push notification to user


		case 5:
			//Delivery has deliveried to Store
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(5,uint(userIdNum),placedOrderUpdate)
			// push notification

		case 6:
			//Store change status to laundring
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(6,uint(userIdNum),placedOrderUpdate)
			// push notification to user & delivery
		case 7:
			//Store change status to finish
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(7,uint(userIdNum),placedOrderUpdate)
			// push notification to user & delivery

		case 8:
			//Delivery change status to deliver
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(8,uint(userIdNum),placedOrderUpdate)
			// push notification to user
		case 9:
			//Customer pay
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(9,uint(userIdNum),placedOrderUpdate)
		case 10:
			//Store cancel order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(10,uint(userIdNum),placedOrderUpdate)
		default :
			c.JSON(http.StatusBadRequest,common.NewError("param",errors.New("Sai thông tin trạng thái")))
			return

	}

	c.JSON(http.StatusOK,placedOrderUpdate)
}