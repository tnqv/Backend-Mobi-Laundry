package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"d2d-backend/common"
)

func OrdersRouterRegister(router *gin.RouterGroup){
	//router.GET("/orders",AccountsLogin)
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