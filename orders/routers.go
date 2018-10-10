package orders

import (
	"d2d-backend/common"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
	router.POST("/", CreateOrder)
    router.GET("/user/:userId",GetOrdersbyAccountID)
	router.GET("/",GetOrders)
	//router.POST("/order/service", CreateOrderServicesForPlacedOrder)
	router.POST("/category", CreateCategory)
	router.PUT("/category", UpdateCategory)
	router.GET("/category/:categoryId", GetCategory)
	router.DELETE("/:categoryId", DeleteCategory)}

func ServicesRouterRegister(router *gin.RouterGroup){
	router.GET("/", GetListServices)
	router.GET("/:serviceId", GetService)
	router.POST("/", CreateService)
	router.PUT("/", UpdateService)
	router.DELETE("/:serviceId", DeleteService)
}

func ServiceOrdersRouterRegister(router *gin.RouterGroup)  {
	router.GET("/", GetListServiceOrders)
	router.GET("/:serviceOrderId", GetServiceOrder)
	router.POST("/", CreateServiceOrder)
	router.PUT("/", UpdateServiceOrder)
	router.DELETE("/:serviceOrderId", DeleteServiceOrder)
}

func OrderStatusesRouterRegister(router *gin.RouterGroup)  {
	router.GET("/", GetListOrderStatuses)
	router.GET("/:orderStatusId", GetOrderStatus)
	router.POST("/", CreateOrderStatus)
	router.PUT("/", UpdateOrderStatus)
	router.DELETE("/:orderStatusId", DeleteOrderStatus)
}

func NotificationsRouterRegister(router *gin.RouterGroup)  {
	//router.GET("/", GetListNotifications)
	router.GET("/:notificationId", GetNotifications)

}

func GetServices(c *gin.Context) {
	data, err := getAllServicesBasedOnCategory()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, data)
}

//DuyNQ's function
func CreateOrder (c *gin.Context) {
	userID, _ := strconv.ParseUint(c.PostForm("userID"), 10, 64)
	var order PlacedOrder
	orderModelValidator := NewOrderModelValidator()
	if err := orderModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	c.Bind(&order)
	order.UserID = getCustomerInformations(uint(userID)).ID
	order.TimePlaced = time.Now()
	order.OrderCode = time.Now().Format("20060102150405")
	//order.OrderStatusID = CreateOrderStatus(1, uint(userID))
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
}

/*func UpdateOrderStatus(c * gin.Context) {
	userID, _ := strconv.ParseUint(c.PostForm("userID"), 10, 64)
	statusID, _ := strconv.ParseUint(c.PostForm("statusID"), 10, 64)
	orderID, _ := strconv.ParseUint(c.PostForm("orderID"), 10, 64)
	orderStatusID := CreateOrderStatus(uint(statusID), uint(userID))
	order, err := updateOrderStatus(uint(orderID), orderStatusID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, order)
}

func CreateOrderStatus(statusID uint, userID uint) (uint) {
	var orderStatus OrderStatus
	orderStatus.StatusID = statusID
	orderStatus.UserID = userID
	orderStatus.StatusChangedTime = time.Now()
	createOrderStatus(&orderStatus)
	return orderStatus.ID
}*/


//Minh's function

func GetOrdersbyAccountID(c *gin.Context){
	userID, _ := strconv.ParseUint(c.Param("userId"),10,64)
	fmt.Println(userID)
	//data,err := getAllOrdersBasedOnAccountID(uint(accountID))
	//if err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
	//	return
	//}
	var orders []PlacedOrder
	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	db := common.GetDB()

	db = db.Where("user_id = ?", userID)

	paginator := pagination.Pagging(&pagination.Param{
		DB: db,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &orders)


	c.JSON(http.StatusOK,paginator)
}

func GetOrders(c *gin.Context){
	//orders,err := getOrders()
	//if err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
	//	return
	//}

	var orders []PlacedOrder

	page, _ := strconv.Atoi(c.DefaultQuery(common.Page, common.PageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(common.Limit, common.LimitDefault))
	db := common.GetDB()
	paginator := pagination.Pagging(&pagination.Param{
		DB: db,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &orders)
	c.JSON(http.StatusOK,paginator)
}

func CreateCategory(c *gin.Context){
	var category Category
	c.Bind(&category)
	createCategory(&category)
	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context)  {
	var category Category
	c.Bind(&category)
	err := updateCategory(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, category)
}

func GetCategory(c *gin.Context){
	cateID, _ := strconv.ParseUint(c.Params.ByName("categoryId"),10,64)
	cate,err := getCategory(uint(cateID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, cate)
}
func DeleteCategory(c *gin.Context){
	cateID, _ := strconv.ParseUint(c.Params.ByName("categoryId"), 10, 64)
	err := deleteCategory(uint(cateID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, "The category has been deleted!")
}

//Order Status
func CreateOrderStatus(c *gin.Context){
	var orderstatus OrderStatus
	c.Bind(&orderstatus)
	orderstatus.StatusChangedTime = time.Now()
	createOrderStatus(&orderstatus)
	c.JSON(http.StatusCreated, orderstatus)
}

func UpdateOrderStatus(c *gin.Context)  {
	var orderstatus OrderStatus
	c.Bind(&orderstatus)
	err := updateOrderStatus(&orderstatus)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, orderstatus)
}

func GetOrderStatus(c *gin.Context){
	orderStatusID, _ := strconv.ParseUint(c.Params.ByName("orderStatusId"),10,64)
	orderstatus,err := getOrderStatus(uint(orderStatusID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, orderstatus)
}

func GetListOrderStatuses(c *gin.Context)  {
	list, err := getListOrderStatuses()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func DeleteOrderStatus(c *gin.Context){
	orderstatusID, _ := strconv.ParseUint(c.Params.ByName("orderStatusId"), 10, 64)
	err := deleteOrderStatus(uint(orderstatusID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, "The order status has been deleted!")
}
//SERVICE_ORDERS ENTITY
func GetListServiceOrders(c *gin.Context)  {
	list, err := getListServiceOrders()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetServiceOrder(c *gin.Context)  {
	serviceOrderId, _ := strconv.ParseUint(c.Params.ByName("serviceOrderId"), 10, 64)
	serviceOrder, err := getServiceOrder(uint(serviceOrderId))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, serviceOrder)
}

func CreateServiceOrder(c *gin.Context){
	var serviceOrder ServiceOrder
	c.Bind(&serviceOrder)
	err := createServiceOrder(&serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, serviceOrder)
}

func UpdateServiceOrder(c * gin.Context)  {
	var serviceOrder ServiceOrder
	c.Bind(&serviceOrder)
	err := updateServiceOrder(&serviceOrder)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, serviceOrder)
}

func DeleteServiceOrder(c * gin.Context)  {
	id, _ := strconv.ParseUint(c.Params.ByName("serviceOrderId"), 10, 64)
	err := deleteServiceOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
}
//END SERVICE_ORDERS ENTITY

//SERVICE ENTITY
func GetListServices(c *gin.Context) {
	list, err := getListServices()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetService(c *gin.Context)  {
	serviceId, _ := strconv.ParseUint(c.Params.ByName("serviceId"), 10, 64)
	service, err := getService(uint(serviceId))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, service)
}

func CreateService(c *gin.Context)  {
	var service Service
	c.Bind(&service)
	err := createService(&service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, service)
}

func UpdateService(c *gin.Context)  {
	var service Service
	c.Bind(&service)
	err := updateService(&service)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, service)
}

func DeleteService(c *gin.Context)  {
	serviceId, _ := strconv.ParseUint(c.Params.ByName("serviceId"), 10, 64)
	err := deleteService(uint(serviceId))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, "Deleted!")
}
//END SERVICE ENTITY

//Notification
func GetNotifications(c *gin.Context){
	notificationID, _ := strconv.ParseUint(c.Params.ByName("notificationId"),10,64)
	notification,err := getOrderStatus(uint(notificationID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, notification)
}

//func GetListNotifications(c *gin.Context)  {
//	list, err := getListNotifications()
//	if err != nil {
//		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
//		return
//	}
//	c.JSON(http.StatusOK, list)
//}
