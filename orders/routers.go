package orders

import (
	"d2d-backend/common"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
	router.POST("/createorder", CreateOrder)
	router.GET("/cusid",GetOrdersbyCustomerID)
	router.POST("/getorders",GetOrders)}

func ServicesRouterRegister(router *gin.RouterGroup){
	router.GET("/",GetServices)
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
	//var order PlacedOrder
	//c.Bind(&order)
	//customer := getCustomerInformations(1)
	//order.CustomerID = customer.ID
	//order.DeliveryAddress = customer.ShippingAddress
	//order.DeliveryLongitude = customer.Longitude
	//order.DeliveryLatitude = customer.Latitude
	//order.OrderStatusID = 1
	//createPlaceOrder(&order)
	//c.JSON(http.StatusCreated, order)
	accountID, _ := strconv.ParseUint(c.PostForm("accountID"), 10, 64)
	var order PlacedOrder
	orderModelValidator := NewOrderModelValidator()
	if err := orderModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	c.Bind(&order)
	order.TimePlaced = time.Now()
	order.OrderStatusID = CreateOrderStatus(1, uint(accountID))
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
}

func CreateOrderStatus(statusID uint, accountID uint) (uint) {
	var orderStatus OrderStatus
	orderStatus.StatusID = statusID
	orderStatus.AccountID = accountID
	orderStatus.StatusChangedTime = time.Now()
	createOrderStatus(&orderStatus)
	return orderStatus.ID
}

//Minh's function
func GetOrdersbyCustomerID(c *gin.Context){
	var userid uint
	userid = 1
	//c.Bind(&userid)
	data,err := getAllOrdersBasedOnCustomerID(&userid)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
		return
	}
	c.JSON(http.StatusOK,data)
}

func GetOrders(c *gin.Context){
	orders,err := getOrders()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
		return
	}
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultPostForm("limit", "10"))
	db := common.GetDB()
	paginator := pagination.Pagging(&pagination.Param{
		DB: db,
		Page: page,
		Limit: limit,
		ShowSQL: true,
	}, &orders)
	c.JSON(http.StatusOK,paginator)
}