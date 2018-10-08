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
	router.PUT("/status", UpdateOrderStatus)
	router.PUT("/serviceOrder/:serviceOrderId", UpdateServiceOrder)
}

func ServicesRouterRegister(router *gin.RouterGroup){
	router.GET("/",GetServices)
	router.DELETE("/")
}

func ServiceOrderRouterRegister(router *gin.RouterGroup)  {
	router.DELETE("/:serviceOrderId", DeleteServiceOrder)
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
	order.OrderStatusID = CreateOrderStatus(1, uint(userID))
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
}

func UpdateOrderStatus(c * gin.Context) {
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
}

func UpdateServiceOrder(c * gin.Context)  {
	id, _ := strconv.ParseUint(c.Params.ByName("serviceOrderId"), 10, 64)
	quantity, _ := strconv.ParseUint(c.PostForm("quantity"), 10, 64)
	err := updateQuantity(uint(id), uint(quantity))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
}

func DeleteServiceOrder(c * gin.Context)  {
	id, _ := strconv.ParseUint(c.Params.ByName("serviceOrderId"), 10, 64)
	err := deleteServiceOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
}

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

