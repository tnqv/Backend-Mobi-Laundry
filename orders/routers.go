package orders

import (
	"d2d-backend/common"
	"github.com/gin-gonic/gin"
	"net/http"
)
func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
	router.POST("/createorder", CreateOrder)
	router.GET("/cusid",GetOrdersbyCustomerID)
	router.GET("/tenorders",GetTenOrders)}

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
	var order PlacedOrder
	orderModelValidator := NewOrderModelValidator()
	if err := orderModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	c.Bind(&order)
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
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

func GetTenOrders(c *gin.Context){
	data,err := getTenOrders()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
		return
	}
	c.JSON(http.StatusOK,data)
}