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
	"d2d-backend/user"

	//models2 "d2d-backend/models/message"
	"encoding/json"
	"d2d-backend/serviceOrder"
	"d2d-backend/store"
	"fmt"
	"log"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpPlacedOrderHandler struct {
	placedOrderService placedOrder.PlacedOrderService
	orderStatusService orderStatus.OrderStatusService
	serviceOrderService serviceOrder.ServiceOrderService
	storeService store.StoreService
	userSer user.UserService
}

type HttpOrderStatusHandler struct {
	orderStatusService orderStatus.OrderStatusService
}

func NewPlacedOrderHttpHandler(e *gin.RouterGroup,
							   service placedOrder.PlacedOrderService,
							   	osService orderStatus.OrderStatusService,
							   		servOrdService serviceOrder.ServiceOrderService,
							   			sService store.StoreService) (*HttpPlacedOrderHandler){
	handler := &HttpPlacedOrderHandler{
		placedOrderService: service,
		orderStatusService: osService,
		serviceOrderService: servOrdService,
		storeService: sService,
	}


	handler.UnauthorizedRoutes(e)
	return handler
}

func (s *HttpPlacedOrderHandler) SetUserService(userServ user.UserService){
		s.userSer = userServ
}

func (s *HttpPlacedOrderHandler) UnauthorizedRoutes(e *gin.RouterGroup){

}

func (s *HttpPlacedOrderHandler) AuthorizedRequiredRoutes(e *gin.RouterGroup){
	e.GET("/", s.GetAllPlacedOrders)
	e.GET("/:id", s.GetPlacedOrderById)
	e.PUT("/:id/status/:statusId",s.UpdateStatusPlacedOrder)
	//e.GET("/order-code/:orderCode",s.GetPlacedOrderByOrderCode)
	e.POST("/", s.CreatePlacedOrder)
	e.PUT("/:id",s.UpdatePlacedOrder)
	e.DELETE("/:id", s.DeletePlacedOrder)

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

	if !common.IsLocationaInHoChiMinhCity(placedOrderModel.DeliveryLatitude,placedOrderModel.DeliveryLongitude) {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("Invalid params", errors.New("Địa điểm không nằm trong Hồ Chí Minh")))
		return
	}

	placedOrderModel.TimePlaced = time.Now()
	placedOrderModel.OrderCode = time.Now().Format("20060102150405")
	placedOrderModel.VerifyCode = common.RandNumberString(6)
	var tempOrderStatus models.OrderStatus
	tempOrderStatus.StatusID = common.ORDER_CREATED_STATUS
	tempOrderStatus.UserId = placedOrderModel.UserID
	tempOrderStatus.StatusChangedTime = time.Now()
	tempOrderStatus.Description = fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_1,placedOrderModel.OrderCode)
	newOrderStatusModel , err := orderStatus.OrderStatusService.CreateNewOrderStatus(s.orderStatusService, &tempOrderStatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	placedOrderModel.OrderStatusId = newOrderStatusModel.StatusID
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

	//Mapping order

	stores,err := s.storeService.GetAllStores()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	var shortestDistanceStoreId uint
	var shortestDistance float64
	shortestDistanceStoreId = 0
	for i:= 0; i < len(stores); i++ {

		distance := common.Distance(float64(placedOrderModel.DeliveryLatitude),float64(placedOrderModel.DeliveryLongitude),float64(stores[i].Latitude),float64(stores[i].Longitude))
		if i == 0 && distance != 0 {
			shortestDistance = distance
			shortestDistanceStoreId = stores[i].ID
		}

		if distance < shortestDistance {
			shortestDistance = distance
			shortestDistanceStoreId =  stores[i].ID
		}
	}

	if shortestDistance == 0 || shortestDistanceStoreId == 0 {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", errors.New("Lỗi cửa hàng")))
		return
	}

	placedOrderModel.StoreID = shortestDistanceStoreId
	_,err = s.placedOrderService.UpdatePlacedOrder(&placedOrderModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	//Insert to firebase store
	common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderModel)

	// push notification
	common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderModel)

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

	idStatusNum,err := strconv.ParseUint(statusId,10,32)

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
		c.JSON(http.StatusNotAcceptable, common.NewError("param", errors.New("Invalid format user_id")))
		return

	}

	userModel,err := s.userSer.GetUserById(int(userIdNum))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Tài khoản không hợp lệ")))
		return

	}

//case 1:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_1,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 2:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_2,notifMess.OrderCode,notifMess.Store.Name)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 3:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_3,notifMess.Delivery.Name)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 4:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_4,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 5:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_5,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 6:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_6,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	tokensPush = append(tokensPush,notifMess.Delivery.Account.FcmToken)
//	case 7:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_7,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	tokensPush = append(tokensPush,notifMess.Delivery.Account.FcmToken)
//	case 8:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_8,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 9:
//	message =	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_9)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
//	case 10:
//	message = fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_10,notifMess.OrderCode)
//	tokensPush = append(tokensPush,notifMess.User.Account.FcmToken)
	switch(idStatusNum){
		case common.ORDER_CREATED_STATUS:

		case common.ORDER_ACCEPTED_BY_STORE:

				if userModel.StoreId == 0{
					c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Tài khoản hiện tại không hợp lệ cho chức năng này")))
					return
				}
				//Store accept order

				description := fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_2,placedOrderUpdate.OrderCode,placedOrderUpdate.Store.Name)
				placedOrderUpdate,err = s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_ACCEPTED_BY_STORE,uint(userIdNum),description,placedOrderUpdate)
				if err != nil {
					c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Lỗi xảy ra khi cập nhật")))
					return
				}
				//message queue

				common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderUpdate)
				common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)

		case common.ORDER_ACCEPTED_BY_DELIVERY:
			//Delivery take order
			placedOrderUpdate.DeliveryID = userModel.ID
			description := fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_3,userModel.Name)
			placedOrderUpdate,err = s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_ACCEPTED_BY_DELIVERY,uint(userIdNum),description,placedOrderUpdate)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Lỗi xảy ra khi cập nhật")))
				return
			}
			// delete from firebase
			common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderUpdate)
			// push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
		case common.ORDER_CONFIRM:
			//Delivery confirm order
			verifyCode := c.PostForm("verify_code")

			if placedOrderUpdate.VerifyCode != verifyCode {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Mã xác nhận không hợp lệ")))
				return
			}

			serviceOrdersJson := c.PostForm("service_orders")
			var serviceOrdersReq []*models.ServiceOrder
			err = json.Unmarshal([]byte(serviceOrdersJson),&serviceOrdersReq)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Dịch vụ đơn hàng không hợp lệ")))
				return
			}

			serviceOrdersReq,err := s.serviceOrderService.CreateListServiceOrders(serviceOrdersReq)

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Lỗi tạo đơn")))
				return
			}

			var total float32
			for _,value := range serviceOrdersReq {
					total += value.Price
			}

			description  :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_4,placedOrderUpdate.OrderCode)
			placedOrderUpdate.ServiceTotalPrice = total

			placedOrderUpdate,err = s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_CONFIRM,uint(userIdNum),description,placedOrderUpdate)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Lỗi xảy ra khi cập nhật")))
				return
			}

			log.Println(placedOrderUpdate.Delivery.Account.FcmToken)

			// push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)

		case common.ORDER_IN_STORE:
			verifyCode := c.PostForm("verify_code")

			if placedOrderUpdate.VerifyCode != verifyCode {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Mã xác nhận không hợp lệ")))
				return
			}
			//Delivery has deliveried to Store
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_5,placedOrderUpdate.OrderCode)
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_IN_STORE,uint(userIdNum),description,placedOrderUpdate)

			// push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)

		case common.ORDER_LAUNDRYING:
			//Store change status to laundring
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_6,placedOrderUpdate.OrderCode)
			placedOrderUpdate,err = s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_LAUNDRYING,uint(userIdNum),description,placedOrderUpdate)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, common.NewError("param", errors.New("Lỗi xảy ra khi cập nhật")))
				return
			}


			// push notification to user & delivery
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
		case common.ORDER_FINISH_LAUNDRYING:
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_7,placedOrderUpdate.OrderCode)
			//Store change status to finish
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_FINISH_LAUNDRYING,uint(userIdNum),description,placedOrderUpdate)

			// push notification to user & delivery
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)

		case common.ORDER_DELIVERY_BACK_TO_CUSTOMER:
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_8,placedOrderUpdate.OrderCode)
			//Delivery change status to deliver
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_DELIVERY_BACK_TO_CUSTOMER,uint(userIdNum),description,placedOrderUpdate)

			// push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
		case common.ORDER_COMPLETE:
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_9,placedOrderUpdate.OrderCode)
			//Customer pay
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_COMPLETE,uint(userIdNum),description,placedOrderUpdate)

			//Push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
		case common.ORDER_CANCEL:
			description :=	fmt.Sprintf(common.MESSAGE_PATTERN_STATUS_10,placedOrderUpdate.OrderCode)
			//Store cancel order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.ORDER_CANCEL,uint(userIdNum),description,placedOrderUpdate)

			//Push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
			// delete from firebase
			common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderUpdate)
		case common.DELIVERY_CANNOT_RECEIVE_CLOTHES:
			description := ""
			//Store cancel order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.DELIVERY_CANNOT_RECEIVE_CLOTHES,uint(userIdNum),description,placedOrderUpdate)

			//Push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
			//common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderUpdate)

		case common.DELIVERY_CANNOT_GIVE_BACK_CLOTHES:
			description := ""
			//Store cancel order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.DELIVERY_CANNOT_GIVE_BACK_CLOTHES,uint(userIdNum),description,placedOrderUpdate)

			//Push notification to user
			common.ProduceMessage(common.NOTIFICATION_QUEUE,placedOrderUpdate)
		case common.DELIVERY_REFUSE_TO_DELIVER:
			description := ""
			//Store cancel order
			s.placedOrderService.UpdatePlacedOrderAndCreateNewOrderStatus(common.DELIVERY_REFUSE_TO_DELIVER,uint(userIdNum),description,placedOrderUpdate)

			// Insert back to firebase
			common.ProduceMessage(common.FIREBASE_QUEUE,placedOrderUpdate)
		default :
			c.JSON(http.StatusBadRequest,common.NewError("param",errors.New("Sai thông tin trạng thái")))
			return

	}

	c.JSON(http.StatusOK,placedOrderUpdate)
}