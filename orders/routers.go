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
	router.POST("/createorderservice", CreateOrderServicesForPlacedOrder)
	router.POST("/createcategory", CreateCategory)
	router.PUT("/updatecategory", UpdateCategory)
	router.GET("/getcategory/:categoryId", GetCategory)
	router.DELETE("/:categoryId", DeleteCategory)}

func ServicesRouterRegister(router *gin.RouterGroup){
	router.GET("/",GetServices)
	router.DELETE("/")
}

func ServiceOrderRouterRegister(router *gin.RouterGroup)  {
	router.DELETE("/:serviceOrderId",)
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
	accountID, _ := strconv.ParseUint(c.PostForm("userID"), 10, 64)
	var order PlacedOrder
	orderModelValidator := NewOrderModelValidator()
	if err := orderModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	c.Bind(&order)
	order.UserID = getCustomerInformations(uint(accountID)).ID
	order.TimePlaced = time.Now()
	order.OrderCode = time.Now().Format("20060102150405")
	order.OrderStatusID = CreateOrderStatus(1, uint(accountID))
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
}

func CreateOrderStatus(statusID uint, accountID uint) (uint) {
	var orderStatus OrderStatus
	orderStatus.StatusID = statusID
	orderStatus.UserID = accountID
	orderStatus.StatusChangedTime = time.Now()
	createOrderStatus(&orderStatus)
	return orderStatus.ID
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

func CreateOrderServicesForPlacedOrder(c *gin.Context){
	orderserviceID, _ := strconv.ParseUint(c.PostForm("service_id"),10,64)
	placedorderID, _ := strconv.ParseUint(c.PostForm("placed_order_id"),10,64)
	quantity, _ := strconv.ParseUint(c.PostForm("quantity"),10,64)
	var orderservice ServiceOrder
	c.Bind(&orderservice)
	orderservice.ServiceID = uint(orderserviceID)
	orderservice.PlacedOrderID = uint(placedorderID)
	orderservice.Quantity = uint(quantity)
	createOrderService(&orderservice)
	c.JSON(http.StatusCreated, orderservice)
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