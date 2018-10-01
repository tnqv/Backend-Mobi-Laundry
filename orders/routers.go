package orders

import (
	"d2d-backend/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
	router.POST("/", CreateOrder)
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

func CreateOrder (c *gin.Context)  {
	var order PlacedOrder
	err := c.Bind(&order)
	if err != nil {
		c.AbortWithError(400, err)
	}
	createPlaceOrder(&order)
	c.JSON(http.StatusCreated, order)
}