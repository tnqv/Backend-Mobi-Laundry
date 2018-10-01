package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"d2d-backend/common"
)

func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
	router.GET("/cusid",GetOrdersbyCustomerID)
	router.GET("/tenorders",GetTenOrders)
}

func ServicesRouterRegister(router *gin.RouterGroup){
	router.GET("/",GetServices)
}

func GetServices(c *gin.Context){
	data,err := getAllServicesBasedOnCategory()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database",err))
		return
	}

	c.JSON(http.StatusOK,data)
}

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