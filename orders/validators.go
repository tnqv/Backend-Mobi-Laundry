package orders

import (
	"d2d-backend/common"
	"github.com/gin-gonic/gin"
)

type OrderModelValidator struct {
	Order struct{
		DeliveryAddress string		`form:"delivery_address" json:"delivery_address" binding:"exists,min=4,max=255"`
		DeliveryLatitude float32	`form:"delivery_latitude" json:"delivery_latitude" binding:"exists,min=4,max=255"`
		DeliveryLongitude float32	`form:"delivery_longitude" json:"delivery_longitude" binding:"exists,min=4,max=255"`
		EstimatedCapacity float32	`form:"estimated_capacity" json:"estimated_capacity" binding:"exists,min=4,max=255"`
	} `json:"order"`
	orderModel PlacedOrder `json:"-"`
}

func (self *OrderModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	return nil
}

func NewOrderModelValidator() OrderModelValidator {
	orderModelValidator := OrderModelValidator{}
	return orderModelValidator
}
